package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/appstacking/releasebus_githubapi/rpcservice"
	"github.com/xiaosongfu/asciibot"
)

func main() {
	// 注册 RPC 服务
	err := rpcservice.RegisterGithubApiService(new(rpcservice.GithubApiService))
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":12004")
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	fmt.Println(asciibot.Random())
	log.Printf("server info --> Version: %s, GitCommit: %s, BuildDate: %s, GoVersion: %s\n", Version, GitCommit, BuildDate, GoVersion)
	log.Println("server started.enjoy!")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go rpc.ServeConn(conn)
	}
}
