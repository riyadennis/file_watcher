package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/radovskyb/watcher"
)
type WatchServer struct{
	watcher *watcher.Watcher
	folder string
}
func Server(w http.ResponseWriter, res *http.Request,  _ httprouter.Params){
	w.Write([]byte("hello"))
}
