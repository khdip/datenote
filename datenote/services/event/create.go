package event

import (
	"context"

	"datenote/datenote/storage"
	epb "datenote/gunk/v1/event"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *EventSvc) CreateEvent(ctx context.Context, req *epb.CreateEventRequest) (*epb.CreateEventResponse, error) {
	event := storage.Event{
		ID:       req.GetEvent().ID,
		Name:     req.GetEvent().Name,
		Date:     req.GetEvent().Date,
		Info:     req.GetEvent().Info,
		Category: req.GetEvent().Category,
	}
	id, err := s.core.CreateEvent(context.Background(), event)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to create event")
	}

	return &epb.CreateEventResponse{
		ID: id,
	}, nil
}
