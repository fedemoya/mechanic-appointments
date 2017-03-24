function getCurrentTimeInSeconds() {
    return Math.floor($.now() / 1000);
}

function getTimeInSeconds(date) {
    var milliseconds_time = date.getTime();
    return Math.floor(milliseconds_time / 1000);
}