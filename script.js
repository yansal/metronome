var interval;
var clickAudio = document.getElementById("click");
var tempoDisplay = document.getElementById("tempo");

var taps = [];
document.addEventListener("touchstart", function (evt) {
    taps = taps.concat(performance.now());
    if (taps.length < 2) {
        return;
    }
    if (taps.length > 4) {
        taps = taps.slice(1);
    }

    var diffs = [];
    for (i = 0; i < taps.length-1; i++) {
        diffs = diffs.concat(taps[i+1]-taps[i]);
    }
    var avgms = avg(diffs);

    clearInterval(interval);
    clickAudio.play();
    interval = setInterval(function(){
        clickAudio.play();
    }, avgms)

    tempoDisplay.textContent = (60000/avgms).toFixed(2);
});

function avg(array) {
    var sum = 0;
    for(var i = 0; i < array.length; i++) {
        sum += array[i];
    }
    return sum / array.length;
}