Chart.defaults.global.responsive = true;

var data = {
  labels: [
    "12:00AM", "01:00", "02:00", "03:00", "04:00", "05:00",
      "06:00", "07:00", "08:00", "09:00", "10:00", "11:00",
    "12:00PM", "01:00", "02:00", "03:00", "04:00", "05:00",
      "06:00", "07:00", "08:00", "09:00", "10:00", "11:00"
  ],
  datasets: [
    {
      label: "Week average transfer", 
      fillColor: "rgba(220,220,220,0.2)",
      strokeColor: "rgba(220,220,220,1)",
      pointColor: "rgba(220,220,220,1)",
      pointStrokeColor: "#fff",
      pointHighlightFill: "#fff",
      pointHighlightStroke: "rgba(220,220,220,1)",
      data: [
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
      ]
    }
  ]
};

$(window).load(function() {
  // Initialization
  var ctx = $("#averages").get(0).getContext("2d");
  var chart = new Chart(ctx).Line(data);

  $(".dropdown-menu li").click(function() {
    var origin = $(this).text();
    var url_api = "http://localhost:8080/api/1/pings/" + origin + "/hours";

    // Change the `Origin` button text
    $("#origin-text").text(origin)

    $.get(url_api, function(avgCollection) {
      data.datasets[0].data = [];
      data.labels = [];

      for (i = 0; i < avgCollection.averages.length; i++) {
        data.datasets[0].data.push(avgCollection.averages[i]);
        data.labels.push(avgCollection.times[i]);
      }

      chart.destroy();
      ctx = $("#averages").get(0).getContext("2d");
      chart = new Chart(ctx).Line(data);
    });
  });
})
