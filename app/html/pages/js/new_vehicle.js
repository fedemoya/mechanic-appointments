$(document).on("pageshow", "#new_vehicle", function() {
    $('#btn_new_vehicle_submit').prop('disabled', false);
    var createdClient = $('#new_vehicle').data();
    if (!createdClient || createdClient["Id"] == 0) {
        throw new Error("Missing CreatedClient in client_detail page");
    }
    $('#new_vehicle_client').val(createdClient["Name"]);
});

function submitVehicleForm() {
  var formData = $('#new_vehicle_form').serialize();
  var createdClient = $('#new_vehicle').data();
  if (!createdClient) {
    throw new Error('Missing client data in new_vehicle_form');
  }
  var clientId = createdClient["Id"];
  formData = formData + "&ClientId=" + clientId;
  $.ajax({
      url : '/vehicle',
      type : 'post',
      data : formData,
      success : function() {
          $('#btn_new_vehicle_submit').prop("disabled", true);
          $('#new_vehicle_confirm').fadeIn(1000);
          $('#new_vehicle_confirm').fadeOut(1000);
      }
  });
}

$(document).on("pagehide", "#new_vehicle", function(){
  $('#new_vehicle_form')[0].reset();
});