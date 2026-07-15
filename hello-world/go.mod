module halon-extras/go-grpc/hello-world

go 1.26

require (
	google.golang.org/grpc v1.82.0
	google.golang.org/grpc/examples v0.0.0
)

replace google.golang.org/grpc/examples => /grpc-go/examples

require (
	golang.org/x/net v0.57.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.40.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260713224248-f5fc221cf8c4 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)
