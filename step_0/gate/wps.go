package gate

import (
	"step-0/helpers"
	"step-0/message"
	"time"
)

func WPSOption() error {
	mes := "#"
	for count := 0; count <= 30; count += 1 {
		message.GetMessage().MesOnline(mes)
		time.Sleep(200 * time.Millisecond)

		mes += "#"
		count += 1
	}
	message.GetMessage().MesInfo("\nYou are connected! Who need password those days!")
	helpers.TheAppConfig().IsConnected = true
	return nil
}
