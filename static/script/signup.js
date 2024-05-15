document.addEventListener('DOMContentLoaded', function () {
    document.getElementById('signup-form').addEventListener('submit', function (event) {
        event.preventDefault();
        const email = document.getElementById('signup-email').value;
        const password = document.getElementById('signup-password').value;
        const errorElement = document.getElementById('signup-error');

        fetch('api/v1/signup', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({email: email, password: password})
        })
            .then(response => {
                if (response.ok) {
                    return response.json();
                } else {
                    return response.json().then(json => Promise.reject(new Error(json.message || 'Ошибка при регистрации')));
                }
            })
            .then(data => {
                console.log('Пользователь успешно зарегистрирован, выполнение входа');
                return firebase.auth().signInWithEmailAndPassword(email, password);
            })
            .then(userCredential => {
                console.log("Успешный автоматический вход:", userCredential.user);
                window.location.href = '/';
            })
            .catch(error => {
                errorElement.textContent = error.message;
                errorElement.style.display = 'block';
            });
    });
});