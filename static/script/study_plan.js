function showDescription(subjectName) {
    const description = document.getElementById(subjectName);
    if (description.style.display === 'block') {
        description.style.display = 'none';
    } else {
        const descriptions = document.getElementsByClassName('description');
        for (let i = 0; i < descriptions.length; i++) {
            descriptions[i].style.display = 'none';
        }
        description.style.display = 'block';
    }
}

function loadSubjects() {
    const user = localStorage.getItem('user');
    const isAdmin = user ? JSON.parse(user).email === 'admin@itmo.ru' : false;
    console.log(isAdmin);

    fetch('http://localhost:8080/api/v1/subject')
        .then(response => response.json())
        .then(subjects => {
            const curriculumGrid = document.getElementById('curriculumGrid');
            curriculumGrid.innerHTML = '';
            const semesters = {};

            subjects.forEach(subject => {
                if (!semesters[subject.semester]) {
                    semesters[subject.semester] = [];
                }
                semesters[subject.semester].push(subject);
            });

            Object.keys(semesters).sort().forEach(semester => {
                const semesterDiv = document.createElement('div');
                semesterDiv.className = (parseInt(semester) % 2 !== 0) ? 'semester odd' : 'semester even';
                const header = document.createElement('header');
                header.className = 'curriculum-header';
                header.textContent = `${semester} Семестр`;
                semesterDiv.appendChild(header);

                semesters[semester].forEach(subject => {
                    const subjectId = `${subject.name}-${semester}`;
                    const subjectDiv = document.createElement('div');
                    subjectDiv.className = 'subject';
                    subjectDiv.textContent = subject.name;
                    subjectDiv.onclick = () => showDescription(subjectId);
                    semesterDiv.appendChild(subjectDiv);

                    const descriptionDiv = document.createElement('div');
                    descriptionDiv.id = subjectId;
                    descriptionDiv.className = 'description';
                    descriptionDiv.style.display = 'none';
                    descriptionDiv.innerHTML = `<strong>${subject.name}: </strong>${subject.description}`;
                    semesterDiv.appendChild(descriptionDiv);
                });

                curriculumGrid.appendChild(semesterDiv);
                if (isAdmin) {
                    document.getElementById('adminControls').style.display = 'block';
                }
            });
        })
        .catch(error => console.error('Error loading subjects:', error));
}

function addSubject() {
    const name = document.getElementById('addName').value;
    let semester = document.getElementById('addSemester').value;
    const description = document.getElementById('addDescription').value;

    semester = parseInt(semester, 10);

    if (isNaN(semester)) {
        alert("Please enter a valid number for the semester.");
        return;
    }

    const subjectData = {
        name: name,
        semester: semester,
        description: description
    };

    console.log("Sending data:", subjectData);

    const idToken = localStorage.getItem('idToken');
    if (!idToken) {
        alert('User is not authenticated. Please log in.');
        return;
    }

    fetch('http://localhost:8080/api/v1/subject/', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${idToken}`
        },
        body: JSON.stringify(subjectData)
    })
        .then(response => {
            if (response.ok) {
                loadSubjects();
                document.getElementById('addSubjectForm').reset();
            } else {
                response.text().then(text => alert('Failed to add subject: ' + text));
            }
        })
        .catch(error => console.error('Error adding subject:', error));
}

function deleteSubject() {
    const name = document.getElementById('deleteName').value;
    let semester = document.getElementById('deleteSemester').value;

    semester = parseInt(semester, 10);

    if (isNaN(semester)) {
        alert("Please enter a valid number for the semester.");
        return;
    }

    const idToken = localStorage.getItem('idToken');
    if (!idToken) {
        alert('User is not authenticated. Please log in.');
        return;
    }

    fetch(`http://localhost:8080/api/v1/subject/${name}/${semester}`, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${idToken}`
        },
    })
        .then(response => {
            if (response.ok) {
                loadSubjects();
                document.getElementById('deleteSubjectForm').reset();
            } else {
                response.text().then(text => alert('Failed to delete subject: ' + text));
            }
        })
        .catch(error => console.error('Error deleting subject:', error));
}

function updateSubject() {
    const name = document.getElementById('updateName').value;
    let semester = document.getElementById('updateSemester').value;
    const description = document.getElementById('updateDescription').value;

    semester = parseInt(semester, 10);

    if (isNaN(semester)) {
        alert("Please enter a valid number for the semester.");
        return;
    }

    const subjectData = {
        name: name,
        semester: semester,
        description: description
    };

    const idToken = localStorage.getItem('idToken');
    if (!idToken) {
        alert('User is not authenticated. Please log in.');
        return;
    }

    console.log("Sending data:", subjectData);

    fetch('http://localhost:8080/api/v1/subject/', {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${idToken}`
        },
        body: JSON.stringify(subjectData)
    })
        .then(response => {
            if (response.ok) {
                loadSubjects();
                document.getElementById('updateSubjectForm').reset();
            } else {
                response.text().then(text => alert('Failed to update subject: ' + text));
            }
        })
        .catch(error => console.error('Error updating subject:', error));
}

document.addEventListener('DOMContentLoaded', loadSubjects);