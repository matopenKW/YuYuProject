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

    $(document).on('click', '#modalBtnProductRagist', function(){
        ragistProduct();
    });

});

function clickFloor(obj) {
    var $floor = $(obj);

    var data = {
        floorId : $floor.attr("floor-id")
    }
    var done = function(data){
        var obj = JSON.parse(data);
        showModal($floor, obj);
    }
    var fail = function(data){
        alert("失敗");
    }

    ajaxExecute("/floor", 'POST', data, done, fail);
}

function clickTeamBar(){

}

function showModal($floor, obj){
    $('#modal .modal-title').html($floor.html());

    var temp = '<div id="status-bar" class="row"></div>';
    temp += '<div id="tenant"></div>';
    $('#modal-body').html(temp);
    $('#modal-body').css({
        backgroundImage: 'url("/assets/img/'+ $floor.attr("floor-id") +'.png")'
    });

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
        var obj = JSON.parse(data.responseJSON);
        alert(obj.message);
    }

    ajaxExecute("/registSerial", 'GET', data, done, fail);
}

function ragistProduct(){
    var $tenant = $('.tenant:checked')

    var data = {
        tenantId: $tenant.val(),
        productName: $('#productName').val(),
        productNo: $('#productNo').val()
    }
    var done = function(data){
        var obj = JSON.parse(data);
        alert(obj.message);
    }
    var fail = function(data){
        var obj = JSON.parse(data.responseJSON);
        alert(obj.message);
    }

    ajaxExecute("/ragistProduct", 'GET', data, done, fail);
}