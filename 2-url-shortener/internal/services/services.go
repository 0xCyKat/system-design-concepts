package services

import (
	"context"
	"errors"
	"sd_concepts/url_shortener/internal/models"
	"sd_concepts/url_shortener/internal/repositories"
)

type Service struct {
	repo *repositories.Repository
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (svc *Service) GetUrlService(ctx context.Context, hash string) (string, error) {
	return svc.repo.GetURLRepository(ctx, hash)
}

func (svc *Service) CreateURLService(ctx context.Context, longUrl string) (string, error) {
	url, _ := svc.checkHashExists(ctx, longUrl)
	if url != nil && url.LongURL == longUrl {
		return url.Hash, errors.New("already exists")
	}

	newUrl := "mod" + longUrl
	hash := getHash(newUrl)
	return hash, svc.repo.CreateURLRepository(ctx, hash, longUrl)
}

func (svc *Service) checkHashExists(ctx context.Context, longUrl string) (*models.URL, error) {
	hash := getHash(longUrl)
	return svc.repo.CheckHashRepository(ctx, hash)
}
