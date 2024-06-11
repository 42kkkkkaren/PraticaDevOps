$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(event) {
    event.preventDefault();
    console.log("Handler for form submit invoked.");

    if ($('#senha').val() != $('#confirmar-senha').val()) {
        alert("As senhas não coincidem");
        return;
    }

    $.ajax ({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val(),
        },
    }).done(function() {
        alert("Usuário cadastrado com sucesso!");
        window.location.href = "/login";
    }).fail(function() {
        alert("Erro ao cadastrar usuário");
    });
}