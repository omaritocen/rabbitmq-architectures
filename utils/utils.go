package utils

import (
	"log"
	"os"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s %s", msg, err)
	}
}

func BodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = args[1]
	}

	return s
}

func SeverityFrom(args []string) string {
	var s string
	if (len(args) < 3) || os.Args[2] == "" {
		s = "info"
	} else {
		s = os.Args[2]
	}
	return s
}
