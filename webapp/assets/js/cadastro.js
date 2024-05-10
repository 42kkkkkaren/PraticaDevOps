$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(event) {
    event.preventDefault();
    console.log("Handler for form submit invoked.");

    var senha = $('#senha').val();
    var confirmarSenha = $('#confirmar-senha').val();

    if (senha !== confirmarSenha) {
        alert("As senhas não coincidem");
        return;
    }

    var userData = {
        nome: $('#nome').val(),
        email: $('#email').val(),
        nick: $('#nick').val(),
        senha: senha
    };

    console.log("Prepared user data:", userData);

    $.ajax({
        url: "/http:/localhost:8080/usuarios",
        method: "POST",
        contentType: "application/json",
        data: JSON.stringify(userData),
        success: function(response) {
            console.log('User registration successful:', response);
            alert('User registered successfully!');
        },
        error: function(xhr, status, error) {
            console.error('Registration failed:', status, error);
            alert('Failed to register user.');
        }
    });
}
