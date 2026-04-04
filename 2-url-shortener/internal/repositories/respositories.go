package repositories

import (
	"context"
	"errors"
	"sd_concepts/url_shortener/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{db: pool}
}

func (r *Repository) GetURLRepository(ctx context.Context, hash string) (string, error) {
	var longURL string

	err := r.db.QueryRow(ctx, "SELECT long_url FROM url WHERE hash = $1", hash).Scan(&longURL)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", errors.New("URL Not Found")
		}

		return "", err
	}

	return longURL, nil
}

func (r *Repository) CheckHashRepository(ctx context.Context, hash string) (*models.URL, error) {
	var obj models.URL
	err := r.db.QueryRow(ctx, "SELECT hash, long_url FROM url WHERE hash = $1", hash).Scan(&obj.Hash, &obj.LongURL)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("hash doesn't exist")
		}
		return nil, err
	}

	return &obj, nil
}

func (r *Repository) CreateURLRepository(ctx context.Context, hash, longURL string) error {
	query := "INSERT INTO url (hash, long_url) VALUES ($1, $2)"

	_, err := r.db.Exec(ctx, query, hash, longURL)
	if err != nil {
		return err
	}

	return nil
}
