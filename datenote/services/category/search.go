package category

import (
	"context"

	cpb "datenote/gunk/v1/category"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CategorySvc) SearchCategory(ctx context.Context, req *cpb.SearchCategoryRequest) (*cpb.SearchCategoryResponse, error) {
	query := req.GetSearchCategoryQuery()
	categories, err := s.core.SearchCategory(context.Background(), query)
	var c []*cpb.Category
	for _, category := range categories {
		c = append(c, &cpb.Category{
			ID:    category.ID,
			Title: category.Title,
		})
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to search category")
	}

	return &cpb.SearchCategoryResponse{
		SearchCategoryResult: c,
	}, nil
}
