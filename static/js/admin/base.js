$(function(){
    baseApp.init();
})

var baseApp={
    init: function(){
        this.initAside();
        this.confirmDelete();
    },
    initAside: function(){
        $('.aside h4').onclick(function(){		
			$(this).siblings('ul').slideToggle();
		});
    },
	confirmDelete: function() {
		$(".delete").onclick(function(){
			var flag = confirm("您确定要删除吗？");
			return flag;
		})
	}
}