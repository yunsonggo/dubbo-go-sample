package dubbopb

import hessian "github.com/apache/dubbo-go-hessian2"

type Pong struct {
	Data string `hessian:"data"`
}

func (p *Pong) name() string {
	return "response.pong"
}

func (p *Pong) JavaClassName() string {
	//return GetJavaClassName(p.name())
	return "com.pcmt.grpc.response.pong"
}

func init() {
	hessian.RegisterPOJOs(&Pong{})
}
