async function enviarTecla(keyToSend){

    const key = {
        key: keyToSend
    }

    try{
        const response = await fetch("/controle/send_key/", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body:JSON.stringify(key)
        })

        if (!response.ok){
            throw new Error(`Erro no servidor: ${response.status}`)
        }

        const resultado = await response.json()
        console.log("Sucesso. Resposta do servidor: ", resultado)
    } catch (error){
        console.log("Erro ao enviar requisicao", error)
    }

}



//______________ Botoes _________________________


//=============- D-pad -======================
const botaoUp = document.getElementById("botaoUp")
const botaoOk = document.getElementById("botaoOk")
const botaoLeft = document.getElementById("botaoLeft")
const botaoDown = document.getElementById("botaoDown")
const botaoRight = document.getElementById("botaoRight")
//=============================================

//============- Menu de navegacao -================
const botaoHome = document.getElementById("botaoHome")
const botaoBack = document.getElementById("botaoBack")
//=================================================

//============- Controle do volume -=============
const botaoVolUp = document.getElementById("volUp")
const botaoVolDown = document.getElementById("volDown")
//===============================================



//______________ Event-listeners ________________________


//=============- D-pad -================================
botaoUp.addEventListener("click", async () => enviarTecla("UP")
)

botaoOk.addEventListener("click", async () => enviarTecla("ENTER")
)

botaoLeft.addEventListener("click", async () => enviarTecla("LEFT")
)

botaoDown.addEventListener("click", async () => enviarTecla("DOWN")
)

botaoRight.addEventListener("click", async () => enviarTecla("RIGHT")
)
//=======================================================


//============- Menu de navegacao -======================
botaoHome.addEventListener("click", async () => enviarTecla("HOME")
)

botaoBack.addEventListener("click", async () => enviarTecla("BACK")
)
//========================================================


//============- Controle do volume -=============
botaoVolUp.addEventListener("click", async () => enviarTecla("VOLP")
)

botaoVolDown.addEventListener("click", async () => enviarTecla("VOLM")
)
//===============================================
