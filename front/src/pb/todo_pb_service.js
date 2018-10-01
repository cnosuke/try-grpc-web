// package: todo
// file: todo.proto

var todo_pb = require("./todo_pb");
var grpc = require("grpc-web-client").grpc;

var TodoService = (function () {
  function TodoService() {}
  TodoService.serviceName = "todo.TodoService";
  return TodoService;
}());

TodoService.GetLatest = {
  methodName: "GetLatest",
  service: TodoService,
  requestStream: false,
  responseStream: false,
  requestType: todo_pb.GetLatestRequest,
  responseType: todo_pb.GetLatestResponse
};

exports.TodoService = TodoService;

function TodoServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

TodoServiceClient.prototype.getLatest = function getLatest(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(TodoService.GetLatest, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

exports.TodoServiceClient = TodoServiceClient;

