//@ts-check
import "../css/index.css";
import setUpChart from "./chart";
import countdown from "./countdown";
import eventStreamHandler from "./eventStreamHandler";

countdown();
const chartJsRef = setUpChart();
eventStreamHandler(chartJsRef);
