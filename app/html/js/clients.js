$(document).on("pageshow", "#clients", function(){
  $.get( "/clients", function( data ) {
      var clients = JSON.parse(data);
      clients.forEach(function(client) {
          $('#clients_table tbody').append(
              '<tr><td><a href="#client_detail" onclick="setClientDetailData(' +
              client['Id'] + ')">' + client['Name'] +
              '</a></td></tr>'
          );
      });
  });
})

function setClientDetailData(clientId) {
  $('#client_detail').data("clientId", clientId);
  // $.mobile.pageContainer.pagecontainer("change", "#client_detail");
}