var video;
var canvasElement;
var canvas;
var loadingMessage;
var outputContainer;
var outputMessage;
var outputData;

$(function(){
    $(document).on('click', '.floor', function(){
        clickFloor(this);
    });

    $(document).on('click', '#status-bar .team-bar', function(){
        var teamId = $(this).attr('team-id');

        $('#tenant p').hide();
        $('#tenant').find('.' + teamId).parent('p').show();
    });

    $(document).on('click','#btnRagistSerial', function(){
        showRegistSerialModal();
    });

    $(document).on('click', '#modalBtnRagist', function(){
        ragistSerial();
    });

});

function clickFloor(obj) {
    var floor = $(obj).html();

    var data = {
        floorId : $(obj).attr("floor-id")
    }
    var done = function(data){
        var obj = JSON.parse(data);
        showModal(floor, obj);
    }
    var fail = function(data){
        alert("失敗");
    }

    ajaxExecute("/floor", 'POST', data, done, fail);
}

function clickTeamBar(){

}

function showModal(floor, obj){

    $('#modal .modal-title').html(floor);

    var temp = '<div id="status-bar" class="row"></div>';
    temp += '<div id="tenant"></div>';
    $('#modal-body').html(temp);

    var bar = '';
    $.each(obj.barList, function(){
        bar += '<div team-id="';
        bar += this.ClassName;
        bar += '" class="';
        bar += this.ClassName;
        bar += ' team-bar text-center">';
        bar += this.Id;
        bar += '</div>';
    });
    var $bar = $('#status-bar');
    $bar.html(bar);
    $.each(obj.barList, function(){
        $bar.children('.' + this.ClassName).css('width', this.All + '%');
    });

    var tenanto = '';
    $.each(obj.tenantoList, function(){
        tenanto += '<p><span class="';
        tenanto += this.ClassName;
        tenanto += '">';
        if (this.Acquisition) {
            tenanto += this.Acquisition;
        }
        tenanto += '</span>';
        tenanto += this.Name;
        tenanto += '</p>';
    });
    $('#tenant').html(tenanto);
    $('#modal').modal('show');
}

function showRegistSerialModal(){
    $('#modal .modal-title').html('シリアルコードを登録');
    var html = '';
    html += '<div id="modal-ragist-serial class="text-center"">';
    html += '<div class="row"><p>シリアルコード</p></div>';
    html += '<div id="loadingMessage">🎥 Unable to access video stream (please make sure you have a webcam enabled)</div>';
    html += '<canvas id="canvas" hidden></canvas>';
    html += '<div id="output" hidden>';
    html += '<div id="outputMessage">No QR code detected.</div>';
    html += '<div hidden><b>Data:</b> <span id="outputData"></span></div>';
    html += '</div>';
	html += '<div class="row"><button type="button" id="modalBtnRagist" class="btn btn-primary">登録</button></div>';
	html += '</div>';
    $('#modal-body').html(html);
    $('#modal').modal('show');

    video = document.createElement("video");
    canvasElement = document.getElementById("canvas");
    canvas = canvasElement.getContext("2d");
    loadingMessage = document.getElementById("loadingMessage");
    outputContainer = document.getElementById("output");
    outputMessage = document.getElementById("outputMessage");
    outputData = document.getElementById("outputData");

    navigator.mediaDevices.getUserMedia({ video: { facingMode: "environment" } }).then(function(stream) {
        video.srcObject = stream;
        video.setAttribute("playsinline", true); // required to tell iOS safari we don't want fullscreen
        video.play();
        requestAnimationFrame(tick);
    });
}

function ragistSerial(){

    var serialCode = $('#modalSerialCode').val();

    var data = {
        serialCode : serialCode
    }
    var done = function(data){
        var obj = JSON.parse(data);
        alert(obj.message);
    }
    var fail = function(data){
        var obj = JSON.parse(data);
        var msg = "エラーが発生しました。/n";
        msg += "tetetet";
        alert(msg);
    }

    ajaxExecute("/registSerial", 'POST', data, done, fail);

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