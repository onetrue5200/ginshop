$('.delete').on('click', function (e) {
    var flag = confirm("您确定要删除吗？");
    return flag;
})

$('#photoUploader').diyUpload({
    url: '/admin/goods/imageUpload',
    success: function (data) {
        console.info(data);
    },
    error: function (err) {
        console.info(err);
    }
});