function submitVehicleForm() {
  $("#new_vehicle_form").submit(function(e){
      e.preventDefault();
      var formData = $('#new_vehicle_form').serialize();
      var CreatedClient = $('#new_vehicle').data();
      if (!CreatedClient) {
        throw new Error('Missing client data in new_vehicle_form');
      }
      var clientId = CreatedClient["Id"];
      formData = formData + "&ClientId=" + clientId;
      $.ajax({
          url : '/vehicle',
          type : 'post',
          data : formData,
          success : function(){
              $('#new_vehicle_confirm').fadeIn(1000);
              $('#new_vehicle_confirm').fadeOut(1000);
          }
      });
      return false;
  });
}