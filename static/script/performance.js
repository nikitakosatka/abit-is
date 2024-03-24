(function () {
    const startTime = performance.now();

    window.addEventListener("load", function () {
        const endTime = performance.now();
        const loadTime = endTime - startTime;
        const loadTimeElement = document.getElementById("loadTime");
        loadTimeElement.textContent = Math.round(loadTime);
    });
})();
