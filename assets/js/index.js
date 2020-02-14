var viewSpeed = 'normal';

$(function(){
    $(document).on('click', '.floor', function(){
        clickFloor(this);
    });

    $(document).on('click', '#status-bar .team-bar', function(){
        var teamId = $(this).attr('team-id');

        $('#tenant p').hide();
        $('#tenant').find('.' + teamId).parent('p').show('normal');
    });

    $(document).on('click','#btnRagistSerial', function(){
        showRegistSerialModal();
    });

    $(document).on('click','#view-graph-area', function(){
        $(".view-graph-toggle").toggle(viewSpeed);
    });

    $(document).on('click','#view-product-graph-area', function(){
        $(".view-product-graph-toggle").toggle(viewSpeed);
    });

    $(document).on('click', '.product-bar', function(){
        showProductDetil(this);
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
    if (!serialCode) {
        alert('しりあるこーどがみにゅうりょくです。');
        return
    }

    var data = {
        serialCode : serialCode
    }
    var done = function(data){
        var obj = JSON.parse(data);
        alert(obj.message);
        $('#serial-modal').modal('hide');
    }
    var fail = function(data){
        var obj = JSON.parse(data.responseJSON);
        alert(obj.message);
    }

    ajaxExecute("/registSerial", 'GET', data, done, fail);
}

function ragistProduct(){
    var $tenant = $('.tenant:checked');
    var productName = $('#productName').val();
    var productNo = $('#productNo').val();
    if (!$tenant[0]){
        alert('テナントを選択から出直してね');
        return
    }
    if (productName === '') {
        alert('商品名が空です。入力の仕方分かりますか？？');
        return
    }
    if (productNo === '') {
        alert('商品番号が空です。この並びだと普通入れると思うんですけど…');
        return
    }

    var data = {
        tenantId: $tenant.val(),
        productName: productName,
        productNo: $('#productNo').val()
    }
    var done = function(data){
        var obj = JSON.parse(data);
        alert(obj.message);
        $('#serial-modal').modal('hide');
    }
    var fail = function(data){
        var obj = JSON.parse(data.responseJSON);
        alert(obj.message);
    }

    ajaxExecute("/ragistProduct", 'GET', data, done, fail);
}

function showProductDetil(obj){
    var teamName = $(obj).attr('class-data');
    var tenantId = $(obj).attr('tenant-id');

    var $targetObj = $('.product-detil.' + teamName + '.' + tenantId);
    $('.product-detil').hide(viewSpeed);
    $targetObj.show(viewSpeed);
}