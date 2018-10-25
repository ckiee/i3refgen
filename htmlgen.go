package main

import (
	"strings"
	"fmt"
	"io/ioutil"
	// "log"
	"os"
)

var cmdHuman = map[string]string{
	"workspace": "Switch to workspace",
	"exec": "Execute command",
	"split": "Split containers",
	"layout": "Change container layout",
	"focus": "Focus a container",
	"move": "Move container",
	"swap": "Swap",
	"sticky": "Sticky floating windows",
	"nop": "Do nothing",
	"restart": "Restart i3wm in place",
	"reload": "Reload i3wm",
	"exit": "Exit your X session",
	"kill": "Close current window",
	"fullscreen": "Fullscreen window",
	"floating": "Floating window mode",
	"mode": "Window mode",
}

func generate(kbs []keybinding) {
	str := ""
	for _, kb := range kbs {
		rcmds := strings.Split(kb.cmd, string(0x20))
		rcmd := rcmds[1];
		cmd := "<b>"+cmdHuman[rcmd] + "</b>" + strings.TrimPrefix(kb.cmd, rcmds[0]+" "+rcmd)
		str += fmt.Sprintf("<tr><td><kbd>%s</kbd></td><td>%s</td></tr>", kb.key, cmd)
	}
	ioutil.WriteFile("out.html", []byte(template(str)), os.FileMode(int(0777)));
	// log.Println(template(str));
}

func template(str string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
	<html>
	<head>
		<meta charset="utf-8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<title>i3 Reference Card (i3refgen)</title>
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<style>
			h2 {
				margin: 7px 0 2px;
				padding: 2px 4px;
				font-size: 1.1em;
				font-family: sans;
				background-color: #b3b3b3;
			}
			kbd {
				/* border-spacing: 2px; */
				background: rgba(128, 128, 128, 0.596);
				padding: 5px;
			}
		</style>
	</head>
	<body>
		<header>
			<h1>i3 Reference Card (i3refgen)</h1>
			<p>You can print this file and put it on your desk.</p>
		</header>
		<section>
			<h2>Keybindings</h2>
			<br>
			<table>
				<tbody>
					%s
				</tbody>
			</table>
		</section>
	</body>
	</html>`, str)
}