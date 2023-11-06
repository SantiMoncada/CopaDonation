//@ts-check
import {
  Chart,
  BarController,
  CategoryScale,
  LinearScale,
  BarElement,
} from "chart.js";

Chart.register(BarController, CategoryScale, LinearScale, BarElement);

export const setUpChart = () => {
  const xValues = [55, 49, 44, 24, 15];
  const yValues = ["WEB", "UX/UI", "DATA"];
  const barColors = [
    "rgb(253, 224, 71)",
    "rgb(192, 132, 252)",
    "rgb(134, 239, 172)",
  ];

  new Chart("myChart", {
    type: "bar",
    data: {
      labels: yValues,
      datasets: [
        {
          backgroundColor: barColors,
          data: xValues,
        },
      ],
    },
    options: {
      indexAxis: "y",
      plugins: {
        legend: {
          display: false,
        },
      },
      responsive: true,
      maintainAspectRatio: false,
      scales: {
        y: {
          stacked: true,
        },
        x: {
          grid: {
            display: false,
          },
        },
      },
    },
  });
};
