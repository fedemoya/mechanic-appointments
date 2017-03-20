$(document).on("pageshow", "#new_reparation", function(){
    
    $.get( "/clients", function( data ) {
          var clients = JSON.parse(data);
          clients.forEach(function(client) {
              $('#new_reparation_clients').append(
                  '<li><a href="#" onclick="new_reparation_loadClientVehicle(' +
                  client['Id'] + ')">' + client['Name'] +
                  '</a></li>'
              );
              $('#new_reparation_clients').listview( "refresh" );
              $('#new_reparation_clients').trigger( "updatelayout");
          });
      });

    $('#reparation_full_payment').click(function () {
      if ($(this).is(':checked')) {
        $('#reparation_partial_payment').prop("disabled", "disabled");
      } else {
        $('#reparation_partial_payment').prop("disabled", "");
      }
    });
});

function new_reparation_loadClientVehicle(id) {
    $.get( "/client/" + id, function( data ) {
        clientDetail = JSON.parse(data);
        var vehicle = clientDetail.VehiclesHistory[0];
        $("#new_reparation_vehicle").val(vehicle.VehicleDescription);
        $("#new_reparation_vehicle").data("vehicleId", vehicle.Id);
    });
}

function submitReparationForm() {
  $("#new_reparation_form").submit(function(e){
      e.preventDefault();
      var date = $("#new_reparation_date").datepicker("getDate");
      var milliseconds_time = date.getTime();
      if (!date) {
        throw new Error('Missing date in new_reparation_form');
      }
      var seconds_time = milliseconds_time / 1000; 
      var vehicleId = $("#new_reparation_vehicle").data("vehicleId");
      var formData = "VehicleId=" + vehicleId;
      formData = formData + "&Date=" + seconds_time;
      var price = $("#new_reparation_price").val();
      formData = formData + "&Price=" + price;
      var description = $('#new_reparation_description').val()
      formData = formData + "&Description=" + description;
      var fullPayment = $('#reparation_full_payment').is(':checked') ? 1 : 0;
      formData = formData + "&FullPayment=" + fullPayment
      var partialPayment = $("#reparation_partial_payment").val();
      formData = formData + "&PartialPayment=" + partialPayment;
      $.ajax({
          url : '/reparation',
          type : 'post',
          data : formData,
          success : function(id) {
              $('#reparation_btn_submit').prop('disabled', true);
              if (fullPayment == 0) {
                $('#btn_load_payment').removeClass('ui-disabled');
              }
          }
      });
      return false;
  });
}

$(document).on("pagehide", "#new_reparation", function(){
  $('#new_reparation_form')[0].reset();
  $('#reparation_btn_submit').prop('disabled', false);
  $('#reparation_btn_submit').prop('value', "Enviar");
});