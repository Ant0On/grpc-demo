package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Ant0On/grpc-demo/task-manager/proto"
	"google.golang.org/grpc"
)

type taskServer struct {
	proto.UnimplementedTaskServiceServer
	tasks map[string]proto.Task
}

func (s *taskServer) AddTask(_ context.Context, req *proto.AddTaskRequest) (*proto.AddTaskResponse, error) {
	id := fmt.Sprintf("%d", len(s.tasks)+1)
	task := proto.Task{
		Id:          id,
		Name:        req.GetName(),
		Description: req.GetDescription(),
	}
	s.tasks[id] = task
	return &proto.AddTaskResponse{Name: "Task added successfully!"}, nil
}

func (s *taskServer) GetTask(_ context.Context, req *proto.GetTaskRequest) (*proto.GetTaskResponse, error) {
	task, exists := s.tasks[req.GetId()]
	if !exists {
		return nil, fmt.Errorf("task not found")
	}
	return &proto.GetTaskResponse{
		Id:          task.GetId(),
		Name:        task.GetName(),
		Description: task.GetDescription(),
	}, nil
}

func (s *taskServer) ListTasks() (*proto.ListTaskResponse, error) {
	var tasksList []*proto.Task
	for _, task := range s.tasks {
		taskCopy := task
		tasksList = append(tasksList, &taskCopy)
	}
	return &proto.ListTaskResponse{Task: tasksList}, nil
}

func (s *taskServer) DeleteTask(_ context.Context, req *proto.DeleteTaskRequest) (*proto.DeleteTaskResponse, error) {
	_, exists := s.tasks[req.GetId()]
	if !exists {
		return nil, fmt.Errorf("task not found")
	}
	delete(s.tasks, req.GetId())
	return &proto.DeleteTaskResponse{Message: "Task deleted successfully!"}, nil
}

func main() {
	server := grpc.NewServer()
	taskService := &taskServer{tasks: make(map[string]proto.Task)}
	proto.RegisterTaskServiceServer(server, taskService)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("Server is listening on port 50051...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
