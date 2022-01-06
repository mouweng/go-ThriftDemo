package main

import (
	"context"
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"

	"test/gen-go/echo"
)


type EchoServer struct {
}
func (e *EchoServer) Echo(ctx context.Context, req *echo.EchoReq) (*echo.EchoRes, error) {
    fmt.Printf("message from client: %v\n", req.GetMsg())

    res := &echo.EchoRes{
        Msg: "success",
    }

    return res, nil
}

func (e *EchoServer) Add(ctx context.Context, num1 *echo.Num, num2 *echo.Num) (*echo.Num, error) {
    fmt.Printf("This is Add!")
    num := &echo.Num {
        ID : num1.ID + num2.ID,
    }
    return num, nil
}



func main() {
    transport, err := thrift.NewTServerSocket(":9898")
    if err != nil {
        panic(err)
    }

    handler := &EchoServer{}
    processor := echo.NewEchoProcessor(handler)

    transportFactory := thrift.NewTBufferedTransportFactory(8192)
    protocolFactory := thrift.NewTCompactProtocolFactory()
    server := thrift.NewTSimpleServer4(
        processor,
        transport,
        transportFactory,
        protocolFactory,
    )

    if err := server.Serve(); err != nil {
        panic(err)
    }
}