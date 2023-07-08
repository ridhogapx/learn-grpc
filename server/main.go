package main

import (
	"fmt"
	"learn-grpc/model"

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

func main() {
	// Create gRPC server
	// var (
	// 	port = flag.Int("port", 50051, "gRPC server port")
	// )

	// type server struct {
	// 	pb.UnimplementedMovieServiceServer
	// }

	fmt.Println("Triggering")
}
