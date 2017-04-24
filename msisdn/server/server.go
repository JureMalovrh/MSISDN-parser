package server

import (
	"msisdn/parser"
	"net"
	"log"
	"fmt"
	"net/rpc"
	"net/rpc/jsonrpc"
)


type Arg string
type MsisdnRpc int

func (t *MsisdnRpc) Parse(msisdn *Arg, data *parser.Msisdn) error {
	m := parser.ParseMsisdn(string(*msisdn))
	
	*data = m
	return nil
}

/* Starts server on port 8222 */
func StartRPCServer() error {
	fmt.Println("Starting server");
	server := rpc.NewServer()
	err := server.Register(new(MsisdnRpc))
	if err != nil {
		return err
	}

	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	
	l, e := net.Listen("tcp", ":8222")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	fmt.Println("Waiting for connectons on :8222");
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}