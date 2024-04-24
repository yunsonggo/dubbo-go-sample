package provider

import (
	pcmtpb "202404/github.com/dubbo-go-sample/dubbogo/proto/gen"
	"context"
)

type PcmtService struct {
}

func (p *PcmtService) Ping(ctx context.Context, req *pcmtpb.Empty) (*pcmtpb.Pong, error) {
	return &pcmtpb.Pong{Data: "pong"}, nil
}
