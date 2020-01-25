$(function(){
    $(document).on('click', '.floor', function(){
        clickFloor(this);
    });
});

function clickFloor(obj) {

    var data = {
        floorTd : $(obj).attr("floor-id")
    }
    var done = function(data){
        showModal(obj);
    }
    var fail = function(data){
        alert("失敗");
    }

    ajaxExecute("/floor", 'POST', data, done, fail);
    
}

function showModal(obj){
    $('#modal .modal-title').html($(obj).html());
    // $('#modal .modal-body').html('12234');

    $('#modal').modal('show');
}