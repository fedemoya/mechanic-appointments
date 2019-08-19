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
    var data = 'ReparationId=' + reparationId + '&Amount=' + amount + '&Date=' + getTimeInSeconds(date);
    $.ajax({
        url:'/payment',
        type:'post',
        data: data,
        success: function() {
            clearReparationDetailPage();
            loadReparationData(reparationId);
        }
    });
}

function loadReparationData(id) {
    $.get("/reparation/" + id, function( data ) {
      
        var reparationDetail = JSON.parse(data);

        var stringDate = epochToString(reparationDetail.Date);        
            
        $('#reparation_data').append(
        '<p><strong>Fecha: ' + 
        stringDate +
        '</strong></p>' +
        '<p><strong>Descripci&oacute;n: ' + 
        reparationDetail.Description +
        '</strong></p>' +
        '<p><strong>Precio: ' + 
        reparationDetail.Price +
        '</strong></p>'
        );

        reparationDetail.Payments.forEach(function (payment) {
            var stringDate = epochToString(payment['Date']);
            $('#payments_table tbody').append(
              '<tr><td>' + stringDate +
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