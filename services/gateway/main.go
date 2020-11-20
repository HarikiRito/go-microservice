package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
	pb "protobuf/_go/user"
)

func main() {
	r := gin.Default()

	conn := GrpcDialConnection("localhost:50051")
	//defer conn.Close()
	//defer cancel()
	client := pb.NewGreeterClient(conn)

	r.GET("/users", func(c *gin.Context) {
		reply, _ := client.SayHello(c, &pb.HelloRequest{Name: "wow"})
		log.Println(reply.GetMessage())
		c.JSON(http.StatusOK, gin.H{
			"ok": "bad",
		})

	})
	r.Run()
}

func GrpcDialConnection(address string) *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}
