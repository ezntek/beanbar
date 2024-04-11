package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/lestrrat-go/strftime"
)

const PREFIX = "["
const POSTFIX = "]"
const SEPARATOR = " | "

type Widget interface {
	Deinit()
	Display() string
}

type DateTime struct {
	format string
}

func NewDateTime(format string) *DateTime {
	return &DateTime{
		format: format,
	}
}

func (d *DateTime) Deinit() {
	return
}

func (d *DateTime) Display() string {
	var buf bytes.Buffer

	res, err := strftime.Format(d.format, time.Now())

	if err != nil {
		panic("failed to format the current time")
	}

	buf.WriteString(PREFIX)
	buf.WriteString(res)
	buf.WriteString(POSTFIX)

	return buf.String()
}

func main() {
	var buf bytes.Buffer
	WIDGETS := []Widget{
		NewDateTime("%A, %d %D"),
		NewDateTime("%H:%M:%S"),
	}

	for index, widget := range WIDGETS {
		buf.WriteString(widget.Display())

		if index != len(WIDGETS)-1 {
			buf.WriteString(SEPARATOR)
		}
	}

	fmt.Println(buf.String())
}
