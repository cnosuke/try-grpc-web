// package: todo
// file: todo.proto

import * as todo_pb from "./todo_pb";
import {grpc} from "grpc-web-client";

type TodoServiceGetLatest = {
  readonly methodName: string;
  readonly service: typeof TodoService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof todo_pb.GetLatestRequest;
  readonly responseType: typeof todo_pb.GetLatestResponse;
};

export class TodoService {
  static readonly serviceName: string;
  static readonly GetLatest: TodoServiceGetLatest;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }
export type ServiceClientOptions = { transport: grpc.TransportConstructor; debug?: boolean }

interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: () => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}

export class TodoServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: ServiceClientOptions);
  getLatest(
    requestMessage: todo_pb.GetLatestRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: todo_pb.GetLatestResponse|null) => void
  ): void;
  getLatest(
    requestMessage: todo_pb.GetLatestRequest,
    callback: (error: ServiceError, responseMessage: todo_pb.GetLatestResponse|null) => void
  ): void;
}

