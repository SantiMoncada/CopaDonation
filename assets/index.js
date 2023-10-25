//date for November 9, 2023, 4:00 PM spain in UTC
var partyDate = new Date(Date.UTC(2023, 10, 9, 14, 0, 0));

const SECOND = 1000;
const MINUTE = SECOND * 60;
const HOUR = MINUTE * 60;
const DAY = HOUR * 24;

function updateTimer() {
    var now = new Date();
    var distance = partyDate.getTime() - now.getTime();
    if (distance < 0) {
        clearInterval(timer);
        const cd = document.getElementById("countdown");

        if (cd === null) return;

        cd.innerHTML = "PARTY!";
        return;
    }

    const days = Math.floor(distance / DAY);
    const hours = Math.floor((distance % DAY) / HOUR);
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

    cd.innerText = `${days} ${hours}:${minutesLocal}:${secondsLocal}`;
}

const timer = setInterval(updateTimer, 1000);
updateTimer();