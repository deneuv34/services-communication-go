package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"service1/contracts"
	call "service1/services-call"
)

const port = 9001

func StartServer() {
	fullname := new(FullNameHandler)
	rpc.Register(fullname)

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Service cannot start on port: %v", err))
	}

	http.Serve(l, http.HandlerFunc(httpHandler))
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	serverCodec := jsonrpc.NewServerCodec(&HttpConn{in: r.Body, out: w})
	err := rpc.ServeRequest(serverCodec)
	if err != nil {
		log.Printf("Error while serving JSON request: %v", err)
		http.Error(w, "Error while serving JSON request, detail has been logged.", 500)
	}

	return
}

type HttpConn struct {
	in  io.Reader
	out io.Writer
}

type FullNameHandler struct{}

func (h *FullNameHandler) FullNameCall(args *contracts.Request, reply *contracts.Response) error {
	string1 := args.Name[3:]
	clients := call.CreateClient()
	defer clients.Close()

	string2 := call.PerformRequest(clients)

	reply.FullName = "my name " + string1 + string2.FullName
	return nil
}

func (c *HttpConn) Read(p []byte) (n int, err error) {
	return c.in.Read(p)
}

func (c *HttpConn) Write(d []byte) (n int, err error) {
	return c.out.Write(d)
}

func (c *HttpConn) Close() error {
	return nil
}
