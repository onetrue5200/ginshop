$(function(){
    LoginApp.init();
})

var LoginApp={
    init: function(){
        this.getCaptcha()
        this.imageChange()
    },
    getCaptcha: function(){
        $.get("/admin/captcha?t="+Math.random(), function(resp){
            $("#captcha-id").val(resp.captchaId)
            $("#captcha-img").attr("src", resp.captchaImage)
        })
    },
    imageChange: function(){
        var outer = this;
        $("#captcha-img").click(function(){
            outer.getCaptcha()
        })
    }
}