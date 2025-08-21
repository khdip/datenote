package event

import (
	"context"

	"datenote/datenote/storage"
	epb "datenote/gunk/v1/event"
)

type eventCoreStore interface {
	CreateEvent(context.Context, storage.Event) (int64, error)
	GetEvent(context.Context, int64) (storage.Event, error)
	GetAllEvents(context.Context) ([]storage.Event, error)
	UpdateEvent(context.Context, storage.Event) error
	DeleteEvent(context.Context, int64) error
	SearchEvent(context.Context, string) ([]storage.Event, error)
}

type EventSvc struct {
	epb.UnimplementedEventServiceServer
	core eventCoreStore
}

func NewEventServer(c eventCoreStore) *EventSvc {
	return &EventSvc{
		core: c,
	}
}
