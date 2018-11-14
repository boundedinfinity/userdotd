package system

import (
	"context"
	"log"
)

type System struct {
	ctx    context.Context
	logger *log.Logger
}

func NewSystem(ctx context.Context, logger *log.Logger) *System {
	return &System{
		ctx:    ctx,
		logger: logger,
	}
}
