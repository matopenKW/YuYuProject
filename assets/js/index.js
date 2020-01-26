$(function(){
    $(document).on('click', '.floor', function(){
        clickFloor(this);
    });

    $(document).on('click', '#status-bar .team-bar', function(){
        var teamId = $(this).attr('team-id');

        $('#tenant p').hide();
        $('#tenant').find('.'+teamId).parent('p').show();
    });
});

function clickFloor(obj) {

    var data = {
        floorTd : $(obj).attr("floor-id")
    }
    var done = function(data){
        var list = JSON.parse(data)
        showModal(obj, list.tenantoList);
    }
    var fail = function(data){
        alert("失敗");
    }

    ajaxExecute("/floor", 'POST', data, done, fail);
}

function clickTeamBar(){

}

function showModal(obj, list){
    $('#modal .modal-title').html($(obj).html());

    var tenanto = '';
    $.each(list, function(){
        console.log(this);
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
    $('#tenant').append(tenanto);

    // $('#modal .modal-body').html('12234');

    $('#modal').modal('show');
}