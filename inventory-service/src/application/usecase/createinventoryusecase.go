package usecase

import (
	"example.com/inventory-service/src/application/interfaces"
	"example.com/inventory-service/src/application/logic"
	"example.com/inventory-service/src/domain/entity"
)

type CreateInventoryUseCase struct {
	InventoryRepository interfaces.IInventoryRepository
}

func (c *CreateInventoryUseCase) Execute(inventory entity.Inventory) *logic.Result[entity.Inventory] {
	createdInventoryResult := c.InventoryRepository.Create(inventory)
	if createdInventoryResult.IsFailure {
		return createdInventoryResult
	}
	createdInventory := createdInventoryResult.GetValue()
	return logic.OkResult(createdInventory)
}
