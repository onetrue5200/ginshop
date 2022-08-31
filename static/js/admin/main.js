export function main() {
    initAside();
    changeStatus();
    changeNum();
    changeGoodsCate();
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

function changeGoodsCate() {
    $("#goods_type_id").change(function(){
        var cateId = $(this).val();
        $.get("/admin/goods/goodsTypeAttribute", {"cateId": cateId}, function(resp){
            console.log(resp);
            var str = "";
            if (resp.success) {
                var attrData = resp.result;
                for (var i = 0; i < attrData.length; i++) {
                    if (attrData[i].attr_type == 1) {
                        str += '<li><span>' + attrData[i].title + ': </span><input type="hidden" name="attr_id_list" value="'+ attrData[i].id +'" /><input type="text" name="attr_value_list" /></li>'
                    } else if (attrData[i].attr_type == 2) {
                        str += '<li><span>' + attrData[i].title + ': </span><input type="hidden" name="attr_id_list" value="' + attrData[i].id + '"><textarea cols="50" rows="3" name="attr_value_list"></textarea></li>'
                    } else {
                        var attrArray = attrData[i].attr_value.split("\n")
                        str += '<li><span>' + attrData[i].title + ': </span>  <input type="hidden" name="attr_id_list" value="' + attrData[i].id + '" />';
                        str += '<select name="attr_value_list">'
                        for (var j = 0; j < attrArray.length; j++) {
                            str += '<option value="' + attrArray[j] + '">' + attrArray[j] + '</option>';
                        }
                        str += '</select>'
                        str += '</li>'
                    }
                }
                $("#goods_type_attribute").html(str);
            }
        })
    });
}