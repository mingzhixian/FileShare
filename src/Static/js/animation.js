window.onload = function () {
	$('.file').css('opacity', '1');
	//动画
	anime({
		targets: '.file',
		translateY: 60,
		duration: 500
	});
}