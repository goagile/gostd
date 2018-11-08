

$(document).ready(() => {
    $("#search_posts").click(() => {
        let searchstring = $("#search_string");
        let req = {
            "searchstring": searchstring.val(),
        }
        // $.post("/search/", req, (data, status) => {
        //     alert(data, status)
        // });

        var xhr = new XMLHttpRequest();
        xhr.open('POST', /search/, true);
        xhr.setRequestHeader('Content-type', 'application/json');
        xhr.onload = function () {
            alert(this.responseText);
        };
        xhr.send(req);
    });
});
