$(document).on("pagecreate", "#new_reparation", function () {
  $('#reparation_full_payment').click(function () {
      if ($(this).is(':checked')) {
        $('#reparation_partial_payment').prop("disabled", "disabled");
      } else {
        $('#reparation_partial_payment').prop("disabled", "");
      }
    });
});

$(document).on("pageshow", "#new_reparation", function(){
    
    $('#reparation_partial_payment').prop("disabled", "");
    $('#new_reparation_clients').empty();
    $('#new_reparation_form')[0].reset();
    $('#reparation_btn_submit').prop('disabled', false);
    
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
      var vehicleId = $("#new_reparation_vehicle").data("vehicleId");
      var formData = "VehicleId=" + vehicleId;
      formData = formData + "&Date=" + getTimeInSeconds(date);
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
              $('#btn_new_reparation_submit').prop('disabled', true);
              $('#new_reparation_confirm').fadeIn(1000);
              $('#new_reparation_confirm').fadeOut(1000);
          }
      });
      return false;
  });
}
