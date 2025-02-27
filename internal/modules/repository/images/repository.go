package imagesRepository

import (
	"context"
	"database/sql"
	"github.com/peter-pashchenko/imagesService/internal/models"
	"go.uber.org/zap"
	"time"
)

type Repo struct {
	db     *sql.DB
	logger *zap.Logger
}

func New(db *sql.DB, logger *zap.Logger) *Repo {
	return &Repo{db: db, logger: logger}
}

func (r *Repo) SaveImage(ctx context.Context, image *models.Image) error {
	query := `INSERT INTO images (name,data,created_at) 
    				 VALUES ($1,$2,$3)
    				 ON CONFLICT(name)
    				 DO UPDATE SET 
    				 data=$2,
    				 updated_at=$3;`

	r.logger.Debug("query string is ready", zap.String("query", query))

	_, err := r.db.ExecContext(
		ctx,
		query,
		image.Name,
		image.Data,
		time.Now())

	if err == nil {
		r.logger.Debug("data saved to db successfully")
	}

	return err
}
func (r *Repo) GetByName(ctx context.Context, name string) (*models.Image, error) {
	query := `SELECT name,data FROM images WHERE name = $1`

	row := r.db.QueryRowContext(
		ctx,
		query,
		name)

	var img models.Image
	err := row.Scan(&img.Name, &img.Data)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &img, err
}
func (r *Repo) ListAll(ctx context.Context) ([]*models.Image, error) {
	query := `SELECT name,created_at,updated_at FROM images `

	rows, err := r.db.QueryContext(
		ctx,
		query)

	if err != nil {
		return nil, err
	}

	images := make([]*models.Image, 0)

	for rows.Next() {
		var updatedAt sql.NullTime
		image := &models.Image{}
		err = rows.Scan(&image.Name, &image.CreatedAt, &updatedAt)
		if err != nil {
			return nil, err
		}
		if updatedAt.Valid {
			image.UpdatedAt = &updatedAt.Time
		}
		images = append(images, image)
	}
	return images, nil
}
