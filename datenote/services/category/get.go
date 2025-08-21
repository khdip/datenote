package category

import (
	"context"

	cpb "datenote/gunk/v1/category"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CategorySvc) GetCategory(ctx context.Context, req *cpb.GetCategoryRequest) (*cpb.GetCategoryResponse, error) {
	id := req.ID
	category, err := s.core.GetCategory(context.Background(), id)

	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to get category")
	}

	return &cpb.GetCategoryResponse{
		Category: &cpb.Category{
			ID:    category.ID,
			Title: category.Title,
		},
	}, nil
}
