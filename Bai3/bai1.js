$(document).ready(function() {
    $("#content-a-1").click(function() {
        var str = $("#content-a-1").text();
        if(str === "View more >") {
            $("#content-a-1").text("View less >");
            $("#content-more-1").show();
        } else {
            $("a").text("View more >");
            $("#content-more-1").hide();
        }
    });
    $("#content-a-2").click(function() {
        var str = $("#content-a-2").text();
        if(str === "View more >") {
            $("#content-a-2").text("View less >");
            $("#content-more-2").show();
        } else {
            $("a").text("View more >");
            $("#content-more-2").hide();
        }
    });
});