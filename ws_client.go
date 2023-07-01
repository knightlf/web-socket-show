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
   ws = new WebSocket("ws://0.0.0.0:9055" + path);
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
<form name="msgform" action="#" onsubmit="return send();">
<select onchange="init()" name="path">
<option value="/copy" selected="selected">/copy</option>
<option value="/readWrite">/readWrite</option>
<option value="/sendRecvText">/sendRecvText</option>
<option value="/sendRecvArrayBuffer">/sendRecvArrayBuffer</option>
<option value="/sendRecvBlob">/sendRecvBlob</option>
<option value="/json">/json</option>
</select>
<input type="text" name="message" size="80" value="">
<input type="file" name="file" >
<input type="submit" value="send">
</form>
<div id="msg"></div>
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
