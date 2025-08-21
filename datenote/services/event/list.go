package event

import (
	"context"

	epb "datenote/gunk/v1/event"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *EventSvc) GetAllEvents(ctx context.Context, req *epb.GetAllEventsRequest) (*epb.GetAllEventsResponse, error) {
	events, err := s.core.GetAllEvents(context.Background())
	var e []*epb.Event
	for _, event := range events {
		e = append(e, &epb.Event{
			ID:       event.ID,
			Name:     event.Name,
			Date:     event.Date,
			Info:     event.Info,
			Category: event.Category,
		})
	}

	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to get events")
	}

	return &epb.GetAllEventsResponse{
		Events: e,
	}, nil
}
