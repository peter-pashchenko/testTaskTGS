package grpc

import (
	"context"
	"fmt"
	"github.com/peter-pashchenko/imagesService/internal/application/dto"
	pb "github.com/peter-pashchenko/imagesService/internal/generated/grpc/images"
	"github.com/peter-pashchenko/imagesService/internal/models"
	"go.uber.org/zap"
	"golang.org/x/sync/semaphore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedImageServiceServer

	service Service

	semList *semaphore.Weighted
	semRW   *semaphore.Weighted

	logger *zap.Logger
}

type Service interface {
	Save(ctx context.Context, img *models.Image) error
	Get(ctx context.Context, name string) (*models.Image, error)
	List(ctx context.Context) ([]*models.Image, error)
}

func (s *Server) SaveImage(ctx context.Context, image *pb.Image) (*pb.SaveImageReponse, error) {
	if err := s.semRW.Acquire(ctx, 1); err != nil {
		return nil, status.Errorf(codes.Unavailable, "please try again later")
	}
	defer s.semRW.Release(1)

	err := s.service.Save(ctx, &models.Image{Name: image.Name, Data: image.Data})
	if err != nil {
		s.logger.Error(
			"mistake in saving the image",
			zap.Error(err),
			zap.Any("image name", image.Name))
		return nil, status.Error(codes.Internal, "couldn't save image")
	}

	return &pb.SaveImageReponse{Status: fmt.Sprintf("image with name %s was saved", image.Name)}, nil
}
func (s *Server) GetByName(ctx context.Context, name *pb.Name) (*pb.Image, error) {
	if err := s.semRW.Acquire(ctx, 1); err != nil {
		return nil, status.Errorf(codes.Unavailable, "please try again later")
	}
	defer s.semRW.Release(1)

	res, err := s.service.Get(ctx, name.Value)
	s.logger.Debug("geeting following answer",
		zap.Any("result", res))

	if err != nil {
		s.logger.Error(
			"error getting the image by name",
			zap.Any("name requested", name.Value),
			zap.Error(err))
		return nil, status.Error(codes.Internal, "couldn't find the image")
	}

	if res == nil {
		return &pb.Image{}, status.Error(codes.InvalidArgument, "no image found under this name")
	}

	return &pb.Image{Name: res.Name, Data: res.Data}, nil
}

func (s *Server) ListAll(ctx context.Context, r *pb.Empty) (*pb.ListImages, error) {
	if err := s.semList.Acquire(ctx, 1); err != nil {
		return nil, status.Errorf(codes.Unavailable, "please try again later")
	}
	defer s.semList.Release(1)

	res, err := s.service.List(ctx)

	if err != nil {
		s.logger.Error(
			"error getting the image list",
			zap.Error(err))
		return nil, status.Error(codes.Internal, "couldn't get the image list")
	}

	return dto.ConvertToPBList(res), nil
}
