package service

import (
	"bistro/internal/dal"
	"bistro/models"
	"errors"
)

func SaveItem(item models.InventoryItem, repo *dal.InventoryRepository) error {
	if item.IngredientID == "" {
		return errors.New("ingredient_id cannot be empty")
	}
	if item.Name == "" {
		return errors.New("item name  cannot be empty")
	}
	if item.Quantity <= 0 {
		return errors.New("quantity can not be <= 0")
	}
	if item.Unit == "" {
		return errors.New("Unit cannot be empty")
	}
	err := repo.SaveItem(item)
	if err != nil {
		return err
	}
	return nil
}

func GetAllItems(repo *dal.InventoryRepository) ([]models.InventoryItem, error) {
	return repo.GetAllItems()
}
