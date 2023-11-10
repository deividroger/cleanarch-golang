package usecase

import "github.com/deividroger/cleanarch-golang/internal/entity"

type OrderListOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}
func (c *ListOrderUseCase) Execute() ([]OrderOutputDTO, error) {

	orders, err := c.OrderRepository.GetOrders()

	if err != nil {
		return []OrderOutputDTO{}, err
	}
	ordersResult := []OrderOutputDTO{}

	for _, order := range orders {
		ordersResult = append(ordersResult, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}
	return ordersResult, nil

}
