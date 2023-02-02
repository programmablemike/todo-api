// @generated by protoc-gen-connect-web v0.6.0 with parameter "target=ts"
// @generated from file todo/v1/todo.proto (package todo.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateTaskRequest, CreateTaskResponse, DeleteTaskRequest, DeleteTaskResponse, ListTasksRequest, ListTasksResponse } from "./todo_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service todo.v1.TodoService
 */
export const TodoService = {
  typeName: "todo.v1.TodoService",
  methods: {
    /**
     * @generated from rpc todo.v1.TodoService.CreateTask
     */
    createTask: {
      name: "CreateTask",
      I: CreateTaskRequest,
      O: CreateTaskResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc todo.v1.TodoService.DeleteTask
     */
    deleteTask: {
      name: "DeleteTask",
      I: DeleteTaskRequest,
      O: DeleteTaskResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc todo.v1.TodoService.ListTasks
     */
    listTasks: {
      name: "ListTasks",
      I: ListTasksRequest,
      O: ListTasksResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

