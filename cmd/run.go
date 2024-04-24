package cmd

import (
	"202404/github.com/dubbo-go-sample/dubbogo/consumer"
	"202404/github.com/dubbo-go-sample/dubbogo/provider"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	dubboConfig "dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

func RunDubbo() {
	p := &provider.PcmtService{}
	c := &consumer.PCMTConsumer{}

	go withConfig(p, c)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	defer func() {
	}()
	return
}

func withConfig(p *provider.PcmtService, c *consumer.PCMTConsumer) {
	dubboConfig.SetProviderService(p)
	dubboConfig.SetConsumerService(c)
	if err := dubboConfig.Load(dubboConfig.WithPath("conf/dubbo.yaml")); err != nil {
		log.Fatal(err)
	}
	fmt.Println("run dubbo with config")
}
