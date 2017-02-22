$(document).on("pageshow", "#client_detail", function(){
    var clientId = $('#client_detail').data("clientId");
    if (!clientId) {
        throw new Error("Missing clientId in client_detail page");
    }
    $.get( "/client/" + clientId, function( data ) {
      var clientDetail = JSON.parse(data);
      $('#client_data').append(
        '<p><strong>' + 
        clientDetail.ClientName +
        '</strong></p>'
        );
      clientDetail.Vehicles.forEach(function(vehicle) {
          $('#client_vehicles_table tbody').append(
              '<tr><td>' +
              vehicle.Brand + " " + vehicle.Model +
              '</td></tr>'
          );
      });
    });
})