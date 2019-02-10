package log

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	Debug *log.Logger = log.New(os.Stderr,
		"DEBUG:\t",
		log.Ldate|log.Ltime)
	Info *log.Logger = log.New(os.Stdout,
		"INFO:\t",
		log.Ldate|log.Ltime)
	Warn *log.Logger = log.New(os.Stdout,
		"WARN:\t",
		log.Ldate|log.Ltime)
	Error *log.Logger = log.New(os.Stdout,
		"ERROR:\t",
		log.Ldate|log.Ltime)
)

func Init(logLevel string) {
	switch logLevel {
	case "debug":
		return
	case "info":
		Debug.SetOutput(ioutil.Discard)
	case "warn":
		Debug.SetOutput(ioutil.Discard)
		Info.SetOutput(ioutil.Discard)
	case "error":
		Debug.SetOutput(ioutil.Discard)
		Info.SetOutput(ioutil.Discard)
		Warn.SetOutput(ioutil.Discard)
	default:
		log.Fatal("Invalid log level " + logLevel)
	}
}
