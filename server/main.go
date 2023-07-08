package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"learn-grpc/model"
	pb "learn-grpc/proto"
	"net"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dsn string = "host=localhost user=root password=root dbname=learn-grpc port=5432 sslmode=disable"
var err error

func init() {
	DatabaseConnection()
}

func DatabaseConnection() {
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database!")
	}

	DB.AutoMigrate(model.Movie{})
	fmt.Println("Successfully connect to database!")
}

// Create gRPC server
var (
	port = flag.Int("port", 50051, "gRPC server port")
)

type server struct {
	pb.UnimplementedMovieServiceServer
}

func (*server) CreateMovie(ctx context.Context, req *pb.CreateMovieRequest) (*pb.CreateMovieResponse, error) {
	movie := req.GetMovie()
	movie.Id = uuid.New().String()

	data := model.Movie{
		ID:    movie.GetId(),
		Title: movie.GetTitle(),
		Genre: movie.GetGenre(),
	}

	res := DB.Create(&data)

	if res.RowsAffected == 0 {
		return nil, errors.New("failed to add movie records")
	}

	return &pb.CreateMovieResponse{
		Movie: &pb.Movie{
			Id:    movie.GetId(),
			Title: movie.GetTitle(),
			Genre: movie.GetGenre(),
		},
	}, nil
}

func main() {
	listen, errList := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if errList != nil {
		panic("Failed to listen gRPC server!")
	}

	s := grpc.NewServer()

	pb.RegisterMovieServiceServer(s, &server{})
	fmt.Printf("Server is listening on port %v", *port)

	if errRpc := s.Serve(listen); errRpc != nil {
		panic("Failed to start gRPC Server!")
	}

}
