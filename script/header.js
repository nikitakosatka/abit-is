document.addEventListener("DOMContentLoaded", function () {
    const navMenu = document.getElementById("nav");
    const menuItems = navMenu.getElementsByTagName("a");

    for (let i = 0; i < menuItems.length; i++) {
        if (document.location.href === menuItems[i].href) {
            menuItems[i].className = "nav__link active";
        }
    }
})
