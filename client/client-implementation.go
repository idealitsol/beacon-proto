package impl

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/idealitsol/beacon-proto/pbx"
	util "github.com/idealitsol/beacon-util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// GrpcClient defines all the clients
// type GrpcClient struct {
// 	Context   context.Context
// 	IAMClient *pbx.BeaconIamServiceClient
// 	SRMClient *pbx.BeaconSrmServiceClient
// 	OMAClient *pbx.BeaconOmaServiceClient
// }

// InvokeGrpcService is a client invocation call to a gRPC function on a server for resource
func InvokeGrpcService(ctx context.Context, host, service, function string, args ...interface{}) (interface{}, error) {
	var opts []grpc.DialOption
	opts = registerClientMiddlewares(opts)

	// dial client connection
	conn, err := grpc.DialContext(ctx, host, opts...)
	if err != nil {
		log.Fatalf("Failed to connect to application addr: %v", err)
	}
	defer conn.Close()

	var client interface{}
	var argsv []interface{}
	switch service {
	case "iam":
		client = pbx.NewBeaconIamServiceClient(conn)
		break
	case "oma":
		client = pbx.NewBeaconOmaServiceClient(conn)
		break
	case "srm":
		client = pbx.NewBeaconSrmServiceClient(conn)
		break
	default:
		return nil, fmt.Errorf("Invalid service provided")
	}

	headersIn, _ := metadata.FromIncomingContext(ctx)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	ctx = metadata.NewOutgoingContext(context.Background(), headersIn)
	defer cancel()

	argsv = append(argsv, ctx)
	argsv = append(argsv, args...)

	ret := util.InvokeFunction(client, function, argsv...)
	if !ret[1].IsNil() {
		return nil, ret[1].Interface().(error)
	}

	return ret[0].Interface(), nil
}
