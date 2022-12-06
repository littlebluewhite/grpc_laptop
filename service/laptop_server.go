package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc_test/pb/message"
	"log"
)

// LaptopServer si the server that provides laptop services
type LaptopServer struct {
	Store LaptopStore
}

// NewLaptopServer returns a new LaptopServer
func NewLaptopServer() *LaptopServer {
	return &LaptopServer{}
}

func (server *LaptopServer) CreateLaptop(
	ctx context.Context,
	req *message.CreateLaptopRequest,
) (*message.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("receive a create-laptop request with id: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		// check if it's a valid UUID
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop ID: %v", err)
		}
		laptop.Id = id.String()
	}

	// save th laptop to in-memory store
	err := server.Store.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "cannot save laptop to the store: %v", err)
	}
	log.Printf("saved laptop with id: %s", laptop.Id)

	res := &message.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return res, nil
}
