const formularioParear = document.getElementById("formularioPareamento")

formularioParear.addEventListener("submit", async (e) => {

    e.preventDefault()

    const mensagemStatus = document.getElementById("mensagemStatus")


    const codigoResponse = {
        "codigo": document.getElementById("codigo").value.toUpperCase()
    }

    try{
        const response = await fetch("/pairing/encerrar/", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body:JSON.stringify(codigoResponse)
        })

        if (!response.ok){
            throw new Error(`Erro no servidor: ${response.status}`)
        }

        const resultado = await response.json()
         mensagemStatus.textContent = "Sucesso!"
        console.log("Sucesso. Resposta do servidor: ", resultado)
    } catch (error){
        console.log("Erro ao enviar requisicao", error)
    }

})