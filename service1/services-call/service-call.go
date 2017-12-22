package call

import (
	"fmt"
	"log"
	"net/rpc"

	"service1/contracts"
)

const port = 9002

func CreateClient() *rpc.Client {
	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal("dialing:", err)
	}
	return client
}

func PerformRequest(c *rpc.Client) contracts.Response {
	// lastName := string(contracts.Request)
	args := &contracts.Request{Name: "World"}
	var reply contracts.Response
	err := c.Call("LastNameHandler.GetLastName", args, &reply)

	if err != nil {
		log.Fatal("error:", err)
	}

	return reply
}
