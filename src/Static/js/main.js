/* window.onload = function () {
	$('.file').css('opacity', '1');
	//动画
	anime({
		targets: '.file',
		translateY: 60,
		duration: 500
	});
} */

//获取url参数
function getQueryVariable(variable) {
	var query = window.location.search.substring(1);
	var vars = query.split("&");
	for (var i = 0; i < vars.length; i++) {
		var pair = vars[i].split("=");
		if (pair[0] == variable) {
			return pair[1];
		}
	}
	return ("");
}
//当前目录
var FilePath = "";
if (decodeURIComponent(getQueryVariable("dir")) != null) {
	FilePath = decodeURIComponent(getQueryVariable("dir"));
}

//前往私人空间
function toUsrSpace() {
	var UsrSpace = prompt("请输入私人标识", "");
	if (UsrSpace != null && IsIllegal(UsrSpace)) {
		ToDir(UsrSpace)
	} else {
		alert("输入违规!");
	}
}

//前往子文件夹
function ToDir(folderName) {
	window.open("./?dir=" + FilePath + "/" + folderName);
}

var fileid = 0;
var fileButton = new Array();
//处理文件
function Upload(files) {
	if (files.length === 0) {
		alert('请选择文件！');
		return;
	}
	handle(files, 0, files.length);
}

function handle(files, i, ele) {
	//获取文件类型
	var prefix = files[i].name.substring(files[i].name.indexOf(".") + 1);
	if (prefix == "") {
		prefix = "file"
	}
	//加入数据
	let data = new FormData();
	data.append('file', files[i]);
	data.append('dir', FilePath)
	$('#WebBody').append("<div class='item'>" +
		"		<img src='./Static/img/icons/" + prefix + ".svg'>" +
		"		<span>" + files[i].name + "</span>" +
		"		<div class='buttonArea' id='UploadFiles" + i +
		"'><span></span></div></div>");
	fileButton[i] =
		"		<div class='download' onclick='Download(\"" + FilePath + "/" + files[i].name + "\")'>下载</div>" +
		"		<div class='delete' onclick='Delete(\"" + FilePath + "/" + files[i].name + "\")'>删除</div>";
	push(data, i, fileid);
	fileid++;
	if (i + 1 < ele) {
		handle(files, i + 1, ele);
	}
}

//上传数据
function push(data, i, id) {
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
				var table = '#UploadFiles' + id + ' span';
				$(table).html(progressRate);
			})
			return xhr;
		}
	}).done(function (output) {
		if (output !== "done") {
			alert("第" + (i + 1).toString() + "个文件上传失败！请稍候再试！");
		} else {
			var table = '#UploadFiles' + id;
			$(table).html(fileButton[id]);
		}
	}).fail(function (xhr, status) {
		console.log(status);
	});
}

//新建文件夹
function NewFolder() {
	var folderName = prompt("请输入文件夹名", "");
	if (folderName != null && IsIllegal(folderName)) {
		$.ajax({
			url: "./?dir=" + FilePath + "/" + folderName,
			type: "get",
		}).done(function () {
			location.reload();
		}).fail(function (xhr, status) {
			console.log(status);
		});
	} else {
		alert("输入违规!");
	}
}

//下载数据
function Download(filePath) {
	window.open("./Download?dir=" + filePath);
}

//删除数据
function Delete(filePath) {
	window.event.cancelBubble = true;
	$.ajax({
		url: "./Delete?dir=" + filePath,
		type: "get",
	}).done(function () {
		location.reload();
	}).fail(function (xhr, status) {
		console.log(status);
	});
}

//判断是否名称违法,违法返回false
function IsIllegal(str) {
	var patt = /^[\a-zA-Z0-9_\u4e00-\u9fa5]+$/;
	return patt.test(str);
}