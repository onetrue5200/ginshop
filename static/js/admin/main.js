export function main() {
    changeStatus();
}

function changeStatus() {
    $(".chStatus").click(function(){
        var id = $(this).attr("data-id");
        var table = $(this).attr("data-table");
        var field = $(this).attr("data-field");
        var el = $(this);
        $.get("/admin/changeStatus", {
            "id": id,
            "table": table,
            "field": field,
        }, function(resp){
            if (resp.success) {
                if (el.attr("src").indexOf("yes") != -1) {
                    el.attr("src", "/static/images/admin/focus/no.gif");
                } else {
                    el.attr("src", "/static/images/admin/focus/yes.gif");
                }
            }
        })
    });
}