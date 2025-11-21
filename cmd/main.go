package main

import (
	"bistro/internal/handler"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	flagDir := flag.String("dir", "data", "dir name")
	flagPort := flag.Int("port", 8000, "Port number")
	flag.Parse()
	initStorage(*flagDir)
	addr := fmt.Sprintf(":%d", *flagPort)
	http.HandleFunc("/inventory", inventoryHandler)
	http.ListenAndServe(addr, nil)
}

func inventoryHandler(w http.ResponseWriter, r *http.Request) {

	handler.AddInventoryItem(w, r)
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
