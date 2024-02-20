module helloworld

go 1.12

require (
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/google/wire v0.3.0
	go-common v1.45.2
	go.uber.org/automaxprocs v1.4.0
	google.golang.org/genproto v0.0.0-20220503193339-ba3ae3f07e29
	google.golang.org/grpc v1.46.0
)

replace go-common => git.bilibili.co/platform/go-common v1.45.2
