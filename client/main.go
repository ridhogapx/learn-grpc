package main

import (
	"flag"
	pb "learn-grpc/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "Address gRPC")
)

type Movie struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Genre string `json:"genre"`
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic("failed to connect grpc server")
	}

	defer conn.Close()

	client := pb.NewMovieServiceClient(conn)

	r := gin.Default()

	r.POST("/movie", func(ctx *gin.Context) {
		var movie Movie

		err := ctx.ShouldBind(&movie)

		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err,
			})
			return
		}

		data := &pb.Movie{
			Title: movie.Title,
			Genre: movie.Genre,
		}

		res, err := client.CreateMovie(ctx, &pb.CreateMovieRequest{
			Movie: data,
		})

		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err,
			})
			return
		}

		ctx.JSON(201, gin.H{
			"movie": res.Movie,
		})
	})
}
