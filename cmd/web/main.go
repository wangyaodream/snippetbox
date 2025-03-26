package main

import (
	"database/sql"
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
    DB_DSN string
}

func main() {
	// Get configuration
	if err := godotenv.Load(".env"); err != nil {
		log.Print("INFO: 未找到环境文件！将使用默认配置")
	}
	cfg := &ServerConfig{
		Host: envutil.GetEnv("HOST", ""),
		Port: envutil.GetInt("PORT", 4000),
        DB_DSN: envutil.GetEnv("DB_DSN", "web:pass@/snippetbox?parseTime=true"),
	}


	// flag
	addr := flag.String("addr", ":4000", "HTTP network address")
	runMode := flag.String("m", "env", "Run mode") // env | flag
	flag.Parse()

	// Leveled logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Lshortfile)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    db, err := openDB(cfg.DB_DSN)
    if err != nil {
        errorLog.Fatal(err)
    }

    defer db.Close()

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
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}
