package main

import (
	"msisdn/server"
	"msisdn/parser"
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	go server.StartRPCServer()
	conn, err := net.Dial("tcp", "localhost:8222")

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := jsonrpc.NewClient(conn)

	var reply parser.Msisdn
	var args[11] server.Arg
	//ok, some known provider, some unknown
	//MNO: mobitel, CI: SI, CDC: 386, SN: 999666
	args[0] = "+(386)41999666"	
	//MNO: Unknown, CI: US, CDC: 1541, SN: 7543010
	args[1] = "+1-541-754-3010"	
	//MNO: Unknown, CI: UK, CDC: 44, SN: 79460353
	args[2] = "0044 7946 0353"
	//MNO: Unknown, CI: MW, CDC: 265, SN: 22266677
	args[3] = "+265-22266677"	
	//MNO: Skylink (Russia), CI: RU, CDC: 7, SN: 2023321
	args[4] = "+79012023321"
	//MNO: WIND, CI: IT, CDC: 39, SN: 111777
	args[5] = "0039329-111777"
	//MNO: Pelephone / Walla Mobile / YouPhone, CI: IL, CDC: 972, SN: 123456
	args[6] = "+972558123456"	
	// too long
	args[7] = "+385112332166666666666"
	// too short	
	args[8] = "+12312"
	// wrong country code	
	args[9] = "+999123123"
	// not only numbers
	args[10] = "444999echo"
	
	
	for i := 0; i < 11; i++ {
		// passing Args to RPC call
		var arg server.Arg
		arg = args[i]

		err = c.Call("MsisdnRpc.Parse", &arg, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		if(reply.ERR != "") {
			fmt.Printf("There was an error when parsing your input: %s, error: %s\n", arg, reply.ERR)
		} else {
			fmt.Printf("Number %s, parsed as:  %s\n", arg, parser.ReturnString(reply))
		}
	}
}
