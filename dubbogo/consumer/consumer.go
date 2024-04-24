package consumer

import (
	"202404/github.com/dubbo-go-sample/dubbogo/dubbopb"
	"context"
)

type PCMTConsumer struct {
	Ping func(ctx context.Context) (res *dubbopb.Pong, err error) `dubbo:"ping"`
}
