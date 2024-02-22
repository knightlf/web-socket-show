package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
	"web-socket-show/conf"
	"web-socket-show/pkg/db"
)

// ConfigInit 配置文件初始化
func ConfigInit() {
	if err := conf.ParserConfig("/conf/tsconfig.json"); err != nil {
		fmt.Printf("Error: %#v\n", err)
		//panic(err)
	}

}

//make id of the Msg struct
func makeid() int64 {
	ts := time.Now().Unix()
	return ts
}

//var port *int = flag.Int("p", 9070, "Port to listen.")

// copyServer echoes back messages sent from client using io.Copy.
func copyServer(ws *websocket.Conn) {
	fmt.Printf("copyServer %#v\n", ws.Config())

	io.Copy(ws, ws)
	//var buf string
	//err := websocket.Message.Send(ws, buf)
	//if err != nil {
	//	fmt.Println(err)
	//}
	msg := db.Msg{Id: time.Now().Unix(),
		Select: "copyServer",
		Send:   "&websocket.Dial()"}
	errinsert := db.InsertOper(msg)
	if errinsert != nil {
		fmt.Println(errinsert)
	}

	fmt.Println("copyServer finished")
}

// readWriteServer echoes back messages sent from client using Read and Write.
func readWriteServer(ws *websocket.Conn) {
	fmt.Printf("readWriteServer %#v\n", ws.Config())
	for {
		buf := make([]byte, 100)
		// Read at most 100 bytes.  If client sends a message more than
		// 100 bytes, first Read just reads first 100 bytes.
		// Next Read will read next at most 100 bytes.
		n, err := ws.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("recv:%q\n", buf[:n])
		// Write send a message to the client.
		n, err = ws.Write(buf[:n])
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("send:%q\n", buf[:n])
		msg := db.Msg{Id: time.Now().Unix(),
			Select: "readWriteServer",
			Send:   string(buf[:n])}
		errinsert := db.InsertOper(msg)
		if errinsert != nil {
			fmt.Println(errinsert)
		}
	}
	fmt.Println("readWriteServer finished")

}

// sendRecvSell
//rver echoes back text messages sent from client
// using websocket.Message.
func sendRecvServer(ws *websocket.Conn) {
	fmt.Printf("sendRecvServer %#v\n", ws)
	for {
		var buf string
		// Receive receives a text message from client, since buf is string.
		err := websocket.Message.Receive(ws, &buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("recv:%q\n", buf)
		// Send sends a text message to client, since buf is string.
		err = websocket.Message.Send(ws, buf)
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Printf("send:%q\n", buf)
		msg := db.Msg{Id: time.Now().Unix(),
			Select: "sendRecvServer",
			Send:   buf}
		errinsert := db.InsertOper(msg)
		if errinsert != nil {
			fmt.Println(errinsert)
		}
	}
	fmt.Println("sendRecvServer finished")
}

// sendRecvBinaryServer echoes back binary messages sent from clent
// using websocket.Message.
// Note that chrome supports binary messaging in 15.0.874.* or later.
func sendRecvBinaryServer(ws *websocket.Conn) {
	fmt.Printf("sendRecvBinaryServer %#v\n", ws)
	for {
		var buf []byte
		// Receive receives a binary message from client, since buf is []byte.
		err := websocket.Message.Receive(ws, &buf)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("recv:%#v\n", buf)
		// Send sends a binary message to client, since buf is []byte.
		err = websocket.Message.Send(ws, buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("send:%#v\n", buf)
		msg := db.Msg{Id: time.Now().Unix(),
			Select: "sendRecvBinaryServer",
			Send:   string(buf)}
		errinsert := db.InsertOper(msg)
		if errinsert != nil {
			fmt.Println(errinsert)
		}
	}
	fmt.Println("sendRecvBinaryServer finished")
}

type T struct {
	Msg  string
	Path string
}

// jsonServer echoes back json string sent from client using websocket.JSON.
func jsonServer(ws *websocket.Conn) {
	fmt.Printf("jsonServer %#v\n", ws.Config())
	for {
		var msg T
		// Receive receives a text message serialized T as JSON.
		err := websocket.JSON.Receive(ws, &msg)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("recv:%#v\n", msg)
		// Send send a text message serialized T as JSON.
		err = websocket.JSON.Send(ws, msg)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("send:%#v\n", msg)
		msgCk := db.Msg{Id: time.Now().Unix(),
			Select: "jsonServer",
			Send:   msg.Msg + " $$$ " + msg.Path}
		errinsert := db.InsertOper(msgCk)
		if errinsert != nil {
			fmt.Println(errinsert)
		}
	}
}

//web sever main function.

func MainServer(w http.ResponseWriter, req *http.Request) {
	//serverPort, _ := strconv.Atoi(conf.Cfg.Sever.Sport)
	serverPort := conf.Cfg.Sever.Sport
	serverAddr := conf.Cfg.Sever.Saddr
	webstr := `<html>
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
   ws = new WebSocket("ws://(((svrip)))" + path);
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
<option value="/sendRecvBlob">/sendSun1</option>
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
`
	//(((svrip))):(((svrport)))  --->  serverAddr:serverPort
	svrip := serverAddr + ":" + serverPort
	//strings.Replace(webstr, "(((svrip)))", svrip, -1)
	//fmt.Printf(webstr, serverAddr, serverPort)
	//fmt.Printf(webstr, "192.168.24.20", serverPort)
	//io.WriteString(w, fmt.Sprintf(webstr, serverAddr, serverPort))
	io.WriteString(w, strings.Replace(webstr, "(((svrip)))", svrip, -1))
}

//test server main
func main() {
	flag.Parse()
	ConfigInit()
	//serverPort, _ := strconv.Atoi(conf.Cfg.Sever.Sport)
	//serverAddr := conf.Cfg.Sever.Saddr
	webPort, _ := strconv.Atoi(conf.Cfg.Web.Wport)
	webAddr := conf.Cfg.Web.Waddr

	http.Handle("/copy", websocket.Handler(copyServer))
	http.Handle("/readWrite", websocket.Handler(readWriteServer))
	http.Handle("/sendRecvText", websocket.Handler(sendRecvServer))
	http.Handle("/sendRecvArrayBuffer", websocket.Handler(sendRecvBinaryServer))
	http.Handle("/sendRecvBlob", websocket.Handler(sendRecvBinaryServer))
	http.Handle("/json", websocket.Handler(jsonServer))
	http.HandleFunc("/", MainServer)
	//println(serverAddr)
	//fmt.Printf("http://%v:%v/\n", serverAddr, serverPort)
	fmt.Printf("http://%v:%v/\n", webAddr, webPort)
	err := http.ListenAndServe(fmt.Sprintf("%v:%v", webAddr, webPort),
		nil)
	if err != nil {
		panic("ListenANdServe: " + err.Error())
	}
}
