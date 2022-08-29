export function main() {
    initAside();
    changeStatus();
    changeNum();
}

function initAside(){
    $(".aside h4").click(function(){		
        $(this).siblings('ul').slideToggle();
    });
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

function changeNum() {
    $(".chSpanNum").click(function(){
        var id=$(this).attr("data-id");
        var table=$(this).attr("data-table");
        var field=$(this).attr("data-field");
        var num=$(this).html();
        
        var SpanEl = $(this)

        var input = $("<input style='width:40px' />");
        $(this).html(input);
        $(input).trigger("focus");
        $(input).val(num);
        $(input).click(function(e){
            e.stopPropagation();
        });
        $(input).blur(function(){
            var inputNum = $(this).val();
            SpanEl.html(inputNum);
            $.get("/admin/changeNum", {
                "id": id,
                "table": table,
                "field": field,
                "num": inputNum,
            }, function(resp){
            })  
        })
    })
}