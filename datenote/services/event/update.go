package event

import (
	"context"

	"datenote/datenote/storage"
	epb "datenote/gunk/v1/event"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *EventSvc) UpdateEvent(ctx context.Context, req *epb.UpdateEventRequest) (*epb.UpdateEventResponse, error) {
	event := storage.Event{
		ID:       req.Event.ID,
		Name:     req.Event.Name,
		Date:     req.Event.Date,
		Info:     req.Event.Info,
		Category: req.Event.Category,
	}

	err := s.core.UpdateEvent(context.Background(), event)

	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to update event")
	}

	return &epb.UpdateEventResponse{}, nil
}
