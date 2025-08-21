package event

import (
	"context"

	epb "datenote/gunk/v1/event"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *EventSvc) GetEvent(ctx context.Context, req *epb.GetEventRequest) (*epb.GetEventResponse, error) {
	id := req.ID
	event, err := s.core.GetEvent(context.Background(), id)

	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to get event")
	}

	return &epb.GetEventResponse{
		Event: &epb.Event{
			ID:       event.ID,
			Name:     event.Name,
			Date:     event.Date,
			Info:     event.Info,
			Category: event.Category,
		},
	}, nil
}
