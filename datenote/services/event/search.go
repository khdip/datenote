package event

import (
	"context"

	epb "datenote/gunk/v1/event"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *EventSvc) SearchEvent(ctx context.Context, req *epb.SearchEventRequest) (*epb.SearchEventResponse, error) {
	query := req.GetSearchEventQuery()
	events, err := s.core.SearchEvent(context.Background(), query)
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
		return nil, status.Error(codes.Internal, "Failed to search event")
	}

	return &epb.SearchEventResponse{
		SearchEventResult: e,
	}, nil
}
