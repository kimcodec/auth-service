package main

import (
	"context"
	"fmt"
	user_api "github.com/kimcodec/microservices/auth-service/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

type server struct {
	user_api.UnimplementedUserApiV1Server
}

func (s *server) Get(cxt context.Context, req *user_api.GetRequest) (*user_api.GetResponse, error) {
	log.Println("Id :", req.GetId())
	return nil, nil
}

func (s *server) Create(ctx context.Context, req *user_api.CreateRequest) (*user_api.CreateResponse, error) {
	log.Printf("Name: %s, Password: %s, Confirm_password: %s, Email: %s, Role: %s",
		req.GetName(), req.GetPassword(), req.GetPasswordConfirm(), req.GetEmail(), req.GetRole())
	return nil, nil
}

func (s *server) Update(ctx context.Context, req *user_api.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("Id: %d, Name: %s, Email: %s", req.GetId(), req.GetName(), req.GetName())
	return nil, nil
}

func (s *server) Delete(ctx context.Context, req *user_api.DeleteRequest) (*emptypb.Empty, error) {
	log.Println("Id: ", req.GetId())
	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}

	s := grpc.NewServer()
	reflection.Register(s)
	user_api.RegisterUserApiV1Server(s, &server{})

	log.Printf("server listening at %d port", 50051)

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to server: ", err.Error())
	}
}
