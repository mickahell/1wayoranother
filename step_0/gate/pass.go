package gate

import (
	"step-0/helpers"
	"step-0/message"
)

func PassOption() error {
	password := StringPrompt("Please enter the password : ")
	forget_pass := "no"

	if password == helpers.TheAppConfig().RouterPass {
		message.GetMessage().MesInfo("You are connected!")
		helpers.TheAppConfig().IsConnected = true
	} else {
		message.GetMessage().MesWarning("Wrong password!")
		forget_pass = StringPrompt("Forget my password ? yes, no : ")
	}

	if forget_pass == "yes" {
		err = ForgetPass()
	}

	if err != nil {
		return err
	}

	return nil
}

func ForgetPass() error {
	const img = `
	~The hyper router
	vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
	vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvuvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
	vvvvvxfn(vvutvvnfrfrvv\\/)t/vvvtv/vvu(/|/\/(j{|/(x\nvv\uvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
	vvvvvnjvxvvuruvununuvvfft/r\vvvfvfvvutfjrffxrjrrfnxvvvruvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
	vvvvvrvvvvvvvvjuvvxvvvjjvvvxvvvxvfvvvnvrrrrvvvvvvvvvvuxvvvjurnvvvvvvvvvvvxvjxxvvvvvvvvvvvv
	vvvvvjvvvvvvvu}|})}rfn{/)fr})1x]u)f)|fvxx\t1{1{j|uvn]t))[1tfj/vv}}|-t]x]r|vrnrnvvvvvvvvvvv
	vvvvvrvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvuvvvvvvvvvvvvvvvvvvvvvvvvvvjvvvvvvvvvvvvuvvvvvvvvvvvv
	vvvvvjrvvnrutururnvvrxvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
	vvvvvrjvvftutn/r/xvvjxvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
	vvvvvxvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvuvvvvvvvvvvvvvvvvvvuvuvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
	vvvvvjvvvvvvvf-t]j[f}1{/uu{(r\x|f/t(xfvtt(|\\\/vvr(f(t(t)v|xvv\\\\r)vvr\/(t(r(frjttvfvvvvv
	vvvvvjvvvvvvvn\xfvfvtffvrxrjnnunufnxvnvvvnurjrfvvnfnnrfrxxurvvxxrfvfvvr/xfrjxjrnnvvunvvvvv
	vvvvvrnvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
	vvvvvrtvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
	vvvvvuvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
	`

	message.GetMessage().MesDraw(img)

	return nil
}
