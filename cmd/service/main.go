package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/peter-pashchenko/imagesService/config"
	imagesServer "github.com/peter-pashchenko/imagesService/internal/application/grpc"
	pb "github.com/peter-pashchenko/imagesService/internal/generated/grpc/images"
	_ "github.com/peter-pashchenko/imagesService/internal/models"
	"github.com/peter-pashchenko/imagesService/internal/modules/repository/images"
	"github.com/peter-pashchenko/imagesService/internal/modules/services/images"
	"github.com/peter-pashchenko/imagesService/pkg/logger"
	"github.com/peter-pashchenko/imagesService/pkg/psql"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const (
	exitStatusOk          = 0
	exitStatusError       = 1
	maxListRequests int64 = 100
	maxRWRequests   int64 = 10
)

func main() {
	_ = godotenv.Load( /*"../../.env"*/)

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config init error %s", err)
	}

	os.Exit(run(cfg))
}

func run(cfg *config.Config) (exitCode int) {
	log := logger.New(cfg.Log.Level)
	defer log.Sync()

	defer func() {
		if paniceErr := recover(); paniceErr != nil {
			log.Error(
				"recover panic",
				zap.Any("error", paniceErr),
			)
			exitCode = exitStatusError
		}
	}()

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM)
	defer stop()

	db, err := psql.Connect(
		ctx,
		log,
		int(maxListRequests+maxListRequests),
		psql.WithHost(cfg.PG.Host),
		psql.WithPort(cfg.PG.Port),
		psql.WithUser(cfg.PG.User),
		psql.WithPass(cfg.PG.Pass),
		psql.WithDatabase(cfg.PG.Database),
		psql.WithMigrations("./db/migrations"),
	)

	defer func() {
		if db != nil {
			err = db.Close()
			log.Debug("closing db...")
			if err != nil {
				log.Error(
					"error closing db",
					zap.Error(err))
			}
		}

	}()

	if err != nil {
		log.Error(
			"db init error",
			zap.Any("error", err),
		)
		return exitStatusError
	}

	repo := imagesRepository.New(db, log)
	service := imagesServices.New(repo, log)
	serviceHandler := imagesServer.New(
		imagesServer.WithLogger(log),
		imagesServer.WithImagesService(service),
		imagesServer.WithSemaphoreList(maxListRequests),
		imagesServer.WithSemaphoreRW(maxRWRequests),
	)

	grpcServer := grpc.NewServer()
	pb.RegisterImageServiceServer(grpcServer, serviceHandler)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPC.Port))
	if err != nil {
		log.Error(
			"grpc failed to listen",
			zap.Error(err),
		)
		return exitStatusError
	}

	errChan := make(chan error)

	go func() {
		log.Info("starting grpc server")

		if err = grpcServer.Serve(lis); err != nil {
			log.Error(
				"grpc failed to serve",
				zap.Error(err))
			errChan <- err
		}
	}()

	defer func() {
		grpcServer.GracefulStop()

		if err != nil {
			log.Error(
				"http server shutdown error",
				zap.Error(err))
		}
	}()

	select {
	case err = <-errChan:
		log.Error(
			"grpc server error,shutting down",
			zap.Error(err),
		)
	case <-ctx.Done():
		log.Info("gracefully shutting down")
	}

	return exitStatusOk
}
