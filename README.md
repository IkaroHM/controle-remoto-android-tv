# Controle-Remoto-Android-TV
Um controle remoto feito pra Android TVs, usando o protocolo V2

## Como funciona
O projeto se comunica com Android TVs através do protocolo V2, com http e tls, em Go.
O projeto usa a biblioteca [atvremote](https://github.com/drosoCode/atvremote)

 **Pareamento**: feito uma única vez, gera certificados TLS que autenticam
  o cliente com a TV.
- **Comandos**: enviados via TLS para a porta 6466 da TV.
- **Interface**: acessada pelo navegador em `http://localhost:8080`.

## Requisitos
- Go 1.21+
- Mi stick ou qualquer Android tv que use o protocolo V2 na mesma rede
-(Opcional) Termux, para rodar no telefone

## Como rodar

### Clonar e configurar
git clone https://github.com/IkaroHM/Controle-Remoto-Android-TV.git.
Crie um .env e coloque o ip da sua tv(Seguindo o .envexample).

### Rodar
Go run . ou go build.

### Criar certificados
Usando a interface, aperte no botao "Criar certs" e depois no botao "Criar certificados".
Se preferir, use a rota "/certificados/criar/", com metodo GET ou POST(corpo vazio).

### Parear
Pela interface, clique em "emparelhar", insirá o código que aparecer na TV e aperte em "Confirmar".
Se preferir, use a rota "/pairing/iniciar/" com metodo POST, em seguida, use a rota "/pairing/encerrar/" com metodo POST e com o corpo dessa maneira: {"codigo": "codigoDadoPelaTv"}.

### Controlar a tv
Pela interface, vá em "Controle".
Se preferir, use a rota "/controle/send_key/" com metodo POST e o corpo dessa maneira: {"key": "chave"}. Porém, atenção, pois a chave precisa ser valida. Olhe o pacote commom da [atvremote](https://github.com/drosoCode/atvremote)

## Estrutura do projeto

- controle/
- ├── main.go           # ponto de entrada
- ├── Tv/               # gerenciamento da conexão com a TV
- ├── routes/           # handlers HTTP
- ├── dtos/             # structs de request/response
- └── public/           # frontend embutido no binário

## Tecnologias

- Go (net/http, embed, crypto/tls).
- [atvremote](https://github.com/drosoCode/atvremote) (protocolo Android TV Remote v2).
- HTML/CSS/JS puro (sem framework).

## Licença

MIT.
