document.addEventListener('DOMContentLoaded', () => {
    const loadDataBtn = document.getElementById('load-data');
    const gallery = document.getElementById('gallery');
    const preloader = document.getElementById('preloader');
    const errorMessage = document.getElementById('error-message');

    let firstLoad = true;

    loadDataBtn.addEventListener('click', () => {
        fetchData(firstLoad);
        firstLoad = !firstLoad;
    });

    function initSwiper() {
        new Swiper('.swiper-container', {
            navigation: {
                nextEl: '.swiper-button-next',
                prevEl: '.swiper-button-prev',
            },
            pagination: {
                el: '.swiper-pagination',
                clickable: true,
            },
        });
    }

    function fetchData(isFirstLoad) {
        preloader.classList.remove('hidden');
        errorMessage.classList.add('hidden');

        const url = isFirstLoad
            ? 'https://jsonplaceholder.typicode.com/photos?id_gte=100'
            : 'https://jsonplaceholder.typicode.com/photos?id_lte=200';

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

        data.forEach(({ title, thumbnailUrl, url }) => {
            const slide = document.createElement('div');
            slide.className = 'swiper-slide';

            const img = document.createElement('img');
            img.src = thumbnailUrl;
            img.alt = title;
            img.onclick = () => {
                window.open(url, '_blank');
            };

            const text = document.createElement('p');
            text.textContent = title;

            slide.appendChild(img);
            slide.appendChild(text);

            fragment.appendChild(slide);
        });

        gallery.innerHTML = '';
        gallery.appendChild(fragment);
        initSwiper();
    }
});