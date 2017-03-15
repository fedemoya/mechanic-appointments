$(document).on("pageshow", "#new_appointment", function() {
    $.get( "/clients", function( data ) {
          var clients = JSON.parse(data);
          clients.forEach(function(client) {
              $('#new_appointment_clients').append(
                  '<li><a href="#" onclick="new_appointment_loadClientVehicle(' +
                  client['Id'] + ')">' + client['Name'] +
                  '</a></li>'
              );
              $('#new_appointment_clients').listview( "refresh" );
              $('#new_appointment_clients').trigger( "updatelayout");
          });
      }); 
});

function new_appointment_loadClientVehicle(id) {
    $.get( "/client/" + id, function( data ) {
        var clientDetail = JSON.parse(data);
        var vehicle = clientDetail.Vehicles[0];
        console.log(vehicle.Brand + " " + vehicle.Model);
        $("#new_appointment_vehicle").val(vehicle.Brand + " " + vehicle.Model);
        $("#new_appointment_vehicle").data("vehicleId", vehicle.Id);
    });
}

function submitAppointmentForm() {
  $("#new_appointment_form").submit(function(e){
      e.preventDefault();
      var date = $("#new_appointment_date").datepicker("getDate");
      var milliseconds_time = date.getTime();
      if (!date) {
        throw new Error('Missing date in new_appointment_form');
      }
      var seconds_time = milliseconds_time / 1000; 
      var vehicleId = $("#new_appointment_vehicle").data("vehicleId");
      var formData = "VehicleId=" + vehicleId;
      formData = formData + "&Date=" + seconds_time;
      $.ajax({
          url : '/appointment',
          type : 'post',
          data : formData,
          success : function(){
              $('#new_appointment_confirm').fadeIn(1000);
              $('#new_appointment_confirm').fadeOut(1000);
          }
      });
      return false;
  });
}

$(document).on("pagehide", "#new_appointment", function(){
  $('#new_appointment_clients').empty();
  $('#new_appointment_form')[0].reset();
});