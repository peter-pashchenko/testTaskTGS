package grpc

import (
	"go.uber.org/zap"
	"golang.org/x/sync/semaphore"
)

func New(opts ...Option) *Server {
	server := &Server{}

	for _, opt := range opts {
		opt(server)
	}

	return server
}

type Option func(*Server)

func WithLogger(logger *zap.Logger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}

func WithImagesService(service Service) Option {
	return func(s *Server) {
		s.service = service
	}
}

func WithSemaphoreList(maxListRequest int64) Option {
	return func(s *Server) {
		s.semList = semaphore.NewWeighted(maxListRequest)
	}
}

func WithSemaphoreRW(maxRWRequest int64) Option {
	return func(s *Server) {
		s.semRW = semaphore.NewWeighted(maxRWRequest)
	}
}
