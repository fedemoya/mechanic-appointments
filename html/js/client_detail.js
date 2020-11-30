$(document).on("pageshow", "#client_detail", function(){
    
    $('#client_data').empty();
    $('#vehicles_history').empty();

    var clientId = $('#client_detail').data("clientId");
    
    if (!clientId) {
        throw new Error("Missing clientId in client_detail page");
    }
    
    $.get( "/api/client/" + clientId, function( data ) {
      
      var clientDetail = JSON.parse(data);
      
      $('#client_detail').data("ClientName", clientDetail.ClientName);
      
      $('#client_data').append(
        '<p><strong>' + 
        clientDetail.ClientName +
        '</strong></p>'
      );
      
      var i = 0;
      var vehiclesHistory = clientDetail.VehiclesHistory
      if (vehiclesHistory != null) {
        vehiclesHistory.forEach(function(vehicleHistory) {
            
            var html = '<div data-role="collapsible" data-mini="true">' +
              '<h4>' +
              vehicleHistory.VehicleDescription +
              '</h4>' +
              '<table data-role="table" data-mode="columntoggle" class="ui-responsive ui-body-d table-stripe" id="vehicle_' + i + '_history">' +
              '<thead>' +
              ' <tr class="ui-bar-d">' +
              '   <th>Fecha</th>' +
              '   <th>Precio</th>' +
              '   <th>Descripci&oacute;n</th>' +
              ' </tr>' +
              '</thead>' +
              '<tbody>' +
              '</tbody>' +
              '</table>' +
              '</div>';

            $('#vehicles_history').append(html).enhanceWithin();
            
            vehicleHistory.Reparations.forEach(function function_name(reparation) {
              var stringDate = epochToString(reparation['Date']);
              $('#vehicle_' + i + '_history tbody').append(
                '<tr><td>' + stringDate + '</td><td>' + reparation['Price'] + '</td><td>' + reparation['Description'] + '</td></tr>'
              );
            });
        });
      }
    });

});

function client_detail_setNewVehicleData() {
  var clientId = $('#client_detail').data("clientId");
  var clientName = $('#client_detail').data("ClientName");
  var CreatedClient = {"Id": clientId, "Name": clientName};
  $("#new_vehicle").data(CreatedClient);
}
