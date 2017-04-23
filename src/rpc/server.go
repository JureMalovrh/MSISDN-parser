package server

import (
	"msisdn/parser"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)


type Arg string

type MsisdnRpc int

func (t *MsisdnRpc) Parse(msisdn *Arg, data *parser.Msisdn) error {
	m, err := parser.ParseMsisdn(string(*msisdn))
	if err != nil {
		return err
	}

	*data = *m
	return nil
}

// StartRPCServer starts the rpc server on given port number
func StartRPCServer(port string) error {
	server := rpc.NewServer()
	err := server.Register(new(MsisdnRpc))
	if err != nil {
		return err
	}

	l, e := net.Listen("tcp", ":8222")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}