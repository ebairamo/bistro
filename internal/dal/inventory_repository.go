package dal

import (
	"bistro/models"
	"encoding/json"
	"os"
)

type InventoryRepository struct {
	dataDir string
}

func NewInventoryRepository(dataDir string) *InventoryRepository {
	return &InventoryRepository{
		dataDir: dataDir,
	}
}

func (r *InventoryRepository) SaveItem(item models.InventoryItem) error {
	filepath := r.dataDir + "/inventory.json"

	// Шаг 1: Прочитать весь файл в []byte
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	// Шаг 2: JSON → Go структура (unmarshal)
	var items []models.InventoryItem
	err = json.Unmarshal(data, &items)
	if err != nil {
		return err
	}

	// Шаг 3: Добавить новый item в массив
	items = append(items, item)

	// Шаг 4: Go структура → JSON (marshal)
	newData, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return err
	}

	// Шаг 5: Записать обратно в файл
	err = os.WriteFile(filepath, newData, 0644)
	if err != nil {
		return err
	}

	return nil
}
