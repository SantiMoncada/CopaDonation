//@ts-check
import { Chart } from "chart.js";
/**
 *
 * @param {Chart<"bar", number[], string> | undefined} chart
 */
export default (chart) => {
  let firstVisit = true;
  let i = 0;

  /**
   * @param {String} bootcamp
   * @param {String} name
   * @param {String} amount
   * @param {String} message
   * @returns {HTMLLIElement}
   */
  const generateDonationRow = (bootcamp, name, amount, message) => {
    const li = document.createElement("li");

    let colorName = "";

    switch (bootcamp) {
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

    const listElement = `
<li class="grid grid-cols-4 gap-2 mb-4">
  <div class="col-span-1 break-words">
    <span class="${colorName}">${name}</span>
  </div>

  <div class="text-white text-center col-span-1">${parseFloat(amount).toFixed(
    2
  )}€</div>

  <div class="text-white col-span-2">${message}</div>
</li>
`;
    li.innerHTML = listElement;
    return li;
  };

  window.onfocus = () => {
    if (!firstVisit) {
      fetch(`${window.location.origin}/api/data`)
        .then((data) => data.json())
        .then((response) => {
          const { donations, total, uxTotal, webTotal, dataTotal } = response;

          //update the total
          const pot = document.getElementById("potAmount");
          if (pot != null) {
            pot.innerText = parseFloat(total).toFixed(2);
          }

          //update the chart
          if (chart != null) {
            chart.data.datasets[0].data = [webTotal, uxTotal, dataTotal];
          }

          //replace the list
          const donatinoList = document.getElementById("listOfDontaions");
          const donationsRows = donations.map((donation) =>
            generateDonationRow(
              donation.Bootcamp,
              donation.Name,
              donation.Amount,
              donation.Message
            )
          );
          if (donatinoList != null) {
            donatinoList.replaceChildren(...donationsRows);
          }
        });
    }

    firstVisit = false;
  };

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
    const li = generateDonationRow(
      newDonation.Bootcamp,
      newDonation.Name,
      newDonation.Amount,
      newDonation.Message
    );

    donatinoList?.prepend(li);
  };
};
