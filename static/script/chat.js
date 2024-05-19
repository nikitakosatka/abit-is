let socket;

function connectWebSocket() {
    socket = new WebSocket("wss://is-y25-website.onrender.com/ws");

    socket.onopen = function (event) {
        console.log("Connected to WebSocket server.");
    };

    socket.onclose = function (event) {
        console.log("Disconnected from WebSocket server.");
    };

    socket.onmessage = function (event) {
        let msgObj = JSON.parse(event.data);
        addMessageToChat(msgObj);
    };
}

function sendMessage() {
    let input = document.getElementById("msg");
    let email = localStorage.getItem("userEmail");
    if (input.value.trim() === "") {
        return;
    }
    if (email) {
        let msgObj = {email: email, message: input.value};
        socket.send(JSON.stringify(msgObj));
        input.value = '';
    } else {
        alert("Для отправки сообщения необходимо войти в систему.");
        input.value = '';
    }
}

function addMessageToChat(msgObj) {
    let chat = document.getElementById('chat');
    let newMessage = document.createElement('li');
    newMessage.innerText = msgObj.email + ": " + msgObj.message;
    chat.appendChild(newMessage);
    chat.scrollTop = chat.scrollHeight;
}

document.addEventListener('DOMContentLoaded', function () {
    connectWebSocket();
});

auth.onAuthStateChanged(user => {
    if (user) {
        user.getIdToken().then(idToken => {
            localStorage.setItem('idToken', idToken);
        });
        localStorage.setItem('userEmail', user.email);
    } else {
        localStorage.removeItem('user');
        localStorage.removeItem('idToken');
        localStorage.removeItem('userEmail');

        if (socket && socket.readyState === WebSocket.OPEN) {
            socket.close();
        }
    }
});
