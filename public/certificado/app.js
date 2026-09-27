const botaoCriarCerts = document.getElementById("botaoCriarCerts")
const mensagemStatus = document.getElementById("mensagemStatus")

async function criarCerts(){
    try{
        const response = await fetch("/certificados/criar/", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
        })

        if (!response.ok){
            throw new Error(`Erro no servidor: ${response.status}`)
        }

        const resultado = await response.json()
        console.log("Sucesso. Resposta do servidor: ", resultado)
        mensagemStatus.textContent = "Sucesso!"
    } catch (error){
        console.log("Erro ao enviar requisicao", error)
    }
}

botaoCriarCerts.addEventListener("click",  async () => criarCerts())