package server

import (
//	"net/http"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"log"
)

var domain = "127.0.0.1"
var port = "20000"

type Remote struct {}

func Run(args []string) {
	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "-D":
			if i == len(args) {
				log.Fatalln("invalid domain")
			}
			domain = args[i + 1]
		case "-p":
			if i == len(args) {
				log.Fatalln("invalid port")
			}
			port = args[i + 1]
		}
	}

	log.Println("listening @ " + domain + ":" + port)

	listen, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer listen.Close()

	server := rpc.NewServer()
	err = server.Register(Remote{})
	if err != nil {
		log.Println(err.Error())
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
