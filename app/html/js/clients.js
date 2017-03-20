$(document).on("pageshow", "#clients", function(){
  
  loadClients(false);

  $('#filter_debtors').click(function () {
      $('#clients_table tbody').empty();
      if ($(this).is(':checked')) {
        loadClients(true);
      } else {
        loadClients(false);
      }
  });

})

function loadClients(filterDebtors) {
  var route = filterDebtors ? "/clients/debtors" : "/clients"
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

$(document).on("pagehide", "#clients", function(){
  $('#clients_table tbody').empty();
})