// package: todo
// file: todo.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class GetLatestRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetLatestRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetLatestRequest): GetLatestRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetLatestRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetLatestRequest;
  static deserializeBinaryFromReader(message: GetLatestRequest, reader: jspb.BinaryReader): GetLatestRequest;
}

export namespace GetLatestRequest {
  export type AsObject = {
  }
}

export class GetLatestResponse extends jspb.Message {
  hasTodo(): boolean;
  clearTodo(): void;
  getTodo(): Todo | undefined;
  setTodo(value?: Todo): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetLatestResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetLatestResponse): GetLatestResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetLatestResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetLatestResponse;
  static deserializeBinaryFromReader(message: GetLatestResponse, reader: jspb.BinaryReader): GetLatestResponse;
}

export namespace GetLatestResponse {
  export type AsObject = {
    todo?: Todo.AsObject,
  }
}

export class Todo extends jspb.Message {
  hasTimestamp(): boolean;
  clearTimestamp(): void;
  getTimestamp(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTimestamp(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getText(): string;
  setText(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Todo.AsObject;
  static toObject(includeInstance: boolean, msg: Todo): Todo.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Todo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Todo;
  static deserializeBinaryFromReader(message: Todo, reader: jspb.BinaryReader): Todo;
}

export namespace Todo {
  export type AsObject = {
    timestamp?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    text: string,
  }
}

