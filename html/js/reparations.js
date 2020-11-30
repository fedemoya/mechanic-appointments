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

  loadReparations(getCurrentTimeInSeconds());
});

function loadReparations(date) {
  $.get( "/api/reparations/" + date, function( data ) {
      var reparations = JSON.parse(data);
      reparations.forEach(function(reparation) {
        console.log('<tr><td><a href="#reparation_detail" onclick="setReparationDetailData(' +
              + reparation.Id + ')">' + reparation['ClientName'] +
              '</a></td><td>' + reparation['VehicleDescription'] +
              '</td><td>' + reparation['Price'] +
              '</td><td>' + reparation['Description'] + '</td></tr>');
          $('#reparations_table tbody').append(
              '<tr><td><a href="#reparation_detail" onclick="setReparationDetailData(' +
              + reparation.Id + ')">' + reparation['ClientName'] +
              '</a></td><td>' + reparation['VehicleDescription'] +
              '</td><td>' + reparation['Price'] +
              '</td><td>' + reparation['Description'] + '</td></tr>'
          );
      });
  });
}

function setReparationDetailData(id) {
  $('#reparation_detail').data("reparationId", id);
}
