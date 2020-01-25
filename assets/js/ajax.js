
function ajaxExecute(url, type, data, done, fail, always){
    $.ajax({
        url: url,
        type: type,
        data: data
    })
    .done(function(data){
        console.log(data);
        done(data);
    })
    .fail(function(data){
        console.log(data);
        fail(data);
    })
    .always(always);
}