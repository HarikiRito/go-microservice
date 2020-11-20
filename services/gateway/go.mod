module gateway

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	google.golang.org/grpc v1.33.2
	protobuf v0.0.0-00010101000000-000000000000
)

replace protobuf => ../protobuf
