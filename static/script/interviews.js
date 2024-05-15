document.addEventListener('DOMContentLoaded', () => {
    const gallery = document.getElementById('gallery');
    const preloader = document.getElementById('preloader');
    const errorMessage = document.getElementById('error-message');

    fetchData();

    function initSwiper() {
        new Swiper('.swiper-container', {
            navigation: {
                nextEl: '.swiper-button-next',
                prevEl: '.swiper-button-prev',
            },
            pagination: {
                el: '.swiper-pagination',
                clickable: true,
            }
        });
    }

    function fetchData() {
        preloader.classList.remove('hidden');
        errorMessage.classList.add('hidden');

        const url = `https://is-y25-website.onrender.com/api/v1/interview`;

        fetch(url)
            .then((response) => {
                if (response.ok) {
                    return response.json();
                } else {
                    throw new Error('Network error');
                }
            })
            .then((data) => {
                render(data);
                preloader.classList.add('hidden');
            })
            .catch((error) => {
                console.error('Error fetching data:', error);
                errorMessage.textContent = '⚠️ Что-то пошло не так';
                errorMessage.classList.remove('hidden');
                preloader.classList.add('hidden');
            });
    }

    function render(data) {
        const fragment = document.createDocumentFragment();

        data.forEach(({ interview_id, title }) => {
            const slide = document.createElement('div');
            slide.className = 'swiper-slide';
            const text = document.createElement('h3');
            text.textContent = title;
            slide.onclick = () => {
                window.location.href = `https://is-y25-website.onrender.com/interview/${interview_id}`;
            };

            slide.appendChild(text);
            fragment.appendChild(slide);
        });

        gallery.innerHTML = '';
        gallery.appendChild(fragment);
        initSwiper();
    }
});
