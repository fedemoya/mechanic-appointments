$(document).on("pagecreate", "#appointments", function () {
  $('#appointments_date').change(function (e) {

        $('#appointments_table tbody').empty();
        
        var date = $(this).datepicker("getDate");
                
        $.get( "/appointments/" + getTimeInSeconds(date), function( data ) {
          var appointments = JSON.parse(data);
          appointments.forEach(function(appointment) {
              console.log('<tr><td>' + appointment['ClientName'] + '</td><td>' + appointment['VehicleDescription'] + '</td></tr>');
              $('#appointments_table tbody').append(
                  '<tr><td>' + appointment['ClientName'] + '</td><td>' + appointment['VehicleDescription'] + '</td></tr>'
              );
          });
      });
    }
  );
});

$(document).on("pageshow", "#appointments", function() {

  $('#appointments_table tbody').empty();
  $('#appointments_date').val("");
  
  $.get( "/appointments", function( data ) {
      var appointments = JSON.parse(data);
      appointments.forEach(function(appointment) {
          console.log('<tr><td>' + appointment['ClientName'] + '</td><td>' + appointment['VehicleDescription'] + '</td></tr>');
          $('#appointments_table tbody').append(
              '<tr><td>' + appointment['ClientName'] + '</td><td>' + appointment['VehicleDescription'] + '</td></tr>'
          );
      });
  });

});
