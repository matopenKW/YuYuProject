var video;
var canvasElement;
var canvas;
var loadingMessage;
var outputContainer;
var outputMessage;
var outputData;

function standby(){
    video = document.createElement("video");
    canvasElement = document.getElementById("canvas");
    canvas = canvasElement.getContext("2d");
    loadingMessage = document.getElementById("loadingMessage");
    outputContainer = document.getElementById("output");
    outputMessage = document.getElementById("outputMessage");
    outputData = document.getElementById("outputData");
}
    

function drawLine(begin, end, color) {
    canvas.beginPath();
    canvas.moveTo(begin.x, begin.y);
    canvas.lineTo(end.x, end.y);
    canvas.lineWidth = 4;
    canvas.strokeStyle = color;
    canvas.stroke();
}

function tick() {
    loadingMessage.innerText = "⌛ Loading video..."
    if (video.readyState === video.HAVE_ENOUGH_DATA) {
        loadingMessage.hidden = true;
        canvasElement.hidden = false;
        outputContainer.hidden = false;

        canvasElement.height = video.videoHeight;
        canvasElement.width = video.videoWidth;
        canvas.drawImage(video, 0, 0, canvasElement.width, canvasElement.height);
        var imageData = canvas.getImageData(0, 0, canvasElement.width, canvasElement.height);
        var code = jsQR(imageData.data, imageData.width, imageData.height, {
        inversionAttempts: "dontInvert",
        });
        if (code) {
        drawLine(code.location.topLeftCorner, code.location.topRightCorner, "#FF3B58");
        drawLine(code.location.topRightCorner, code.location.bottomRightCorner, "#FF3B58");
        drawLine(code.location.bottomRightCorner, code.location.bottomLeftCorner, "#FF3B58");
        drawLine(code.location.bottomLeftCorner, code.location.topLeftCorner, "#FF3B58");
        outputMessage.hidden = true;
        outputData.parentElement.hidden = false;
        outputData.innerText = code.data;
        } else {
        outputMessage.hidden = false;
        outputData.parentElement.hidden = true;
        
        }
    }
    requestAnimationFrame(tick);
}