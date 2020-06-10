package router

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

var (
    muxDispatcher = mux.NewRouter()
)

type muxRouter struct{}

func NewMuxRouter() RouteInterface {
    return &muxRouter{}
}

func (*muxRouter) GET( uri string, f func(w http.ResponseWriter, r *http.Request)){
    muxDispatcher.HandleFunc(uri,f).Methods("GET")
}

func (*muxRouter) POST( uri string, f func(w http.ResponseWriter, r *http.Request)){
    muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) SERVE( port string ){
    fmt.Printf("starting server on port %s\n",port)
    http.ListenAndServe(port,muxDispatcher)
}
