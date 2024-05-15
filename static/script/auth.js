const firebaseConfig = {
    apiKey: "AIzaSyAL7V5NwzBr2kQzTzibR75l4AleQCdA7Ds",
    authDomain: "abit-21c09.firebaseapp.com",
    projectId: "abit-21c09",
    storageBucket: "abit-21c09.appspot.com",
    messagingSenderId: "563877259748",
    appId: "1:563877259748:web:82eb92d7c3d0a2df677ff1"
};

const app = firebase.initializeApp(firebaseConfig);
const auth = firebase.auth();

auth.onAuthStateChanged(user => {
    if (user) {
        console.log("Пользователь вошел в систему:", user);
        document.getElementById('auth-buttons').style.display = 'none';
        document.getElementById('user-info').style.display = 'block';
        document.getElementById('user-email').textContent = user.email;
        user.getIdToken().then(idToken => {
            localStorage.setItem('user', JSON.stringify(user))
            localStorage.setItem('idToken', idToken);
        });
    } else {
        console.log("Пользователь вышел из системы");
        document.getElementById('auth-buttons').style.display = 'block';
        document.getElementById('user-info').style.display = 'none';
        localStorage.removeItem('user')
        localStorage.removeItem('idToken');
    }
});

document.addEventListener('DOMContentLoaded', function () {
    document.getElementById('login-btn').addEventListener('click', function () {
        window.location.href = '/login';
    });

    document.getElementById('signup-btn').addEventListener('click', function () {
        window.location.href = '/signup';
    });

    document.getElementById('logout-btn').addEventListener('click', function () {
        firebase.auth().signOut().then(() => {
            console.log("Выход выполнен успешно");
            localStorage.removeItem('user');
            document.location.reload(true);
        }).catch((error) => {
            console.error("Ошибка при выполнении выхода:", error);
        });
    });

    const user = localStorage.getItem('user');
    if (user) {
        document.getElementById('auth-buttons').style.display = 'none';
        document.getElementById('user-info').style.display = 'block';
        document.getElementById('user-email').textContent = JSON.parse(user).email;
    } else {
        document.getElementById('auth-buttons').style.display = 'block';
        document.getElementById('user-info').style.display = 'none';
    }
});