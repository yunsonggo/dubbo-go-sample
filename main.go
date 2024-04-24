package main

// import "202404/github.com/dubbo-go-sample/cmd"
import (
	"log"

    "dubbo.apache.org/dubbo-go/v3/config"
    _ "dubbo.apache.org/dubbo-go/v3/imports"
    _ "202404/github.com/dubbo-go-sample/dubbogo/provider"
)


func main() {
	// cmd.RunDubbo()

	if err := config.Load(config.WithPath("/Users/admin/work/pcm/dubbo-go-sample/conf/dubbogo.yaml")); err != nil {
		log.Fatalf("Error loading Dubbo config: %v", err)
	}
	

	select {}
    // quit := make(chan os.Signal, 1)
    // signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
    // <-quit
    // config.Stop()
}
