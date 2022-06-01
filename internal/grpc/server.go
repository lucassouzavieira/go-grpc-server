package grpc

import ( 
	"flag" 
	pbf "github.com/lucassouzavieira/go-grpc-server/pkg/protobuf/schema/fleet"
	pbi "github.com/lucassouzavieira/go-grpc-server/pkg/protobuf/schema/incident"
)


var (
	port = flag.Int("port", 50051, "The server port")
)

type fleetServer struct {
	pbf.UnimplementedFleetServiceServer
}

type incidentServer struct {
	pbi.UnimplementedIncidentServiceServer
}

type grpcAppServer struct {
	fleet fleetServer
	incident incidentServer
}

