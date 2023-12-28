package example

type Application struct {
	Name        string
	FoodService FoodService
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
}

func NewConnection() (Connection, error) {
	return Connection{
		dbHost: "mongodb://0.0.0.0",
	}, nil
}
