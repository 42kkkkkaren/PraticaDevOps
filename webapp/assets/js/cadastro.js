$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(event) {
    event.preventDefault();
    console.log("Handler for form submit invoked.");

    if ($('#senha').val() !== $('#confirmar-senha').val()) {
        alert("As senhas não coincidem");
        return;
    }

    const usuario = {
        nome: $('#nome').val(),
        email: $('#email').val(),
        nick: $('#nick').val(),
        senha: $('#senha').val()
    };

    console.log("Dados do usuário:", usuario);

    $.ajax({
        url: "http://localhost:5000/usuarios", // Certifique-se de que a URL está correta
        method: "POST",
        contentType: "application/json",
        data: JSON.stringify(usuario),
        success: function(data, textStatus, xhr) {
            console.log("Resposta da API: ", data);
            console.log("Status da resposta: ", textStatus);
            console.log("Código de status HTTP: ", xhr.status);
            alert("Usuário cadastrado com sucesso!");
            window.location.href = "/login";
        },
        error: function(xhr, textStatus, errorThrown) {
            console.error("Erro na resposta da API: ", xhr);
            console.log("Status da resposta: ", textStatus);
            console.log("Erro lançado: ", errorThrown);
            if (xhr.responseJSON && xhr.responseJSON.erro) {
                alert("Erro: " + xhr.responseJSON.erro);
            } else {
                alert("Erro ao cadastrar usuário");
            }
        }
    });
}
