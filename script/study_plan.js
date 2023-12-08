function showDescription(subjectId) {
    const description = document.getElementById(subjectId);

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

