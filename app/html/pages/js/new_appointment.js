$(document).on("pageshow", "#new_appointment", function() {
    $('#btn_new_appointment_submit').prop('disabled', false);
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
        clientDetail = JSON.parse(data);
        var vehicle = clientDetail.VehiclesHistory[0];
        $("#new_appointment_vehicle").val(vehicle.VehicleDescription);
        $("#new_appointment_vehicle").data("vehicleId", vehicle.Id);
    });
}

function submitAppointmentForm() {
  $("#new_appointment_form").submit(function(e){
      e.preventDefault();
      var date = $("#new_appointment_date").datepicker("getDate");
      var vehicleId = $("#new_appointment_vehicle").data("vehicleId");
      var formData = "VehicleId=" + vehicleId;
      formData = formData + "&Date=" + getTimeInSeconds(date);
      $.ajax({
          url : '/appointment',
          type : 'post',
          data : formData,
          success : function() {
              $('btn_new_appointment_submit').prop("disabled", true);
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