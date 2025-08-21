package category

import (
	"context"

	"datenote/datenote/storage"
)

type categoryStore interface {
	CreateCategory(context.Context, storage.Category) (int64, error)
	GetCategory(context.Context, int64) (storage.Category, error)
	GetAllCategories(context.Context) ([]storage.Category, error)
	UpdateCategory(context.Context, storage.Category) error
	DeleteCategory(context.Context, int64) error
	SearchCategory(context.Context, string) ([]storage.Category, error)
}

type CoreCategorySvc struct {
	store categoryStore
}

func NewCoreCategorySvc(s categoryStore) *CoreCategorySvc {
	return &CoreCategorySvc{
		store: s,
	}
}

func (cs CoreCategorySvc) CreateCategory(ctx context.Context, t storage.Category) (int64, error) {
	return cs.store.CreateCategory(ctx, t)
}

func (cs CoreCategorySvc) GetCategory(ctx context.Context, id int64) (storage.Category, error) {
	return cs.store.GetCategory(ctx, id)
}

func (cs CoreCategorySvc) GetAllCategories(ctx context.Context) ([]storage.Category, error) {
	return cs.store.GetAllCategories(ctx)
}

func (cs CoreCategorySvc) UpdateCategory(ctx context.Context, t storage.Category) error {
	return cs.store.UpdateCategory(ctx, t)
}

func (cs CoreCategorySvc) DeleteCategory(ctx context.Context, id int64) error {
	return cs.store.DeleteCategory(ctx, id)
}

func (cs CoreCategorySvc) SearchCategory(ctx context.Context, q string) ([]storage.Category, error) {
	return cs.store.SearchCategory(ctx, q)
}
