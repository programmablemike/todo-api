// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: todo/v1/todo.proto

package todov1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/programmablemike/todo-api/autogen/go/todo/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// TodoServiceName is the fully-qualified name of the TodoService service.
	TodoServiceName = "todo.v1.TodoService"
)

// TodoServiceClient is a client for the todo.v1.TodoService service.
type TodoServiceClient interface {
	CreateTask(context.Context, *connect_go.Request[v1.CreateTaskRequest]) (*connect_go.Response[v1.CreateTaskResponse], error)
	DeleteTask(context.Context, *connect_go.Request[v1.DeleteTaskRequest]) (*connect_go.Response[v1.DeleteTaskResponse], error)
	ListTasks(context.Context, *connect_go.Request[v1.ListTasksRequest]) (*connect_go.Response[v1.ListTasksResponse], error)
	// MarkTask as complete or incomplete
	MarkTask(context.Context, *connect_go.Request[v1.MarkTaskRequest]) (*connect_go.Response[v1.MarkTaskResponse], error)
}

// NewTodoServiceClient constructs a client for the todo.v1.TodoService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTodoServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) TodoServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &todoServiceClient{
		createTask: connect_go.NewClient[v1.CreateTaskRequest, v1.CreateTaskResponse](
			httpClient,
			baseURL+"/todo.v1.TodoService/CreateTask",
			opts...,
		),
		deleteTask: connect_go.NewClient[v1.DeleteTaskRequest, v1.DeleteTaskResponse](
			httpClient,
			baseURL+"/todo.v1.TodoService/DeleteTask",
			opts...,
		),
		listTasks: connect_go.NewClient[v1.ListTasksRequest, v1.ListTasksResponse](
			httpClient,
			baseURL+"/todo.v1.TodoService/ListTasks",
			opts...,
		),
		markTask: connect_go.NewClient[v1.MarkTaskRequest, v1.MarkTaskResponse](
			httpClient,
			baseURL+"/todo.v1.TodoService/MarkTask",
			opts...,
		),
	}
}

// todoServiceClient implements TodoServiceClient.
type todoServiceClient struct {
	createTask *connect_go.Client[v1.CreateTaskRequest, v1.CreateTaskResponse]
	deleteTask *connect_go.Client[v1.DeleteTaskRequest, v1.DeleteTaskResponse]
	listTasks  *connect_go.Client[v1.ListTasksRequest, v1.ListTasksResponse]
	markTask   *connect_go.Client[v1.MarkTaskRequest, v1.MarkTaskResponse]
}

// CreateTask calls todo.v1.TodoService.CreateTask.
func (c *todoServiceClient) CreateTask(ctx context.Context, req *connect_go.Request[v1.CreateTaskRequest]) (*connect_go.Response[v1.CreateTaskResponse], error) {
	return c.createTask.CallUnary(ctx, req)
}

// DeleteTask calls todo.v1.TodoService.DeleteTask.
func (c *todoServiceClient) DeleteTask(ctx context.Context, req *connect_go.Request[v1.DeleteTaskRequest]) (*connect_go.Response[v1.DeleteTaskResponse], error) {
	return c.deleteTask.CallUnary(ctx, req)
}

// ListTasks calls todo.v1.TodoService.ListTasks.
func (c *todoServiceClient) ListTasks(ctx context.Context, req *connect_go.Request[v1.ListTasksRequest]) (*connect_go.Response[v1.ListTasksResponse], error) {
	return c.listTasks.CallUnary(ctx, req)
}

// MarkTask calls todo.v1.TodoService.MarkTask.
func (c *todoServiceClient) MarkTask(ctx context.Context, req *connect_go.Request[v1.MarkTaskRequest]) (*connect_go.Response[v1.MarkTaskResponse], error) {
	return c.markTask.CallUnary(ctx, req)
}

// TodoServiceHandler is an implementation of the todo.v1.TodoService service.
type TodoServiceHandler interface {
	CreateTask(context.Context, *connect_go.Request[v1.CreateTaskRequest]) (*connect_go.Response[v1.CreateTaskResponse], error)
	DeleteTask(context.Context, *connect_go.Request[v1.DeleteTaskRequest]) (*connect_go.Response[v1.DeleteTaskResponse], error)
	ListTasks(context.Context, *connect_go.Request[v1.ListTasksRequest]) (*connect_go.Response[v1.ListTasksResponse], error)
	// MarkTask as complete or incomplete
	MarkTask(context.Context, *connect_go.Request[v1.MarkTaskRequest]) (*connect_go.Response[v1.MarkTaskResponse], error)
}

// NewTodoServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTodoServiceHandler(svc TodoServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/todo.v1.TodoService/CreateTask", connect_go.NewUnaryHandler(
		"/todo.v1.TodoService/CreateTask",
		svc.CreateTask,
		opts...,
	))
	mux.Handle("/todo.v1.TodoService/DeleteTask", connect_go.NewUnaryHandler(
		"/todo.v1.TodoService/DeleteTask",
		svc.DeleteTask,
		opts...,
	))
	mux.Handle("/todo.v1.TodoService/ListTasks", connect_go.NewUnaryHandler(
		"/todo.v1.TodoService/ListTasks",
		svc.ListTasks,
		opts...,
	))
	mux.Handle("/todo.v1.TodoService/MarkTask", connect_go.NewUnaryHandler(
		"/todo.v1.TodoService/MarkTask",
		svc.MarkTask,
		opts...,
	))
	return "/todo.v1.TodoService/", mux
}

// UnimplementedTodoServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTodoServiceHandler struct{}

func (UnimplementedTodoServiceHandler) CreateTask(context.Context, *connect_go.Request[v1.CreateTaskRequest]) (*connect_go.Response[v1.CreateTaskResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("todo.v1.TodoService.CreateTask is not implemented"))
}

func (UnimplementedTodoServiceHandler) DeleteTask(context.Context, *connect_go.Request[v1.DeleteTaskRequest]) (*connect_go.Response[v1.DeleteTaskResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("todo.v1.TodoService.DeleteTask is not implemented"))
}

func (UnimplementedTodoServiceHandler) ListTasks(context.Context, *connect_go.Request[v1.ListTasksRequest]) (*connect_go.Response[v1.ListTasksResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("todo.v1.TodoService.ListTasks is not implemented"))
}

func (UnimplementedTodoServiceHandler) MarkTask(context.Context, *connect_go.Request[v1.MarkTaskRequest]) (*connect_go.Response[v1.MarkTaskResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("todo.v1.TodoService.MarkTask is not implemented"))
}
