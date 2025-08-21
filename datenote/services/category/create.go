package category

import (
	"context"

	"datenote/datenote/storage"
	cpb "datenote/gunk/v1/category"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CategorySvc) CreateCategory(ctx context.Context, req *cpb.CreateCategoryRequest) (*cpb.CreateCategoryResponse, error) {
	category := storage.Category{
		ID:    req.Category.ID,
		Title: req.Category.Title,
	}

	id, err := s.core.CreateCategory(context.Background(), category)

	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to create category")
	}

	return &cpb.CreateCategoryResponse{
		ID: id,
	}, nil
}
