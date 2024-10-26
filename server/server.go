package main

import (
	"book-grpc-users/users"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	//start the server

	l, err := net.Listen("tcp", ":9001")

	if err != nil {
		log.Fatal("Couldn't listen on port: 9001 ", err)
	}

	s := grpc.NewServer()

	st := users.UserService{}

	users.RegisterUsersServer(s, &st)

	err = s.Serve(l)

	if err != nil {
		log.Fatal("Couldn't serve the grpc server ", err)
	}
}
