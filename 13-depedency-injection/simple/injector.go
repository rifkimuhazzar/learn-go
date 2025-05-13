//go:build wireinject
// +build wireinject

package simple

import (
	"io"
	"os"

	"github.com/google/wire"
)

func InitializeService(x bool) (*SimpleService, error) {
	wire.Build(NewSimpleService, NewSimpleRepository)
	return nil, nil
}

func InitializeDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabasePostgreSQL,
		NewDatabaseMongoDB,
		NewDatabaseRepository,
	)
	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializeFooBarService() *FooBarService {
    wire.Build(fooSet, barSet, NewFooBarService)
    return nil
}

// func InitializeHelloService() *HelloService {
//     wire.Build(NewHelloService, NewSayHelloImpl)
//     return nil
// }

var helloSet = wire.NewSet(
    NewSayHelloImpl,
    wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializeHelloService() *HelloService {
    wire.Build(helloSet, NewHelloService)
    return nil
}

func InitializeFooBar() *FooBar {
    wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "Foo", "Bar"))
    // wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "*"))
    return nil
}

func InitializeFooBarusingValue() *FooBar {
    wire.Build(wire.Value(&Foo{}), wire.Value(&Bar{}), wire.Struct(new(FooBar), "*"))
    return nil
}

func InitializeReader() io.Reader {
    wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
    return nil
}

func InitializeConfiguration() *Configuration {
    wire.Build(
        NewApplication,
        wire.FieldsOf(new(*Application), "Configuration"),
    )
    return nil
}

func InitializeConnection(name string) (*Connection, func()) {
    wire.Build(NewConnection, NewFile)
    return nil, nil
}
