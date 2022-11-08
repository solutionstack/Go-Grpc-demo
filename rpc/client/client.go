package rpc_client

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "proto_test/proto"
)

const (
	grpcAddress = "0.0.0.0:9000"
)

type UserRPCClient struct {
	client pb.UserClient //client holds a variable of the generated protobuf userClient definition
}

func NewUserClientManager() *UserRPCClient {
	return &UserRPCClient{}
}

func (urpc *UserRPCClient) Connect() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, errors.Wrap(err, "could not create grpc connection")
	}
	log.Printf("Connected to RPC server %v\n", grpcAddress)

	urpc.client = pb.NewUserClient(conn)
	return conn, nil
}

// GetClient return the userClient interface
func (urpc *UserRPCClient) GetClient() *pb.UserClient {

	return &urpc.client
}
