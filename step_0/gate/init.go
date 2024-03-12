package gate

import (
	"bufio"
	"fmt"
	"os"
	"step-0/helpers"
	"step-0/message"
	"strconv"
	"strings"
)

var err error

func StringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		message.GetMessage().MesOnline(label)
		fmt.Fprint(os.Stderr, " ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func InitChoice() error {
	for !helpers.TheAppConfig().IsConnected {
		choice := StringPrompt(`
1. WPS
2. Password
3. RJ45
4. Console port
Choose how you want to connect ?`)

		if choice == strconv.Itoa(1) {
			err = WPSOption()
		}
		if choice == strconv.Itoa(2) {
			err = PassOption()
		}
		if choice == strconv.Itoa(3) || choice == strconv.Itoa(4) {
			message.GetMessage().MesWarning("No cable detected. Did you connected the cable ?")
		}

		if err != nil {
			return err
		}
	}

	return nil
}
