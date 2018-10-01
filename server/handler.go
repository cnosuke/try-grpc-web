package main

import (
	"context"
	"github.com/cnosuke/try-grpc-web/server/pb"
	"github.com/golang/protobuf/ptypes"
	"go.uber.org/zap"
	"time"
)

type handler struct {
	todo.TodoServiceServer

	logger *zap.SugaredLogger
}

func NewHandler(l *zap.SugaredLogger) *handler {
	return &handler{
		logger: l,
	}
}

func (h *handler) GetLatest(ctx context.Context, req *todo.GetLatestRequest) (res *todo.GetLatestResponse, err error) {
	h.logger.Infow("GetLatest", "request", req)
	ts, _ := ptypes.TimestampProto(time.Now())

	return &todo.GetLatestResponse{
		Todo: &todo.Todo{
			Timestamp: ts,
			Text: "hello",
		},
	}, nil
}
