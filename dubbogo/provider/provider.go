package provider

import (
	// pcmtpb "202404/github.com/dubbo-go-sample/dubbogo/proto/gen"
	"context"
	"dubbo.apache.org/dubbo-go/v3/config"
)

type PcmtService struct {
}

func (p *PcmtService) Reference() string {
    return "PcmtService"
}

// func (p *PcmtService) Ping(ctx context.Context, req *pcmtpb.Empty) (*pcmtpb.Pong, error) {
// 	return &pcmtpb.Pong{Data: "pong"}, nil
// }

func (p *PcmtService) Ping(ctx context.Context) (string, error) {
	return "test", nil
}

func (s *PcmtService) MethodMapper() map[string]string {
	return map[string]string{
		"Ping": "ping",
	}
}

func init() {
    config.SetProviderService(new(PcmtService))
}