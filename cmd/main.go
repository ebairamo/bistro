package main

import (
	"bistro/internal/dal"
	"bistro/internal/handler"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	flagDir := flag.String("dir", "data", "dir name")
	flagPort := flag.Int("port", 8000, "Port number")
	flag.Parse()

	initStorage(*flagDir)
	repo := dal.NewInventoryRepository(*flagDir)

	// Создаём Gin роутер
	r := gin.Default()

	// Роуты
	r.POST("/inventory", func(c *gin.Context) {
		handler.AddInventoryItem(c, repo)
	})
	r.GET("/inventory", func(c *gin.Context) {
		handler.GetAllItems(c, repo)
	})
	r.GET("/inventory/:id", func(c *gin.Context) {
		handler.GetItemByID(c, repo)
	})

	// Запуск
	addr := fmt.Sprintf(":%d", *flagPort)
	r.Run(addr)
}

func inventoryHandler(w http.ResponseWriter, r *http.Request, repo *dal.InventoryRepository) {
	fmt.Println(r.URL.Path)
	switch r.Method {
	case http.MethodPost:
		handler.AddInventoryItem(w, r, repo)
	case http.MethodGet:
		handler.GetAllItems(w, r, repo)
	}

}

func initStorage(dir string) {
	err := os.Mkdir(dir, 0666)
	if err != nil {
		fmt.Println(err)
	}
	inventoryDir := dir + "/inventory.json"
	_, err = os.Stat(inventoryDir)
	if err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(inventoryDir)
			if err != nil {
				fmt.Println(err)
			}
			file.WriteString("[]")
			file.Close()
		}
	}

}
