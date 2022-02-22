//新建文件夹
function newFiles(){
	var FilesName=prompt("请输入文件夹名");
	if(FilesName!=null&&FilesName!=""){
		$.ajax({
			url: "./Upload?NewFiles="+FilesName,
			type: "get"
		}).done(function (output) {
			if (output !== "done") {
				alert("出现错误，请稍后再试!");
			} else {
				alert("成功!")
			}
		}).fail(function (xhr, status) {
			console.log(status);
		});
	}
}
var fileid = 0;
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
	$('#Uploaded').append("<div id=\"UploadFiles" + fileid + "\" class=\"UploadFiles\">" + "<span class='UploadFilesName'>" + files[i].name + "</span><span class='UploadFilesDone'></span>" + "</div>")
	push(data, i, fileid);
	fileid++;
	if (i + 1 < ele) {
		handle(files, i + 1, ele);
	}
}
//上传数据
function push(data, i,id) {
	$.ajax({
		url: "./Upload",
		type: "post",
		data: data,
		mimeType: "multipart/form-data",
		contentType: false,
		cache: false,
		processData: false,
		xhr: function () {
			var xhr = new XMLHttpRequest();
			//使用XMLHttpRequest.upload监听上传过程，注册progress事件，打印回调函数中的event事件
			xhr.upload.addEventListener('progress', function (e) {
				//loaded代表上传了多少
				//total代表总数为多少
				var progressRate = parseInt((e.loaded / e.total) * 100) + "%";
				var table = '#UploadFiles' + id + ' .UploadFilesDone';
				$(table).html(progressRate);
			})

			return xhr;
		}
	}).done(function (output) {
		if (output !== "done") {
			alert("第" + (i + 1).toString() + "个文件上传失败！请稍候再试！");
		} else {
			var table = '#UploadFiles' + id + ' .UploadFilesDone';
			$(table).html("完成");
		}
	}).fail(function (xhr, status) {
		console.log(status);
	});
}