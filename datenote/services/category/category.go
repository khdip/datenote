package category

import (
	"context"

	"datenote/datenote/storage"
	cpb "datenote/gunk/v1/category"
)

type categoryCoreStore interface {
	CreateCategory(context.Context, storage.Category) (int64, error)
	GetCategory(context.Context, int64) (storage.Category, error)
	GetAllCategories(context.Context) ([]storage.Category, error)
	UpdateCategory(context.Context, storage.Category) error
	DeleteCategory(context.Context, int64) error
	SearchCategory(context.Context, string) ([]storage.Category, error)
}

type CategorySvc struct {
	cpb.UnimplementedCategoryServiceServer
	core categoryCoreStore
}

func NewCategoryServer(c categoryCoreStore) *CategorySvc {
	return &CategorySvc{
		core: c,
	}
}
