let socket;

function connectWebSocket() {
    socket = new WebSocket("ws://localhost:8080/ws");

    socket.onopen = function(event) {
        console.log("Connected to WebSocket server.");
    };

    socket.onclose = function(event) {
        console.log("Disconnected from WebSocket server.");
    };

    socket.onmessage = function(event) {
        let chat = document.getElementById('chat');
        let newMessage = document.createElement('li');
        let msgObj = JSON.parse(event.data);
        newMessage.innerText = msgObj.email + ": " + msgObj.message;
        chat.appendChild(newMessage);
        chat.scrollTop = chat.scrollHeight;

        saveMessage(event.data);
    };
}

function sendMessage() {
    let input = document.getElementById("msg");
    let email = localStorage.getItem("userEmail");
    if (input.value.trim() === "") {
        return;
    }
    if (email) {
        let msgObj = { email: email, message: input.value };
        socket.send(JSON.stringify(msgObj));
        input.value = '';
    } else {
        alert("Для отправки сообщения необходимо войти в систему.");
        input.value = '';
    }
}

function isValidJSON(str) {
    try {
        JSON.parse(str);
    } catch (e) {
        return false;
    }
    return true;
}

function loadMessages() {
    let chat = document.getElementById('chat');
    let messages = JSON.parse(localStorage.getItem('chatMessages') || '[]');
    messages.forEach(function(message) {
        if (isValidJSON(message)) {
            let msgObj = JSON.parse(message);
            let newMessage = document.createElement('li');
            newMessage.innerText = msgObj.email + ": " + msgObj.message;
            chat.appendChild(newMessage);
            chat.scrollTop = chat.scrollHeight;
        }
    });
}

function saveMessage(message) {
    let messages = JSON.parse(localStorage.getItem('chatMessages') || '[]');
    messages.push(message);
    localStorage.setItem('chatMessages', JSON.stringify(messages));
}

document.addEventListener('DOMContentLoaded', function() {
    loadMessages();
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

        if (socket) {
            socket.close();
        }
    }
});