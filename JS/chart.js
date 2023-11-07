//@ts-check
import {
  Chart,
  BarController,
  CategoryScale,
  LinearScale,
  BarElement,
} from "chart.js";

/**
 * Function that creates and returns a Chart.js chart.
 *
 * @returns {Chart<"bar", number[], string> | undefined} The Chart.js chart instance with specific chart type, data, and label types.
 */
export default () => {
  Chart.register(BarController, CategoryScale, LinearScale, BarElement);

  const chart = document.getElementById("myChart");

  const webAmount = chart?.dataset.web;
  const uxAmount = chart?.dataset.ux;
  const dataAmount = chart?.dataset.data;

  if (!webAmount || !uxAmount || !dataAmount) {
    return;
  }

  const xValues = [
    parseFloat(webAmount),
    parseFloat(uxAmount),
    parseFloat(dataAmount),
  ];

  const yValues = ["WEB", "UX/UI", "DATA"];
  const barColors = [
    "rgb(253, 224, 71)",
    "rgb(192, 132, 252)",
    "rgb(134, 239, 172)",
  ];

  const chartJsRef = new Chart("myChart", {
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
          ticks: {
            color: barColors,
            font: {
              weight: "900",
            },
            autoSkip: false,
            maxRotation: 45,
            minRotation: 45,
          },
        },
      },
    },
  });

  return chartJsRef;
};
