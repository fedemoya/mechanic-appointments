$(document).on("pagecreate", "#appointments", function () {
  $('#appointments_date').change(function (e) {

        $('#appointments_table tbody').empty();
        
        var date = $(this).datepicker("getDate");
                
        $.get( "/api/appointments/" + getTimeInSeconds(date), function( data ) {
          var appointments = JSON.parse(data);
          build_appointments_table(appointments);
      });
    }
  );
});

$(document).on("pageshow", "#appointments", function() {

  $('#appointments_table tbody').empty();
  $('#appointments_date').val("");

  $.get( "/api/appointments", function( data ) {
      var appointments = JSON.parse(data);
      build_appointments_table(appointments);
  });

});

function refresh_appointments_table() {
    $.get( "/api/appointments", function( data ) {
        var appointments = JSON.parse(data);
        build_appointments_table(appointments);
    });
}

function build_appointments_table(appointments) {
    appointments.forEach(function (appointment) {
        var date = new Date(appointment['Date'] * 1000)
        var appointmentId = appointment['AppointmentId']
        var clientName = appointment['ClientName']
        var vehicleDescription = appointment['VehicleDescription']
        var dateString = date.toLocaleDateString("es-AR") +
            ' ' + date.toLocaleTimeString("es-AR")
        $('#appointments_table tbody').append(
            '<tr>' +
            '<td>' + clientName + '</td>' +
            '<td>' + vehicleDescription + '</td>' +
            '<td>' + dateString  + '</td>' +
            '<td>' + '<a href="#delete_appointment" data-role="button" data-icon="delete" ' +
                        'data-iconpos="notext" ' +
                        'data-mini="true" ' +
                        'data-inline="true" ' +
                        'data-transition="pop" ' +
                        'data-client-name="'+clientName+'" ' +
                        'data-vehicle-description="'+vehicleDescription+'" ' +
                        'data-date-string="'+dateString+'" ' +
                        'onclick="setDeleteAppointmentData(' +
                                    appointmentId + ')"' +
                        'id="a_delete_appointment_' + appointmentId + '">' +
                     '</a>' +
            '</td>' +
            '</tr>'
        );
        $('#appointments_table tbody').enhanceWithin();
    });
}

function setDeleteAppointmentData(appointmentId) {
    const aElement = $('#a_delete_appointment_' + appointmentId)
    $('#delete_appointment').data("clientName", aElement.data("clientName"));
    $('#delete_appointment').data("vehicleDescription", aElement.data("vehicleDescription"));
    $('#delete_appointment').data("dateString", aElement.data("dateString"));
    $('#delete_appointment').data("appointmentId", appointmentId);
}
