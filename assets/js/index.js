$(function(){
    $(document).on('click', '.floor', function(){

        

        $('#modal .modal-title').html($(this).html());
        $('#modal .modal-body').html('12234');

        $('#modal').modal('show');

    });
});