package imagesServices

import (
	"context"
	"github.com/peter-pashchenko/imagesService/internal/models"
	"go.uber.org/zap"
)

type Service struct {
	repo Repo

	logger *zap.Logger
}

type Repo interface {
	SaveImage(ctx context.Context, image *models.Image) error
	GetByName(ctx context.Context, name string) (*models.Image, error)
	ListAll(ctx context.Context) ([]*models.Image, error)
}

func New(repo Repo, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) Save(ctx context.Context, img *models.Image) error {
	return s.repo.SaveImage(ctx, img)
}

func (s *Service) Get(ctx context.Context, name string) (*models.Image, error) {
	return s.repo.GetByName(ctx, name)

}
func (s *Service) List(ctx context.Context) ([]*models.Image, error) {
	return s.repo.ListAll(ctx)
}
