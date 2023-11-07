//@ts-check
import { Chart } from "chart.js";
/**
 *
 * @param {Chart<"bar", number[], string> | undefined} chart
 */
export default (chart) => {
  const eventSource = new EventSource(`${window.location.origin}/event-stream`);

  eventSource.onmessage = (event) => {
    const newDonation = JSON.parse(event.data);

    //update the graph

    let updateIndex = -1;
    switch (newDonation.Bootcamp) {
      case "web":
        updateIndex = 0;
        break;
      case "ux":
        updateIndex = 1;
        break;
      case "data":
        updateIndex = 2;
        break;
    }
    if (chart != undefined) {
      chart.data.datasets[0].data[updateIndex] += parseFloat(
        newDonation.Amount
      );
      chart.update();
    }
    //update the total ammount
    const potAmount = document.getElementById("potAmount");
    if (potAmount) {
      const newAmount =
        parseFloat(potAmount.innerText) + parseFloat(newDonation.Amount);
      potAmount.innerText = newAmount.toFixed(2);
    }
    //update the list
    const donatinoList = document.getElementById("listOfDontaions");
    const li = document.createElement("li");

    let colorName = "";

    switch (newDonation.Bootcamp) {
      case "web":
        colorName = "text-yellow-300";
        break;
      case "ux":
        colorName = "text-purple-400";
        break;
      case "data":
        colorName = "text-green-300";
        break;
    }

    const listElement = `<li class="text-white w-full grid grid-cols-4 items-center gap-2 mb-2">
<div class="col-span-1">
  <span class="${colorName}">${newDonation.Name}</span>: 
</div>

<div class="text-center col-span-1">${newDonation.Amount}€</div>
  <div class="col-span-2">${newDonation.Message}</div>
</li>`;

    li.innerHTML = listElement;
    donatinoList?.prepend(li);
  };
};
