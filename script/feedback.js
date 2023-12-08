document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('review-form');
    const reviewContainer = document.getElementById('review-container');

    form.addEventListener('submit', (event) => {
        event.preventDefault();

        const name = document.getElementById('review-name').value;
        const text = document.getElementById('review-text').value;

        const review = document.createElement('div');
        review.classList.add('review');

        const reviewName = document.createElement('p');
        reviewName.classList.add('review-name');
        reviewName.textContent = 'От: ' + name;

        const reviewText = document.createElement('p');
        reviewText.classList.add('review-text');
        reviewText.textContent = text;

        review.appendChild(reviewName);
        review.appendChild(reviewText);

        reviewContainer.appendChild(review);

        form.reset();

        localStorage.setItem('reviews', reviewContainer.innerHTML);
    });

    const savedReviews = localStorage.getItem('reviews');
    if (savedReviews) {
        reviewContainer.innerHTML = savedReviews;
    }
});