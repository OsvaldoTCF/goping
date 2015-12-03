Chart.defaults.global.responsive = true;

var origin = "";
var time = "";

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

function initEmptyChart() {
  var ctx = $("#averages").get(0).getContext("2d");
  var chart = new Chart(ctx).Line(data);

  return chart;
}

function updateChart(chart) {
  chart.destroy();
  ctx = $("#averages").get(0).getContext("2d");
  chart = new Chart(ctx).Line(data);

  return chart;
}

function clearData() {
  data.datasets[0].data = [];
}

function updateOrigin(newOrigin) {
    $("#origin-text").text(origin)
}

var weekday = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"];
var months = ["January", "February", "March", "April", "May", "June", "July",
              "August", "September", "October", "November", "December"];

function updateTime(newTime) {
  time = newTime;
  date = new Date(time);

  $("#day").text(weekday[date.getDay()] + ", "
      + months[date.getMonth()] + " " + date.getDate() + ", "
      + date.getFullYear()
  );
}

$(window).load(function() {
  var chart = initEmptyChart();

  $("#previous-day-btn").click(function() {
    $.get("http://localhost:8080/api/2/pings/" + origin + "/" + time + "/prev", function(avgCollection) {
      clearData();

      if (avgCollection.times.length != 0) {
        updateTime(avgCollection.times[0]);
      }

      for (i = 0; i < avgCollection.averages.length; i++) {
        data.datasets[0].data.push(avgCollection.averages[i]);
      }

      chart = updateChart(chart);
    });
  });

  $("#next-day-btn").click(function() {
    $.get("http://localhost:8080/api/2/pings/" + origin + "/" + time + "/next", function(avgCollection) {
      clearData();

      if (avgCollection.times.length != 0) {
        updateTime(avgCollection.times[0]);
      }

      for (i = 0; i < avgCollection.averages.length; i++) {
        data.datasets[0].data.push(avgCollection.averages[i]);
      }

      chart = updateChart(chart);
    });
  });

  $("#oldest-btn").click(function() {
  });

  $("#now-btn").click(function() {
  });

  $(".dropdown-menu li").click(function() {
    origin = $(this).text();
    updateOrigin($(this).text());

    $.get("http://localhost:8080/api/1/pings/" + origin + "/hours", function(avgCollection) {
      clearData();

      if (avgCollection.times.length != 0) {
        updateTime(avgCollection.times[0]);
      }

      for (i = 0; i < avgCollection.averages.length; i++) {
        data.datasets[0].data.push(avgCollection.averages[i]);
      }

      chart = updateChart(chart);
    });
  });
})
