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

function build_appointments_table(appointments) {
    appointments.forEach(function (appointment) {
        var date = new Date(appointment['Date'] * 1000)
        $('#appointments_table tbody').append(
            '<tr>' +
            '<td>' + appointment['ClientName'] + '</td>' +
            '<td>' + appointment['VehicleDescription'] + '</td>' +
            '<td>' +  date.toLocaleDateString("es-AR") +
                      ' ' + date.toLocaleTimeString("es-AR") + '</td>' +
            '</tr>'
        );
    });
}
