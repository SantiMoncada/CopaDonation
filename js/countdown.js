//@ts-check
export default () => {
  //date for November 10, 2023, 4:00 PM spain in UTC
  var partyDate = new Date(Date.UTC(2023, 10, 10, 14, 0, 0));

  const SECOND = 1000;
  const MINUTE = SECOND * 60;
  const HOUR = MINUTE * 60;
  const DAY = HOUR * 24;

  const timer = setInterval(updateTimer, 1000);

  updateTimer();

  function updateTimer() {
    var now = new Date();
    var distance = partyDate.getTime() - now.getTime();
    if (distance < 0) {
      const cd = document.getElementById("countdown");

      if (cd === null) return;

      cd.innerHTML = "PARTY!";

      clearInterval(timer);
      return;
    }

    const hours = Math.floor(distance / HOUR);
    const minutes = Math.floor((distance % HOUR) / MINUTE);
    const seconds = Math.floor((distance % MINUTE) / SECOND);

    const cd = document.getElementById("countdown");
    if (cd === null) return;

    const minutesLocal = minutes.toLocaleString(undefined, {
      minimumIntegerDigits: 2,
    });

    const secondsLocal = seconds.toLocaleString(undefined, {
      minimumIntegerDigits: 2,
    });

    cd.innerText = `${hours}:${minutesLocal}:${secondsLocal}`;
  }
};
