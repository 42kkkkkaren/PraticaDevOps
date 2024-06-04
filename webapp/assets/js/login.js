$('login').on('submit', fazerLogin);

function fazerLogin(evento) {
    evento.preventDefault();

    $.ajax({
        url: 'http://localhost:8080/login',
        method: 'POST',
        data: {
            email: $('#email').val(),
            senha: $('#senha').val()
        }
    }).done(function(){
        window.location = 'http://localhost:8080/home';
    }).fail(function(){
        alert('Usuário ou senha inválidos');
    });
}