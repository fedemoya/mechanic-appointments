$(document).on("pageshow", "#new_payment", function(){
  var reparationData = $('#new_payment').data();
  if (!reparationData || reparationData["Id"] == 0) {
      throw new Error("Missing reparationData in client_detail page");
  }
  $('#new_payment_client').val(createdClient["Name"]);
});

function submitVehicleForm() {
  var formData = $('#new_payment_form').serialize();
  var createdClient = $('#new_payment').data();
  if (!createdClient) {
    throw new Error('Missing client data in new_payment_form');
  }
  var clientId = createdClient["Id"];
  formData = formData + "&ClientId=" + clientId;
  $.ajax({
      url : '/api/vehicle',
      type : 'post',
      data : formData,
      success : function(){
          $('#new_payment_confirm').fadeIn(1000);
          $('#new_payment_confirm').fadeOut(1000);
      },
      error : function() {
          $('#new_payment_error').fadeIn(1000);
          $('#new_payment_error').fadeOut(1000);
      }
  });
}

$(document).on("pagehide", "#new_payment", function(){
  $('#new_payment_form')[0].reset();
});