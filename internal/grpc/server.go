package grpc

import ( 
	"flag" 
	pbfleet "github.com/lucassouzavieira/go-grpc-server/pkg/protobuf/schema/fleet"
	pbincident "github.com/lucassouzavieira/go-grpc-server/pkg/protobuf/schema/incident"
)


var (
	port = flag.Int("port", 50051, "The server port")
)

type fleetServer struct {
	pbfleet.UnimplementedFleetServiceServer
}

type incidentServer struct {
	pbincident.UnimplementedIncidentServiceServer
}

type grpcAppServer struct {
	fleet fleetServer
	incident incidentServer
}

