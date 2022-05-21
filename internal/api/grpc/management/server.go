package management

import (
	"google.golang.org/grpc"

	"github.com/caos/zitadel/internal/api/authz"
	"github.com/caos/zitadel/internal/api/grpc/server"
	"github.com/caos/zitadel/internal/command"
	"github.com/caos/zitadel/internal/config/systemdefaults"
	"github.com/caos/zitadel/internal/eventstore"
	"github.com/caos/zitadel/internal/query"
	"github.com/caos/zitadel/pkg/grpc/management"
)

const (
	mgmtName = "Management-API"
)

var _ management.ManagementServiceServer = (*Server)(nil)

type Server struct {
	management.UnimplementedManagementServiceServer
	command        *command.Commands
	query          *query.Queries
	systemDefaults systemdefaults.SystemDefaults
	assetAPIPrefix string
	es             *eventstore.Eventstore
}

func CreateServer(command *command.Commands, es *eventstore.Eventstore, query *query.Queries, sd systemdefaults.SystemDefaults, assetAPIPrefix string) *Server {
	return &Server{
		command:        command,
		query:          query,
		systemDefaults: sd,
		assetAPIPrefix: assetAPIPrefix,
		es:             es,
	}
}

func (s *Server) RegisterServer(grpcServer *grpc.Server) {
	management.RegisterManagementServiceServer(grpcServer, s)
}

func (s *Server) AppName() string {
	return mgmtName
}

func (s *Server) MethodPrefix() string {
	return management.ManagementService_MethodPrefix
}

func (s *Server) AuthMethods() authz.MethodMapping {
	return management.ManagementService_AuthMethods
}

func (s *Server) RegisterGateway() server.GatewayFunc {
	return management.RegisterManagementServiceHandlerFromEndpoint
}

func (s *Server) GatewayPathPrefix() string {
	return "/management/v1"
}
