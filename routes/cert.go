package routes

import (
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/IkaroHM/controle-remoto-android-tv/Tv"
	"github.com/IkaroHM/controle-remoto-android-tv/dtos"
	"github.com/drosocode/atvremote/pkg/common"
)

func (router *Router) PaginaCertsGet() (http.Handler, error){
	paginaCriarCerts, err := fs.Sub(router.arquivosWeb, "certificado")
	if err != nil{
		return nil, err
	}
	return http.FileServer(http.FS(paginaCriarCerts)), nil
}

func CriarCerts(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")

    if err := os.MkdirAll("certs", 0755); err != nil {
        http.Error(w, "Erro ao criar pasta", http.StatusInternalServerError)
        return
    }

	ipDaTv := Tv.GetIP()

	err := common.CreateCertificate("atvremote", []string{ipDaTv}, "certs/cert.pem", "certs/key.pem")
	if err != nil{
		log.Print("Erro ao criar certificados", err)
		http.Error(w, "Erro ao criar certificados", http.StatusInternalServerError)
		return
	}

	resposta, err := json.Marshal(dtos.Resposta{Mensagem: "Certificados criados com sucesso", Status: http.StatusOK}) 
	
	w.Write([]byte(resposta))

}