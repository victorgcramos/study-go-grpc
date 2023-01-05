package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/victorgcramos/zi/pkg/user/pb"
)

type User struct {
	ID              uuid.UUID
	Username        string
	EncodedPassword []byte
}

type Server struct {
	pb.UserServiceServer

	db *ZiDB
}

func (s *Server) Version(_ context.Context, _ *pb.VersionRequest) (*pb.VersionResponse, error) {
	fmt.Println("[USER]: Version called.")
	return &pb.VersionResponse{
		Version: "v1",
		Token:   "MyToken",
	}, nil
}

func (s *Server) Register(_ context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	//	Prepare user to send to db
	u := User{
		ID:              uuid.New(),
		Username:        req.Username,
		EncodedPassword: []byte(req.Password),
	}
	// Save user on db
	err := s.db.New(u)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResponse{
		Verification: "Verified",
	}, nil
}

func NewService() *Server {
	return &Server{db: OpenDB()}
}
