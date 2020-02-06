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

    $(document).on('click', '#btnCopy', function(){
        clipboardCopy();
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

    $('#serial-modal').modal('show');
    standby();
    cameraStart(); 
}

function cameraStart(){
    navigator.mediaDevices.getUserMedia({ video: { facingMode: "environment" } }).then(function(stream) {
        video.srcObject = stream;
        video.setAttribute("playsinline", true); // required to tell iOS safari we don't want fullscreen
        video.play();
        requestAnimationFrame(tick);
    });
}

function clipboardCopy() {
    var copyTarget = $("#outputData")[0];
    document.getSelection().selectAllChildren(copyTarget);
    document.execCommand("copy");
}

function ragistSerial(){

    var serialCode = $('#serial-code-text').val();

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
        alert(msg);
    }

    ajaxExecute("/registSerial", 'POST', data, done, fail);

}
