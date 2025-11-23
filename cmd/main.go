package main

import (
	"bistro/internal/dal"
	"bistro/internal/handler"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

func main() {
	flagDir := flag.String("dir", "data", "dir name")
	flagPort := flag.Int("port", 8000, "Port number")
	flagHelp := flag.Bool("help", false, "help flag")
	flag.Parse()
	slog.Info("StartingBistro", "port", *flagPort, "dataDir", *flagDir)
	if *flagHelp {
		help()
		os.Exit(0)
	}
	initStorage(*flagDir)
	slog.Info("Storage initialized")
	repo := dal.NewInventoryRepository(*flagDir)
	addr := fmt.Sprintf(":%d", *flagPort)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		inventoryHandler(w, r, repo)
	})

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		slog.Error("Server failed", "error", err)
	}
}

func inventoryHandler(w http.ResponseWriter, r *http.Request, repo *dal.InventoryRepository) {
	url := strings.Split(r.URL.Path, "/")
	switch url[1] {
	case "inventory":
		if len(url) == 2 {
			switch r.Method {
			case http.MethodPost:
				handler.AddInventoryItem(w, r, repo)
			case http.MethodGet:
				handler.GetAllItems(w, r, repo)
			}
		}
		if len(url) == 3 {
			switch r.Method {
			case http.MethodGet:
				handler.GetItem(w, r, repo)
			case http.MethodPut:
				handler.UpdateInventoryItem(w, r, repo)
			case http.MethodDelete:
				handler.DeleteItem(w, r, repo)
			}
		}
	}

}

func initStorage(dir string) {
	err := os.Mkdir(dir, 0666)
	if err != nil {
		slog.Error("dir exists", "error", err)
	}
	inventoryDir := dir + "/inventory.json"
	_, err = os.Stat(inventoryDir)
	if err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(inventoryDir)
			if err != nil {
				slog.Error("file exist", "error", err)
			}
			file.WriteString("[]")
			file.Close()
		}
	}

}

func help() {
	fmt.Println(`$ ./bistro --help
Bistro Management System

Usage:
  hot-coffee [--port <N>] [--dir <S>] 
  hot-coffee --help

Options:
  --help       Show this screen.
  --port N     Port number.
  --dir S      Path to the data directory.`)
}
