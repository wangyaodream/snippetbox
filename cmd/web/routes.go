package main

import "net/http"

func (app *application) routes() *http.ServeMux {
    mux := http.NewServeMux()

    // 文件服务器
    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Handle("/static", http.StripPrefix("/static", fileServer))

    // 路由
    mux.HandleFunc("/", app.home)
    mux.HandleFunc("/snippet/view", app.snippetView)
    mux.HandleFunc("/snippet/create", app.snippetCreate)

    return mux
}
