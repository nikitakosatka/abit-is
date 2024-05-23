document.addEventListener('DOMContentLoaded', function () {
    const idToken = localStorage.getItem('idToken');
    const user = localStorage.getItem('user');
    const isAdmin = user ? JSON.parse(user).email === 'admin@itmo.ru' : false;
    document.getElementById('delete-button').addEventListener('click', function () {
        const interviewId = window.location.pathname.split('/').pop();
        fetch(`/api/v1/interview/${interviewId}`, {
            method: 'DELETE',
            headers: {
                'Authorization': `Bearer ${idToken}`
            }
        }).then(response => {
            if (response.ok) {
                window.location.href = '/interviews';
            } else {
                console.error('Failed to delete the interview');
            }
        });
    });

    if (isAdmin) {
        document.getElementById('adminControls').style.display = 'block';
    }

    document.getElementById('edit-button').addEventListener('click', function () {
        document.getElementById('interview-title').style.display = 'none';
        document.getElementById('interview-text').style.display = 'none';
        document.getElementById('edit-button').style.display = 'none';
        document.getElementById('delete-button').style.display = 'none';

        document.getElementById('edit-form').style.display = 'block';
    });

    document.getElementById('save-button').addEventListener('click', function () {
        const interviewId = window.location.pathname.split('/').pop();
        const updatedTitle = document.getElementById('edit-title').value;
        const updatedText = document.getElementById('edit-text').value;

        fetch(`/api/v1/interview/${interviewId}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${idToken}`
            },
            body: JSON.stringify({title: updatedTitle, text: updatedText})
        }).then(response => {
            if (response.ok) {
                document.getElementById('interview-title').innerText = updatedTitle;
                document.getElementById('interview-text').innerText = updatedText;

                document.getElementById('interview-title').style.display = 'block';
                document.getElementById('interview-text').style.display = 'block';
                document.getElementById('edit-button').style.display = 'block';
                document.getElementById('delete-button').style.display = 'block';

                document.getElementById('edit-form').style.display = 'none';
            } else {
                console.error('Failed to update the interview');
            }
        });
    });
});
