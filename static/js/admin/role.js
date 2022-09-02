$('.delete').on('click', function (e) {
    var flag = confirm("您确定要删除吗？");
    return flag;
});

$('#photoUploader').diyUpload({
    url: '/admin/goods/imageUpload',
    success: function (response) {
        var photoStr = '<input type="hidden" name="goods_image_list" value=' + response.link + ' />';
        $("#photoList").append(photoStr)
    },
    error: function (err) {
        console.info(err);
    }
});

$('.relation_goods_color').change(function(){
    var goods_image_id=$(this).attr("goods_image_id");
    var color_id=$(this).val();
    $.get("/admin/goods/changeGoodsImageColor",{"goods_image_id":goods_image_id,"color_id":color_id},function(response){
    });
})

$(".goods_image_delete").click(function(){
    var goods_image_id=$(this).attr("goods_image_id")   
    var _that=this;         
    var flag = confirm("确定要删除吗?");
    if (flag){
        $.get("/admin/goods/removeGoodsImage",{"goods_image_id":goods_image_id},function(response){
            if(response.success){
                $(_that).parent().remove()
            }
        })
    }
})