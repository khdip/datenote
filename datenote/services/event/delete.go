package event

import (
	"context"

	epb "datenote/gunk/v1/event"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *EventSvc) DeleteEvent(ctx context.Context, req *epb.DeleteEventRequest) (*epb.DeleteEventResponse, error) {
	id := req.ID

	err := s.core.DeleteEvent(context.Background(), id)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to delete event")
	}

	return &epb.DeleteEventResponse{}, nil
}
