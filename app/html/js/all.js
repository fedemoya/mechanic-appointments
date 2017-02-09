var clientId;

function submitClientForm() {
  $("#new_client_form").submit(function(e){
      e.preventDefault();
      $.ajax({
          url:'/client',
          type:'post',
          data:$('#new_client_form').serialize(),
          success: function(data) {
              clientId = data;
              $('#new_client_confirm').fadeIn(2000);
              $('#new_client_confirm').fadeOut(2000);
          }
      });
      return false;
  });
}

// function goToNewVehicle() {

// }

function getURLParameter(sParam) {
    var sPageURL = window.location.search.substring(1);
    var sURLVariables = sPageURL.split('&');
     for (var i = 0; i < sURLVariables.length; i++) {
         var sParameterName = sURLVariables[i].split('=');
         if (sParameterName[0] == sParam) {
             return sParameterName[1];
         }
     }
}

function submitVehicleForm() {
  $("#new_vehicle_form").submit(function(e){
      e.preventDefault();
      var data = $('#new_vehicle_form').serialize();
      var clientId = getURLParameter("ClientId");
      data = data + "&ClientId=" + clientId;
      $.ajax({
          url:'/vehicle',
          type:'post',
          data:data,
          success:function(){

              $('#new_client_confirm').fadeIn(2000);
              $('#new_client_confirm').fadeOut(2000);
          }
      });
      return false;
  });
}

function load_appointments() {
    $.get( "/appointments", function( data ) {
        var appointments = JSON.parse(data);
        appointments.forEach(function(appointment) {
            console.log('<tr><td>' + appointment['ClientName'] + '</td><td>' + appointment['VehicleDescription'] + '</td></tr>');
            $('#appointments-table tbody').append(
                '<tr><td>' + appointment['ClientName'] + '</td><td>' + appointment['VehicleDescription'] + '</td></tr>'
            );
        });
    });
}