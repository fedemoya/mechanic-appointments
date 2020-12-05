$(document).on("pagecreate", "#reparations", function () {
  $('#reparations_date').change(function () {
      $('#reparations_table tbody').empty();
      $('#reparations_date tbody').empty();   
      var date = $(this).datepicker("getDate");
      loadReparations(getTimeInSeconds(date));
  });
});

$(document).on("pageshow", "#reparations", function() {
  
  $('#reparations_table tbody').empty();
  $('#reparations_date').val("");

  loadReparations();
});

function loadReparations(date = '') {
  var path = "/api/reparations"
  if(date != '') {
      path = path + "/" + date
  }
  $.get(path, function( data ) {
      var reparations = JSON.parse(data);
      reparations.forEach(function(reparation) {
          var date = new Date(reparation['Date'] * 1000)
          $('#reparations_table tbody').append(
              '<tr><td><a href="#reparation_detail" onclick="setReparationDetailData(' +
              + reparation.Id + ')">' + reparation['ClientName'] +
              '</a></td><td>' + reparation['VehicleDescription'] +
              '</td><td>' + reparation['Price'] +
              '</td><td>' + reparation['Description'] +
              '</td><td>' + date.toLocaleDateString('es-AR') + '</td></tr>'
          );
      });
  });
}

function setReparationDetailData(id) {
  $('#reparation_detail').data("reparationId", id);
}
