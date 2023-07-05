package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
)

var port *int = flag.Int("p", 9033, "Port to listen.")

//http://10.1.1.13:9080/ws/
//10.1.1.14:9197/ws
//ws://101.251.223.186:9197
func MainServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `<html>
<head>
<style type="text/css">
:root{
	--border-color: #ccc;
}
*,body{
	margin:0;
	padding: 0;
	outline:none;
}
body{
	background: radial-gradient(circle at 7.2% 13.6%, rgb(37, 249, 245) 0%, rgb(8, 70, 218) 90%);
}
.box-item{
	padding: 20px;
	background: rgba(255,255,255,0.5);
	border-radius: 5px;
	box-sizing:border-box;
}
.container{
	width: 860px;
	margin: 0 auto;
	padding: 40px 0;
	display: flex;
	flex-direction: column;
	height: 100vh;
	box-sizing: border-box;
}
.my-form{
	display: flex;
}
.input-wrap{
	display: block;
	flex: 1;
	margin-right: 10px;
	height: 40px;
	line-height: 40px;
	border: 1px solid var(--border-color);
	border-radius: 5px;
	position: relative;
	padding-left: 168px;
	padding-right: 80px;
	overflow: hidden;
	background: #fff;
}
.select{
	position: absolute;
	left: 0;
	top: 0;
	height: 40px;
	border: none;
	padding: 0 0px;
	box-sizing: border-box;
	display: block;
	width: 168px;
	text-align: center;
	font-size: 14px;
	color:#22c7d5;
	border-right: 1px solid var(--border-color);
}
.input-text{
	display: block;
	width: 100%;
	height: 100%;
	border: none;
	padding: 0 10px;
	font-size: 16px;
}
.input-file-wrap{
	position: absolute;
	right: 0;
	top: 0;
	width: 80px;
	font-size: 14px;
	cursor: pointer;
	height: 100%;
	text-align: center;
	border-left: 1px solid var(--border-color);
	color: #409eff;
}
.input-file-wrap input{
	opacity: 0;
	position: absolute;
	width: 200%;
	height: 100%;
	top: 0;
	left: -70px;
	cursor: pointer;
	z-index: 9;
	display: block;
	color: transparent;
}
.btn-send{
	width: 100px;
	height: 40px;
	border: none;
	color: #fff;
	font-size: 16px;
	border-radius: 5px;
	background: #409eff;
	cursor:pointer;
}
.btn-send:hover,.input-file-wrap:hover{
	opacity:0.8;
}
.msg-container{
	margin-top: 20px;
	min-height: 360px;
	font-size: 15px;
	line-height: 1.5em;
	color: #333;
	letter-spacing: 1px;
	height: 500px;
	flex: 1;
	overflow: auto;
}
.msg-wrap{
	min-height:200px;
}
</style>
<script type="text/javascript">
var path;
var ws;
function init() {
   console.log("init");
   if (ws != null) {
     ws.close();
     ws = null;
   }
   path = document.msgform.path.value;
   console.log("path:" + path);
   var div = document.getElementById("msg");
   div.innerText = "path:" + path + "\n" + div.innerText;
   ws = new WebSocket("ws://%v:%v" + path);
   if (path == "/sendRecvBlob") {
     ws.binaryType = "blob";
   } else if (path == "/sendRecvArrayBuffer") {
     ws.binaryType = "arraybuffer";
   }
   ws.onopen = function () {
      div.innerText = "opened\n" + div.innerText;
   };
   ws.onmessage = function (e) {
      div.innerText = "msg:" + e.data + "\n" + div.innerText;
      if (e.data instanceof ArrayBuffer) {
        s = "ArrayBuffer: " + e.data.byteLength + "[";
        var view = new Uint8Array(e.data);
        for (var i = 0; i < view.length; ++i) {
          s += " " + view[i];
        }
        s += "]";
        div.innerText = s + "\n" + div.innerText;
      }
   };
   ws.onclose = function (e) {
      div.innerText = "closed\n" + div.innerText;
   };
   console.log("init");
   div.innerText = "init\n" + div.innerText;
};
function send() {
   console.log("send");
   var m = document.msgform.message.value;
   if (path == "/sendRecvArrayBuffer" || path == "/sendRecvBlob") {
     var t = m;
     if (t != "") {
       var array = new Uint8Array(t.length);
       for (var i = 0; i < t.length; i++) {
          array[i] = t.charCodeAt(i);
       }
       m = array.buffer;
     } else {
     m = document.msgform.file.files[0];
     }
   } else if (path == "/json") {
     m = JSON.stringify({Msg: m, Path: path})
   }
   console.log("send:" + m);
   if (m instanceof ArrayBuffer) {
     var s = "arrayBuffer:" + m.byteLength + "[";
     var view = new Uint8Array(m);
     for (var i = 0; i < m.byteLength; ++i) {
      s += " " + view[i];
     }
     s += "]";
     console.log(s);
   }
   ws.send(m);
   return false;
};
</script>
<body onLoad="init();">
<div class="container">
<form name="msgform" action="#" onsubmit="return send();" class="my-form box-item">
<div class="input-wrap">
<select class="select" onchange="init()" name="path">
<option value="/copy" selected="selected">/copy</option>
<option value="/readWrite">/readWrite</option>
<option value="/sendRecvText">/sendRecvText</option>
<option value="/sendRecvArrayBuffer">/sendRecvArrayBuffer</option>
<option value="/sendRecvBlob">/sendRecvBlob</option>
<option value="/json">/json</option>
</select>
<input class="input-text" type="text" name="message" size="80" value="" />
<div class="input-file-wrap">选择文件<input type="file" name="file" /></div>
</div>
<input class="btn-send" type="submit" value="send" />
</form>
<div class="msg-container">
<div class="msg-wrap box-item" id="msg"></div>
</div>
</div>
</html>
`)
}

//test main
func main() {
	flag.Parse()

	http.HandleFunc("/", MainServer)
	fmt.Printf("http://0.0.0.0:%d/\n", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		panic("ListenANdServe: " + err.Error())
	}
}
