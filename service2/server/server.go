package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"service2/contracts"
)

const port = 9002

type LastNameHandler struct{}

func (h *LastNameHandler) GetLastName(args *contracts.Request, reply *contracts.Response) error {
	reply.FullName = args.Name

	return nil
}

func StartServer() {
	lastNameService := &LastNameHandler{}

	rpc.Register(lastNameService)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))

	if err != nil {
		log.Fatal(fmt.Sprint("Cannot Listen on port: %v", err))
	}

	log.Printf("Server is starting on port: %v", port)

	http.Serve(l, nil)
}
