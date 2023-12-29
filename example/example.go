package example

import (
	"fmt"
	"io"
)

type Application struct {
	FoodService FoodService
	// we will bind Fooer (interface) to MyFooer (concrete type)
	F Fooer
	// ignored field when passing "*" into injector
	// Ignored string `wire:"-"`

	// Binding value
	B Bar

	// Binding interface value
	R io.Reader

	// Binding from ContainS struct field
	S int
}

type Fooer interface {
	Foo() string
}

type MyFooer string

func NewMyFooer() MyFooer {
	return "interface binding"
}

func (i MyFooer) Foo() string {
	return "foo"
}

// Bar does not contain provider, we will inject value directly from injector
type Bar struct {
	V int
}

// ContainS provide its field
type ContainS struct {
	S int
}

func NewContainS() ContainS {
	return ContainS{S: 99}
}

type FoodService struct {
	orderRepository OrderRepository
}

func NewFoodService(repository OrderRepository) (FoodService, error) {
	return FoodService{
		orderRepository: repository,
	}, nil
}

type OrderRepository struct {
	db Connection
}

func NewOrderRepository(connection Connection) (OrderRepository, error) {
	return OrderRepository{
		db: connection,
	}, nil
}

type Connection struct {
	dbHost string
	dbName string
}

// NewConnection connection with cleanup function
func NewConnection(dbname string) (Connection, func(), error) {
	return Connection{
			dbHost: "mongodb://0.0.0.0",
			dbName: dbname,
		}, func() {
			fmt.Println("connection cleaned up")
		}, nil
}
