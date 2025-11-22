package main

import (
	"bistro/internal/dal"
	"bistro/internal/handler"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	flagDir := flag.String("dir", "data", "dir name")
	flagPort := flag.Int("port", 8000, "Port number")
	flag.Parse()

	initStorage(*flagDir)
	repo := dal.NewInventoryRepository(*flagDir)
	addr := fmt.Sprintf(":%d", *flagPort)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		inventoryHandler(w, r, repo)
	})

	http.ListenAndServe(addr, nil)
}

func inventoryHandler(w http.ResponseWriter, r *http.Request, repo *dal.InventoryRepository) {
	fmt.Println(r.URL.Path)
	url := strings.Split(r.URL.Path, "/")
	fmt.Println(url, url[0])
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
