package main

func main() {
	kbs := readParse()
	generate(kbs)
}

type keybinding struct {
	key string
	cmd string
}
