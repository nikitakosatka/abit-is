document.addEventListener('DOMContentLoaded', () => {
    const user = localStorage.getItem('user');
    const isAdmin = user ? JSON.parse(user).email === 'admin@itmo.ru' : false;
    const gallery = document.getElementById('gallery');
    const preloader = document.getElementById('preloader');
    const errorMessage = document.getElementById('error-message');

    let swiperInstance;

    fetchData();

    if (isAdmin) {
        document.getElementById('adminControls').style.display = 'block';
    }

    function initSwiper() {
        swiperInstance = new Swiper('.swiper-container', {
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

    function addInterview() {
        const title = document.getElementById('addTitle').value;
        let text = document.getElementById('addText').value;

        const data = {
            title: title,
            text: text,
        };

        console.log("Sending data:", data);

        const idToken = localStorage.getItem('idToken');
        if (!idToken) {
            alert('User is not authenticated. Please log in.');
            return;
        }

        fetch('https://is-y25-website.onrender.com/api/v1/interview/', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${idToken}`
            },
            body: JSON.stringify(data)
        })
            .then(response => {
                if (response.ok) {
                    // Check if response has content before attempting to parse JSON
                    return response.text().then(text => text ? JSON.parse(text) : {});
                } else {
                    return response.text().then(text => {
                        throw new Error('Failed to add interview: ' + text);
                    });
                }
            })
            .then(newInterview => {
                if (newInterview && newInterview.interview_id) {
                    addNewSlide(newInterview);
                    document.getElementById('addSubjectForm').reset();
                } else {
                    console.error('Invalid interview data:', newInterview);
                }
            })
            .catch(error => {
                console.error('Error adding interview:', error);
                alert(error.message);
            });
    }

    function addNewSlide({ interview_id, title }) {
        const slide = document.createElement('div');
        slide.className = 'swiper-slide';
        const text = document.createElement('h3');
        text.textContent = title;
        slide.onclick = () => {
            window.location.href = `https://is-y25-website.onrender.com/interview/${interview_id}`;
        };

        slide.appendChild(text);

        const swiperWrapper = gallery.querySelector('.swiper-wrapper') || gallery;
        if (swiperWrapper) {
            swiperWrapper.appendChild(slide);

            if (swiperInstance) {
                swiperInstance.update(); // Update swiper to recognize the new slide
            } else {
                console.error('Swiper instance not initialized');
            }
        } else {
            console.error('.swiper-wrapper or #gallery not found');
        }
    }

    window.addInterview = addInterview;
});