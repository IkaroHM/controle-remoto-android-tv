package routes

import(
	"net/http"
	"io/fs"
)

func (router *Router) HomeGet() (http.Handler, error){
	home, err := fs.Sub(router.arquivosWeb, "home")
	if err != nil{
		return nil, err
	}
	return  http.FileServer(http.FS(home)), nil
}