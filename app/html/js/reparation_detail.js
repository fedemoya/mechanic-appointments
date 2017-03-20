$(document).on("pageshow", "#reparation_detail", function() {

    var reparationId = $('#reparation_detail').data("reparationId");

    if (!reparationId) {
        throw new Error("Missing reparationId in client_detail page");
    }

    loadReparationData(reparationId);

});

function submitPaymentForm(argument) {
    var reparationId = $('#reparation_detail').data("reparationId");
    var amount = $('#new_payment_amount').val();
    var date = new Date()
    var seconds_date = date.getTime() / 1000
    var data = 'ReparationId=' + reparationId + '&Amount=' + amount + '&Date=' + seconds_date;
    $.ajax({
        url:'/payment',
        type:'post',
        data: data,
        success: function() {
            clearReparationDetailPage();
            loadReparationData(reparationId);
        }
    });
    return false;
}

function loadReparationData(id) {
    $.get("/reparation/" + id, function( data ) {
      
        var reparationDetail = JSON.parse(data);

        var milliseconds_date = reparationDetail.Date * 1000;
        var d =  new Date(milliseconds_date);
        var string_date = d.getDate() + '/' + d.getMonth() + '/' + d.getFullYear();
            
        $('#reparation_data').append(
        '<p><strong>Fecha: ' + 
        string_date +
        '</strong></p>' +
        '<p><strong>Descripci&oacute;n: ' + 
        reparationDetail.Description +
        '</strong></p>' +
        '<p><strong>Precio: ' + 
        reparationDetail.Price +
        '</strong></p>'
        );

        reparationDetail.Payments.forEach(function (payment) {
        var milliseconds_date = payment['Date'] * 1000;
        var d =  new Date(milliseconds_date);
        var string_date = d.getDate() + '/' + d.getMonth() + '/' + d.getFullYear();
          $('#payments_table tbody').append(
              '<tr><td>' + string_date +
              '</td><td>' + payment['Amount'] +
              '</td></tr>'
          );
        });

    });
}

function clearReparationDetailPage() {
    $('new_payment_amount').val("");
    $('#reparation_data').empty();
    $('#payments_table tbody').empty();
}

$(document).on("pagehide", "#reparation_detail", function() {
    clearReparationDetailPage();
})