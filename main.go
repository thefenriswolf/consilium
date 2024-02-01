package main

import (
	"embed"
)

// #+TODO: ro database entries
// #+TODO: save data as json?
// #+TODO: PDF export
// #+TODO: graceful shutdown

//const LANG string = "en"
//const VERSION string = "20240130"

//go:embed resources/*
var resources embed.FS

//var tpl *template.Template
//var name = "john"

const PORT string = ":6477"
const ADDRESS string = "localhost" + PORT

func main() {
	WriteDB("dev.db", "devbucket", "devkey", []byte("devdata"))
	LaunchServer(PORT)
}
