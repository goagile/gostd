

$(document).ready(() => {
    $("#search_posts").click(() => {
        let searchstring = $("#search_string");
        let req = {
            "searchstring": searchstring.val(),
        }
        $.post("/search/", JSON.stringify(req), (data, status) => {
            let results = JSON.parse(data);
            console.log(results, status);
            if (results == null || results == undefined || !results.length) {
                return
            }
            $("#searchresults").html("");
            results.forEach(r => {
               $("#searchresults").append(`<div>${r.title}</div>`);
            });
        });

        // var xhr = new XMLHttpRequest();
        // xhr.open('POST', /search/, true);
        // xhr.setRequestHeader('Content-type', 'application/json');
        // xhr.onload = function () {
        //     alert(this.responseText);
        // };
        // xhr.send(req);
    });
});
