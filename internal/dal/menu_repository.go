package dal

import (
	"bistro/models"
	"encoding/json"
	"errors"
	"os"
)

type MenuRepository struct {
	dataDir string
}

func NewMenuRepository(dataDir string) *MenuRepository {
	return &MenuRepository{
		dataDir: dataDir,
	}
}
func (r *MenuRepository) AddMenuItem(menuItem models.MenuItem) error {
	filepath := r.dataDir + "/menu.json"
	file, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	var menuItems []models.MenuItem
	err = json.Unmarshal(file, &menuItems)
	if err != nil {
		return err
	}
	menuItems = append(menuItems, menuItem)

	data, err := json.Marshal(menuItems)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath, data, 0666)
	if err != nil {
		return err
	}

	return nil
}

func (r *MenuRepository) GetMenuAllItems() ([]models.MenuItem, error) {

	filepath := r.dataDir + "/menu.json"
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var menuItems []models.MenuItem
	err = json.Unmarshal(file, &menuItems)
	if err != nil {
		return nil, err
	}
	return menuItems, nil
}

func (r *MenuRepository) GetMenuItem(id string) (models.MenuItem, error) {
	filepath := r.dataDir + "/menu.json"
	file, err := os.ReadFile(filepath)
	if err != nil {
		return models.MenuItem{}, err
	}
	menuItems := []models.MenuItem{}
	err = json.Unmarshal(file, &menuItems)
	if err != nil {
		return models.MenuItem{}, err
	}
	for _, item := range menuItems {
		if item.ID == id {
			return item, nil
		}
	}
	return models.MenuItem{}, errors.New("menuItem not found")
}
