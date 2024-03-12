package message

import (
	"fmt"
	"os"
)

type Message struct {
	drawing string
	level   string
	message string
}

const DrawColor = "\033[96;1m"
const InfoColor = "\033[0;32m"
const WarningColor = "\033[0;33m"
const ErrorColor = "\033[1;31m"
const CriticalColor = "\033[0;31m"
const NoneColor = "\033[0m"
const OnlineColor = "\033[96;1m"

var mymes Message

func (mes *Message) MesDraw(drawing string) {
	mes.drawing = drawing

	fmt.Fprintln(os.Stdout, DrawColor, string(mes.drawing), NoneColor)
}

func (mes *Message) MesOnline(message string) {
	mes.level = "CHOICE"
	mes.message = message

	fmt.Fprint(
		os.Stdout,
		OnlineColor,
		mes.message,
		NoneColor,
	)

}

func (mes *Message) MesInfo(message string) {
	mes.level = "INFO"
	mes.message = message

	fmt.Fprintln(
		os.Stdout,
		InfoColor,
		mes.message,
		NoneColor,
	)

}

func (mes *Message) MesWarning(message string) {
	mes.level = "WARNING"
	mes.message = message

	fmt.Fprintln(
		os.Stderr,
		WarningColor,
		mes.level+" : "+mes.message,
		NoneColor,
	)

}

func (mes *Message) MesError(message string) {
	mes.level = "ERROR"
	mes.message = message

	fmt.Fprintln(
		os.Stderr,
		ErrorColor,
		mes.level+" : "+mes.message,
		NoneColor,
	)
}

func (mes *Message) MesCritical(message string, exit ...bool) {
	mes.level = "CRITICAL"
	mes.message = message

	fmt.Fprintln(
		os.Stderr,
		CriticalColor,
		mes.level+" : "+mes.message,
		NoneColor,
	)
	if len(exit) == 0 {
		os.Exit(1)
	}
}

func GetMessage() *Message {
	return &mymes
}
