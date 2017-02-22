$(document).on("pageshow", "#appointments", function(){
  $.get( "/appointments", function( data ) {
      var appointments = JSON.parse(data);
      appointments.forEach(function(appointment) {
          console.log('<tr><td>' + appointment['ClientName'] + '</td><td>' + appointment['VehicleDescription'] + '</td></tr>');
          $('#appointments-table tbody').append(
              '<tr><td>' + appointment['ClientName'] + '</td><td>' + appointment['VehicleDescription'] + '</td></tr>'
          );
      });
  });
})