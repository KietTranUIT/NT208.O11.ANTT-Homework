function changeColor() {
    var circle = document.getElementById("circle");
    var randomColor = getRandomColor();
    circle.style.backgroundColor = randomColor;
}

function getRandomColor() {
    var letters = "0123456789ABCDEF";
    var color = "#";
    for (var i = 0; i < 6; i++) {
        color += letters[Math.floor(Math.random() * 16)];
    }
    return color;
}

function resizeCircle() {
    var circle = document.getElementById("circle");
    var currentSize = circle.offsetWidth;
    var newSize = currentSize + getRandomNumber(-50, 50); // Tăng hoặc giảm ngẫu nhiên từ -50px đến 50px
    circle.style.width = newSize + "px";
    circle.style.height = newSize + "px";
}

function getRandomNumber(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}