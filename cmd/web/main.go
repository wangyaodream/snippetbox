package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/wangyaodream/snippetbox/internal/envutil"
)

type application struct {
    errorLog *log.Logger
    infoLog  *log.Logger
}

type ServerConfig struct {
	Host string
	Port int
}

func main() {
	// Get configuration
	if err := godotenv.Load(".env"); err != nil {
		log.Print("INFO: 未找到环境文件！将使用默认配置")
	}
	cfg := &ServerConfig{
		Host: envutil.GetEnv("HOST", ""),
		Port: envutil.GetInt("PORT", 4000),
	}


	// flag
	addr := flag.String("addr", ":4000", "HTTP network address")
	runMode := flag.String("m", "env", "Run mode") // env | flag
	flag.Parse()

	// Leveled logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Lshortfile)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    app := &application{
        errorLog: errorLog,
        infoLog: infoLog,
    }
	// TODO 可以支持两种地址模式
	// log.Printf("Starting server on %s", *addr)
	hostStr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	if *runMode == "flag" {
		hostStr = *addr
	}

    srv := http.Server{
        Addr: hostStr,
        ErrorLog: errorLog,
        Handler: app.routes(),
    }


	infoLog.Printf("Starting server on :%d", cfg.Port)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
