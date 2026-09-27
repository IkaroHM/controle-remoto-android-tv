package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"strings"

	"github.com/IkaroHM/controle-remoto-android-tv/dtos"
	"github.com/IkaroHM/controle-remoto-android-tv/Tv"
	"github.com/drosocode/atvremote/pkg/common"
)

func (router *Router) ControleGet() (http.Handler, error){
	paginaControle, err := fs.Sub(router.arquivosWeb, "controle")
	if err != nil{
		return nil, err
	}
	return http.FileServer(http.FS(paginaControle)), nil
}

func (router *Router) CommandSend(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost{
		http.Error(w, "Metodo invalido, tente POST", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	b, err := io.ReadAll(r.Body)
	if err != nil{
		http.Error(w, "Erro ao ler o corpo", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var req dtos.Requisicao
	err = json.Unmarshal(b, &req)
	if err != nil{
		http.Error(w, "Json invalido", http.StatusBadRequest)
		return
	}

	key := strings.ToUpper(req.Key)

	comandos := common.ParseKeys(key)
	if comandos == nil{
		http.Error(w, "Chave invalida", http.StatusBadRequest)
		return
	}

	cmd := Tv.GetComando()
	if cmd == nil{
		http.Error(w, "Tv nao esta conectada. Acesse /parear/ para se conectar", http.StatusServiceUnavailable)
		return
	}

	for _, c := range comandos{
		err := Tv.EnviarTecla(c)
		if err != nil{
			fmt.Println(err)
			http.Error(w, "Erro ao mandar comando para a tv", http.StatusInternalServerError)
			return
		}
	}

	var resposta dtos.Resposta
	resposta.Mensagem = "Comando enviado com sucesso"
	resposta.Status = http.StatusOK
	respostaJson, err := json.Marshal(resposta)
	if err != nil{
		http.Error(w, "Erro na resposta", http.StatusInternalServerError)
		return
	}

	w.Write(respostaJson)

}
