Chart.defaults.global.responsive = true;

var data = {
  labels: ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"],
  datasets: [
    {
      label: "Week average transfer", 
      fillColor: "rgba(220,220,220,0.2)",
      strokeColor: "rgba(220,220,220,1)",
      pointColor: "rgba(220,220,220,1)",
      pointStrokeColor: "#fff",
      pointHighlightFill: "#fff",
      pointHighlightStroke: "rgba(220,220,220,1)",
      data: [17, 98, 83, 48, 10, 99, 67]
    }
  ]
};

$(window).load(function() {
  var ctx = $("#averages").get(0).getContext("2d");
  var avgChart = new Chart(ctx).Line(data);
})
