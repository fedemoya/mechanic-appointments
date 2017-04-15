function getCurrentTimeInSeconds() {
    return Math.floor($.now() / 1000);
}

function getTimeInSeconds(date) {
    var milliseconds_time = date.getTime();
    return Math.floor(milliseconds_time / 1000);
}

function epochToString(seconds) {
    var milliseconds = seconds * 1000;
    var d =  new Date(milliseconds);
    var month = d.getMonth() + 1;
    return d.getDate() + '/' + month + '/' + d.getFullYear();
}