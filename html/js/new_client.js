$(document).on('pageshow', '#new_client', function () {
  $('#btn_new_client_submit').prop('disabled', false);
});

var CreatedClient = {"Name": "", "Id": 0};

function submitClientForm() {
  var formData = $('#new_client_form').serialize();
  CreatedClient["Name"] = $('#new_client_name').val();
  $.ajax({
      url: '/api/client',
      type:'post',
      data: formData,
      success: function(clientId) {
          CreatedClient["Id"] = clientId;
          $('#btn_new_client_submit').prop("disabled", true);
          $('#new_client_confirm').fadeIn(1000);
          $('#new_client_confirm').fadeOut(1000, function () {
            $('#btn_load_vehicle').show();
          });
      }
  });
}

function new_client_setNewVehicleData() {
  if (!CreatedClient) {
    throw new Error('Missing client data in setNewVehicleData()');
  }
  $("#new_vehicle").data(CreatedClient);
}

$(document).on("pagehide", "#new_client", function(){
  $('#new_client_form')[0].reset();
  $('#new_client_confirm').html("");
  $('#btn_load_vehicle').hide();
})