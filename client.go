package main

import (
	"src/rpc/server"
	"src/msisdn/parser"
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	// client connecting to RPC service
	// and calling methods

	conn, err := net.Dial("tcp", "localhost:8222")

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := jsonrpc.NewClient(conn)

	var reply parser.Msisdn
	var args[11] server.Arg
	args[0] = "+386-41-23321"	
	args[1] = "+386-41-23321"	
	args[2] = "+386-41-23321"	
	args[3] = "+386-41-23321"	
	args[4] = "+386-41-23321"	
	args[5] = "+386-41-23321"	
	args[6] = "+386-41-23321"	
	args[7] = "+386-41-23321"	
	args[8] = "+386-41-23321"	
	args[9] = "+386-41-23321"	
	args[10] = "+386-41-23321"	
	for i := 0; i < 11; i++ {
		// passing Args to RPC call
		var arg = server.Arg
		arg = args[i];

		err = c.Call("MsisdnRpc.Parse", &arg, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		if(reply.ERR != "") {
			fmt.Printf("There was an error when parsing your input: %s, error: %s", arg, reply.ERR)
		} else {
			fmt.Printf("Number %s, parsed as:  %s\n", arg, parser.returnString(reply))
		}
	}
}
