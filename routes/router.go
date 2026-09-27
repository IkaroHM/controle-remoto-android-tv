package routes

import (
	"io/fs"
	"log"
	"net/http"

)

type Router struct{
	arquivosWeb fs.FS
}

func CriarRouter(arquivosWeb fs.FS) *Router{
	return &Router{arquivosWeb: arquivosWeb}
}

func (r *Router) StartRotas(){

	handlerControleGet, err := r.ControleGet()
	if err != nil{
		log.Fatal(err)
	}

	handlerHomeGet, err := r.HomeGet()
	if err != nil{
		log.Fatal(err)
	}

	handlerPaginaCerts, err := r.PaginaCertsGet()
	if err != nil{
		log.Fatal(err)
	}

	handlerPaginapairing, err := r.PaginaPairingGet()
	if err != nil{
		log.Fatal(err)
	}

	http.Handle("/home/", http.StripPrefix("/home", handlerHomeGet))
	

	http.Handle("/controle/", http.StripPrefix("/controle", handlerControleGet))

	http.HandleFunc("/controle/send_key/", r.CommandSend)

	
	http.Handle("/certificados/", http.StripPrefix("/certificados", handlerPaginaCerts))

	http.HandleFunc("/certificados/criar/", CriarCerts)

	
	http.Handle("/pairing/", http.StripPrefix("/pairing", handlerPaginapairing))

	http.HandleFunc("/pairing/iniciar/", r.IniciarPareamento)

	http.HandleFunc("/pairing/encerrar/", r.EncerrarPareamento)
}