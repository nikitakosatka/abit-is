document.addEventListener('DOMContentLoaded', function() {
    document.getElementById('login-form').addEventListener('submit', function(event) {
        event.preventDefault();
        const email = document.getElementById('login-email').value;
        const password = document.getElementById('login-password').value;
        const errorMessageElement = document.getElementById('error-message');

        firebase.auth().signInWithEmailAndPassword(email, password)
            .then((userCredential) => {
                console.log("Пользователь авторизован:", userCredential.user);
                window.location.href = '/';
            })
            .catch((error) => {
                const errorCode = error.code;
                const errorMessage = error.message;

                if (errorCode === 'auth/invalid-login-credentials') {
                    errorMessageElement.textContent = 'Неверный логин или пароль!';
                    errorMessageElement.style.display = 'block';
                } else {
                    errorMessageElement.textContent = errorMessage;
                    errorMessageElement.style.display = 'block';
                }
                console.error("Ошибка при входе:", errorMessage);
            });
    });
});