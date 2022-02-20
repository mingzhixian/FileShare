//处理文件
function handleFiles(files) {
	if (files.length === 0) {
		alert('请选择文件！');
		return;
	}
	handle(files, 0, files.length);
}

function handle(files, i, ele) {
	let data = new FormData();
	data.append('file', files[i]);
	push(data, i);
	if (i + 1 < ele) {
		handle(files, i + 1, ele);
	} else if (i + 1 === ele) {
		window.open("./");
	}
}
//上传数据
function push(data, i) {
	$.ajax({
		url: "./Upload",
		type: "post",
		data: data,
		mimeType: "multipart/form-data",
		contentType: false,
		cache: false,
		processData: false
	}).done(function (output) {
		if (output !== "done") {
			alert("第" + (i + 1).toString() + "个文件上传失败！请稍候再试！");
		}
	}).fail(function (xhr, status) {
		console.log(status);
	});
}