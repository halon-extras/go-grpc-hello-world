package main

// #cgo CFLAGS: -I/opt/halon/include
// #cgo LDFLAGS: -Wl,--unresolved-symbols=ignore-all
// #include <HalonMTA.h>
// #include <stdlib.h>
import "C"
import (
	"context"
	"log"
	"time"
	"unsafe"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

func main() {}

//export Halon_version
func Halon_version() C.int {
	return C.HALONMTA_PLUGIN_VERSION
}

//export greet
func greet(hhc *C.HalonHSLContext, args *C.HalonHSLArguments, ret *C.HalonHSLValue) {
	var argument = C.HalonMTA_hsl_argument_get(args, 0)
	if argument == nil {
		return
	}

	var _name *C.char
	if !C.HalonMTA_hsl_value_get(argument, C.HALONMTA_HSL_TYPE_STRING, unsafe.Pointer(&_name), nil) {
		return
	}
	name := C.GoString(_name)

	// Set up a connection to the server.
	conn, err := grpc.NewClient("[::1]:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	_output := C.CString(r.GetMessage())
	pointer := unsafe.Pointer(_output)
	defer C.free(pointer)
	C.HalonMTA_hsl_value_set(ret, C.HALONMTA_HSL_TYPE_STRING, pointer, 0)
}

//export Halon_hsl_register
func Halon_hsl_register(hhrc *C.HalonHSLRegisterContext) C.bool {
	C.HalonMTA_hsl_module_register_function(hhrc, C.CString("greet"), nil)
	return true
}
