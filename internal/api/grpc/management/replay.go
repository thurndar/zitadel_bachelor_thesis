package management

import (
	"context"

	"github.com/caos/zitadel/pkg/grpc/management"
)

func (s *Server) ReplayEvents(ctx context.Context, in *management.ReplayEventsRequest) (*management.ReplayEventsResponse, error) {
	return &management.ReplayEventsResponse{}, s.es.ReplayEvents(ctx, in.Limit)
}
