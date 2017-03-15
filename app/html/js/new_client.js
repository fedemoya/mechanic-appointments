var CreatedClient = {"Name": "", "Id": 0};

function submitClientForm() {
  $("#new_client_form").submit(function(e){
      e.preventDefault();
      var formData = $('#new_client_form').serialize();
      CreatedClient["Name"] = $('#new_client_name').val();
      $.ajax({
          url:'/client',
          type:'post',
          data: formData,
          success: function(clientId) {
              CreatedClient["Id"] = clientId;
              $('#new_client_confirm').html("<strong>Nuevo cliente creado</strong>");
              $('#btn_load_vehicle').show()
          }
      });
      return false;
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