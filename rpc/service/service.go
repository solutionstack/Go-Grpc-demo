package main

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	pb "proto_test/proto"
	"sync"
	"syscall"
)

const (
	port = "0.0.0.0:9000"
)

// UserServer implements the generated protobuf server methods
type UserServer struct {
	pb.UnimplementedUserServer
	cache map[string]*pb.NewUser //in memory cache (or db store etc)
}

func NewUserService() *UserServer {
	return &UserServer{
		cache: make(map[string]*pb.NewUser),
	}
}

func (u *UserServer) CreateUSer(ctx context.Context, data *pb.PostParam) (*pb.NewUser, error) {

	new := &pb.NewUser{
		Id:      uuid.New().String(),
		Name:    data.GetName(),
		Age:     data.GetAge(),
		IsAdmin: data.GetIsAdmin(),
	}
	u.cache[new.Id] = new

	return new, nil

}
func (u *UserServer) GetUser(ctx context.Context, data *pb.GetParam) (*pb.NewUser, error) {

	user, ok := u.cache[data.GetId()]
	if !ok {
		return nil, errors.New("not found")
	}
	return user, nil
}

func (u *UserServer) GetUsers(ctx context.Context, empty *empty.Empty) (*pb.UserList, error) {
	v := make([]*pb.NewUser, 0, len(u.cache))

	for _, value := range u.cache {
		v = append(v, value)
	}

	return &pb.UserList{
		Users: v,
	}, nil
}

var wg = sync.WaitGroup{}

func main() {
	//create grpc server
	ls, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("could not bind address: %v", err)
	}

	svc := NewUserService()

	sv := grpc.NewServer()
	pb.RegisterUserServer(sv, svc)

	go startGRPCServe(sv, ls)
	log.Println("RPC Server listening on " + port)

	go shutdown(sv)

	wg.Add(1)
	wg.Wait()
}

func startGRPCServe(sv *grpc.Server, ls net.Listener) {
	err := sv.Serve(ls)
	if err != nil {
		log.Printf("could not start grp server: %v", err)
		wg.Done()
	}

}
func shutdown(sv *grpc.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	sv.Stop()
	log.Println("RPC Server shutdown")

	close(stop)
	wg.Done()

}
