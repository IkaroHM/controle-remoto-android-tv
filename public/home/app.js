async function iniciarPareamento() {
  try {
    const response = await fetch("/pairing/iniciar/", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      throw new Error(`Erro no servidor: ${response.status}`);
    }

    const resultado = await response.json();
    console.log("Sucesso. Resposta do servidor: ", resultado);
  } catch (error) {
    console.log("Erro ao enviar requisicao", error);
  }
}

const botaoEmparelhar = document.getElementById("botaoEmparelhar").addEventListener("click", () => iniciarPareamento())
