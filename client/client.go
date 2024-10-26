package main

import (
	"book-grpc-users/users"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	c, err := grpc.NewClient(":9001", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("Error conneacting to port 9001 ", err)
	}

	defer c.Close()

	u := users.NewUsersClient(c)

	req := users.UserGetRequest{
		Email: "nasas@hotmail.com",
		Id:    "56",
	}

	res, err := u.GetUser(context.Background(), &req)

	if err != nil {
		log.Fatal("Error getting response from server: ", err)
	}
	log.Println("Response from server: ", res)
}
