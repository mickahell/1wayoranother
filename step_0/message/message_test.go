package message

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	//setUp()
	//log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	//log.Println("Do stuff AFTER the tests!")
	//tearDown()
	os.Exit(exitVal)
}

func TestMesInfo(t *testing.T) {
	GetMessage().MesInfo("this is a unit test with info log.")
}

func TestMesWarning(t *testing.T) {
	GetMessage().MesWarning("this is a unit test with warning log.")
}

func TestMesError(t *testing.T) {
	GetMessage().MesError("this is a unit test with errror log.")
}

func TestMesCritical(t *testing.T) {
	GetMessage().MesCritical("this is a unit test with critical log.", false)
}

func TestTheMessage(t *testing.T) {
	var got interface{} = GetMessage()

	_, ok := got.(*Message)

	if ok == false {
		log.Fatalln("got is not what i want !")
	}

}
