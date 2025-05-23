package simple

type SayHello interface {
	Hello(name string) string
}

type HelloService struct {
	SayHello
}

func NewHelloService(sayHello SayHello) *HelloService {
	return &HelloService{SayHello: sayHello}
}

type SayHelloImpl struct {
	Name string
}

func NewSayHelloImpl() *SayHelloImpl {
	return &SayHelloImpl{}
}

func (s SayHelloImpl) Hello(name string) string {
	return "Hello " + name
}
