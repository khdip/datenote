package category

import (
	"context"

	cpb "datenote/gunk/v1/category"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CategorySvc) GetAllCategories(ctx context.Context, req *cpb.GetAllCategoriesRequest) (*cpb.GetAllCategoriesResponse, error) {
	categories, err := s.core.GetAllCategories(context.Background())
	var c []*cpb.Category
	for _, category := range categories {
		c = append(c, &cpb.Category{
			ID:    category.ID,
			Title: category.Title,
		})
	}

	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to get categories")
	}

	return &cpb.GetAllCategoriesResponse{
		Categories: c,
	}, nil
}
