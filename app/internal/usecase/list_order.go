package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type ListOrdersOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type TListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func ListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *TListOrdersUseCase {
	return &TListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *TListOrdersUseCase) Execute() ([]ListOrdersOutputDTO, error) {
	orders, err := c.OrderRepository.List()
	if err != nil {
		return []ListOrdersOutputDTO{}, err
	}
	dto := convertOrdersToDTO(orders)
	return dto, nil
}

func convertOrdersToDTO(orders []entity.Order) []ListOrdersOutputDTO {
	output := make([]ListOrdersOutputDTO, len(orders))

	for i, order := range orders {
		output[i] = ListOrdersOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
	}

	return output
}
