$(document).on("pageshow", "#appointments", function(){
  
  $.get( "/appointments", function( data ) {
      var appointments = JSON.parse(data);
      appointments.forEach(function(appointment) {
          console.log('<tr><td>' + appointment['ClientName'] + '</td><td>' + appointment['VehicleDescription'] + '</td></tr>');
          $('#appointments_table tbody').append(
              '<tr><td>' + appointment['ClientName'] + '</td><td>' + appointment['VehicleDescription'] + '</td></tr>'
          );
      });
  });

  $('#appointments_date').change(function (e) {

        $('#appointments_table tbody').empty();
        
        var date = $(this).datepicker("getDate");
        var milliseconds_time = date.getTime();
        var seconds_time = milliseconds_time / 1000;
                
        $.get( "/appointments/" + seconds_time, function( data ) {
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

$(document).on("pagehide", "#appointments", function(){
  $('#appointments_table tbody').empty();
  $('#appointments_date').val("");
});
