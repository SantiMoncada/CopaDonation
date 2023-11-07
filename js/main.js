//@ts-check
import setUpChart from "./chart";
import eventStreamHandler from "./eventStreamHandler";
import countdown from "./countdown";


countdown();
const chartJsRef = setUpChart();
eventStreamHandler(chartJsRef);
