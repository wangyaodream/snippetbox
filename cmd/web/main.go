package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/wangyaodream/snippetbox/internal/envutil"
)

type ServerConfig struct{
    Host string
    Port int
}

func main() {
    // Get configuration
    if err := godotenv.Load(".env"); err != nil {
        log.Print("INFO: 未找到环境文件！")
    }
    cfg := &ServerConfig{
        Host: envutil.GetEnv("HOST", ""),
        Port: envutil.GetInt("PORT", 4000),
    }
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("Starting server on :%d", cfg.Port)
    hostStr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	err := http.ListenAndServe(hostStr, mux)
	log.Fatal(err)
}
