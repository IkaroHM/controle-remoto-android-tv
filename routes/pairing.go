package routes

import (
	"encoding/json"
	"io"
	"io/fs"
	"log"
	"net/http"

	"github.com/IkaroHM/controle-remoto-android-tv/Tv"
	"github.com/IkaroHM/controle-remoto-android-tv/dtos"
)

func (router *Router) PaginaPairingGet() (http.Handler, error){
	home, err := fs.Sub(router.arquivosWeb, "pairing")
	if err != nil{
		return nil, err
	}
	return  http.FileServer(http.FS(home)), nil
}

func (router *Router) IniciarPareamento(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost{
		http.Error(w, "Metodo invalido, tente POST", http.StatusBadRequest)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")

	_, err := Tv.IniciarPareamento()
	if err != nil{
		http.Error(w, "Erro ao iniciar pareamento", http.StatusInternalServerError)
		return
	}

	resposta, err := json.Marshal(dtos.Resposta{
		Mensagem: "Pareamento iniciado",
		Status: http.StatusOK,
	})
	if err != nil{
		http.Error(w, "Erro na resposta", http.StatusInternalServerError)
		return
	}

	w.Write(resposta)
}

func (router *Router) EncerrarPareamento(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost{
		http.Error(w, "Metodo invalido, tente POST", http.StatusBadRequest)
		return
	}
	
	w.Header().Set("Content-Type", "application.Json")

	b, err := io.ReadAll(r.Body)
	if err != nil{
		http.Error(w, "Erro ao ler o corpo", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var req dtos.RequisicaoPairing
	err = json.Unmarshal(b, &req)
	if err != nil{
		http.Error(w, "Json invalido", http.StatusBadRequest)
		return
	}

	err = Tv.FinalizarPareamento(req.Codigo)
	if err != nil{
		log.Print(err)
		http.Error(w, "Erro ao encerrar pareamento", http.StatusInternalServerError)
		return
	}

	resposta, err := json.Marshal(dtos.Resposta{
		Mensagem: "Pareamento encerrado com sucesso",
		Status: http.StatusOK,
	})
	if err != nil{
		http.Error(w, "Erro na resposta", http.StatusInternalServerError)
		return
	}

	w.Write(resposta)
}