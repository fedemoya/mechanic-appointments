$(document).on("pageshow", "#delete_appointment", function(){

    var clientName = $('#delete_appointment').data("clientName");
    var vehicleDescription = $('#delete_appointment').data("vehicleDescription");
    var dateString = $('#delete_appointment').data("dateString");

    $('#delete_appointment_message').empty()
    $('#delete_appointment_message').append(
        '<p><strong>' +
        '¿Borrar turno para ' + clientName +
        ', vehiculo ' + vehicleDescription + ', ' +
        'el día ' + dateString + '?' +
        '</strong></p>'
    );

});

function deleteAppointmentConfirmed() {
    var appoinmentId = $('#delete_appointment').data("appointmentId");
    $.ajax({
        url: '/api/appointment/' + appoinmentId,
        type: 'delete',
        success: function (clientId) {
            refresh_appointments_table()
        },
        error: function () {
            $('#delete_appointment_error').fadeIn(1000);
            $('#delete_appointment_error').fadeOut(1000);
        }
    })
}

