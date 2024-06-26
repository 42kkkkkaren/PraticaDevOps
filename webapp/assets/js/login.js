$(document).ready(function() {
    $('#login').on('submit', fazerLogin);
});

function fazerLogin(evento) {
    evento.preventDefault();

    const email = $('#email').val();
    const senha = $('#senha').val();

    if (!email || !senha) {
        alert("Por favor, preencha todos os campos.");
        return;
    }

    $.ajax({
        url: "http://localhost:5000/login", // Certifique-se de que a URL está correta
        method: "POST",
        contentType: "application/json",
        data: JSON.stringify({ email: email, senha: senha }),
        success: function(response) {
            console.log("Login bem-sucedido", response);
            alert("Login bem-sucedido!");

            // Armazena o token JWT no localStorage
            localStorage.setItem('token', response.token);

            // Redireciona para a página /home
            window.location.href = "/home";
        },
        error: function(xhr, textStatus, errorThrown) {
            console.error("Erro no login: ", xhr);
            console.log("Response status:", xhr.status);
            console.log("Response text:", xhr.responseText);
            console.log("Response JSON:", xhr.responseJSON);
            if (xhr.responseJSON && xhr.responseJSON.erro) {
                alert("Erro ao realizar login: " + xhr.responseJSON.erro);
            } else {
                alert("Erro ao realizar login: " + xhr.statusText);
            }
        }
    });
}
