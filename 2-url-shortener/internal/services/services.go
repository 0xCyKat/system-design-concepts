package services

import (
	"context"
	"errors"
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
	baseHash := getHash(longUrl)

	existing, err := svc.repo.CheckHashRepository(ctx, baseHash)
	if err == nil && existing != nil {
		if existing.LongURL == longUrl {
			return existing.Hash, nil
		}

		saltedHash := getHash("mod" + longUrl)

		existingSalted, errSalted := svc.repo.CheckHashRepository(ctx, saltedHash)
		if errSalted == nil && existingSalted != nil {
			return "", errors.New("salted collision occurred")
		}

		if err := svc.repo.CreateURLRepository(ctx, saltedHash, longUrl); err != nil {
			return "", err
		}
		return saltedHash, nil
	}

	if err := svc.repo.CreateURLRepository(ctx, baseHash, longUrl); err != nil {
		return "", err
	}
	return baseHash, nil
}
