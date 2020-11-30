$(document).on("pagecreate", "#clients", function () {
  $('#filter_debtors').click(function () {
      loadClients();
  });
});

$(document).on("pageshow", "#clients", function() {
  loadClients();
});

function loadClients() {
  $('#clients_table tbody').empty();
  var route = "/api/clients";
  if ($('#filter_debtors').is(':checked')) {
    route = route + "/debtors";
  }
  $.get(route, function( data ) {
      var clients = JSON.parse(data);
      clients.forEach(function(client) {
          $('#clients_table tbody').append(
              '<tr><td><a href="#client_detail" onclick="setClientDetailData(' +
              client['Id'] + ')">' + client['Name'] +
              '</a></td></tr>'
          );
      });
  });
}

function setClientDetailData(clientId) {
  $('#client_detail').data("clientId", clientId);
}