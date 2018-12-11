package main

import (
	"bytes"
	"io"
	"log"
)

func ConsumeMssg() {
	var err error
	Mssg, err = Ch.Consume(
		"PostQ",
		"",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}
