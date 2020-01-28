$(function(){
    $(document).on('click', '.floor', function(){
        clickFloor(this);
    });

    $(document).on('click', '#status-bar .team-bar', function(){
        var teamId = $(this).attr('team-id');

        $('#tenant p').hide();
        $('#tenant').find('.' + teamId).parent('p').show();
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
        bar += this.Name;
        bar += '</div>';
    });
    var $bar = $('#status-bar');
    $bar.html(bar);
    $.each(obj.barList, function(){
        $bar.children('.' + this.ClassName).css('width', this.Rate + '%');
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

    // $('#modal .modal-body').html('12234');

    $('#modal').modal('show');
}