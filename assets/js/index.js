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

    var bar = '';
    $.each(obj.barList, function(){
        console.log(this);
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
    html += '<p>シリアルコード</p>'
    html += '<input type="text" id="modalSerialCode">';
    html += '<button type="button" id="modalBtnRagist" class="btn btn-primary">登録</button>';

    $('#modal-body').html(html);
    $('#modal').modal('show');
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