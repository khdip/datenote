package event

import (
	"context"

	"datenote/datenote/storage"
)

type eventStore interface {
	CreateEvent(context.Context, storage.Event) (int64, error)
	GetEvent(context.Context, int64) (storage.Event, error)
	GetAllEvents(context.Context) ([]storage.Event, error)
	UpdateEvent(context.Context, storage.Event) error
	DeleteEvent(context.Context, int64) error
	SearchEvent(context.Context, string) ([]storage.Event, error)
}

type CoreEventSvc struct {
	store eventStore
}

func NewCoreEventSvc(s eventStore) *CoreEventSvc {
	return &CoreEventSvc{
		store: s,
	}
}

func (cs CoreEventSvc) CreateEvent(ctx context.Context, t storage.Event) (int64, error) {
	return cs.store.CreateEvent(ctx, t)
}

func (cs CoreEventSvc) GetEvent(ctx context.Context, id int64) (storage.Event, error) {
	return cs.store.GetEvent(ctx, id)
}

func (cs CoreEventSvc) GetAllEvents(ctx context.Context) ([]storage.Event, error) {
	return cs.store.GetAllEvents(ctx)
}

func (cs CoreEventSvc) UpdateEvent(ctx context.Context, t storage.Event) error {
	return cs.store.UpdateEvent(ctx, t)
}

func (cs CoreEventSvc) DeleteEvent(ctx context.Context, id int64) error {
	return cs.store.DeleteEvent(ctx, id)
}

func (cs CoreEventSvc) SearchEvent(ctx context.Context, q string) ([]storage.Event, error) {
	return cs.store.SearchEvent(ctx, q)
}
