//go:build wireinject
// +build wireinject
package example

func InitializeApplication() (*Application, err) {
	wire.Build(NewFoodService, NewOrderRepository, db.NewConnection, wire.Struct(new(Application), "*"))
	return &Application{}, nil
}

