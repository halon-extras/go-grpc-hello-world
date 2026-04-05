module halon-extras/go-grpc/hello-world

go 1.26

require (
	google.golang.org/grpc v1.80.0
	google.golang.org/grpc/examples v0.0.0
)

replace google.golang.org/grpc/examples => /grpc-go/examples

require (
	golang.org/x/net v0.52.0 // indirect
	golang.org/x/sys v0.42.0 // indirect
	golang.org/x/text v0.35.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260401024825-9d38bb4040a9 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)
