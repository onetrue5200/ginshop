$('.delete').on('click', function (e) {
    var flag = confirm("您确定要删除吗？");
    return flag;
})

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