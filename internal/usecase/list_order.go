package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type ListOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type TListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func ListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *TListOrderUseCase {
	return &TListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *TListOrderUseCase) Execute() ([]ListOrderOutputDTO, error) {
	orders, err := c.OrderRepository.List()
	if err != nil {
		return []ListOrderOutputDTO{}, err
	}
	dto := convertOrdersToDTO(orders)
	return dto, nil
}

func convertOrdersToDTO(orders []entity.Order) []ListOrderOutputDTO {
	output := make([]ListOrderOutputDTO, len(orders))

	for i, order := range orders {
		output[i] = ListOrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
	}

	return output
}
