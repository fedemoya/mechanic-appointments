$(document).on("pageshow", "#reparations", function() {

  var milliseconds_time = $.now();
  var seconds_time = Math.floor(milliseconds_time / 1000);

  loadReparations(seconds_time);

  $('#reparations_date').change(function (e) {

      $('#reparations_table tbody').empty();
      $('#reparations_date tbody').empty();
        
      var date = $(this).datepicker("getDate");
      var milliseconds_time = date.getTime();
      var seconds_time = Math.floor(milliseconds_time / 1000);
                
      loadReparations(seconds_time);
  });

});

function loadReparations(date) {
  $.get( "/reparations/" + date, function( data ) {
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

$(document).on("pagehide", "#reparations", function(){
  $('#reparations_table tbody').empty();
  $('#reparations_date').val("");
})
