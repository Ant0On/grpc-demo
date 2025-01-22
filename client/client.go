package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Ant0On/grpc-demo/task-manager/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewTaskServiceClient(conn)

	// Dodawanie zadania
	addRes, err := client.AddTask(context.Background(), &proto.AddTaskRequest{
		Name:        "First Task",
		Description: "This is the first task",
	})
	if err != nil {
		log.Fatalf("could not add task: %v", err)
	}
	fmt.Println(addRes.GetName())

	// Pobieranie zadania
	getRes, err := client.GetTask(context.Background(), &proto.GetTaskRequest{
		Id: "1",
	})
	if err != nil {
		log.Fatalf("could not get task: %v", err)
	}
	fmt.Println("Task:", getRes.GetName(), getRes.GetDescription())
}
