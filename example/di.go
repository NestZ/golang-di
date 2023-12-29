//go:build wireinject
// +build wireinject

package example

import (
	"github.com/google/wire"
	"io"
	"os"
)

// these lines will be copied into generated file
var serviceSet = wire.NewSet(NewFoodService)
var repositorySet = wire.NewSet(NewOrderRepository)
var connectionSet = wire.NewSet(NewConnection)

// struct provider, we can specify which struct fields to be injected
var app = wire.Struct(new(Application), "F", "FoodService", "B", "R", "S")

// injector
func InitializeApplication(dbname string) (*Application, func(), error) {
	// wire will generate code, then replace these lines will be replaced
	wire.Build(serviceSet, repositorySet, connectionSet, NewMyFooer, wire.Bind(new(Fooer), new(MyFooer)), wire.Value(Bar{V: 66}), wire.InterfaceValue(new(io.Reader), os.Stdin), NewContainS, wire.FieldsOf(new(ContainS), "S"), app)
	return &Application{}, nil, nil
}
