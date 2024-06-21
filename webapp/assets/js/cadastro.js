$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(event) {
    event.preventDefault();
    console.log("Handler for form submit invoked.");

    if ($('#senha').val() !== $('#confirmar-senha').val()) {
        alert("As senhas não coincidem");
        return;
    }

    $.ajax({
        url: "http://localhost:5000/usuarios", // Use a URL correta da API
        method: "POST",
        contentType: "application/json", // Adicionando contentType para JSON
        data: JSON.stringify({
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val(),
        }),
        success: function(data) {
            alert("Usuário cadastrado com sucesso!");
            window.location.href = "/login";
        },
        error: function(xhr) {
            if (xhr.status === 409) { // Conflito de duplicação
                alert("Erro: " + xhr.responseJSON.erro);
            } else {
                alert("Erro ao cadastrar usuário");
            }
        }
    });
}
