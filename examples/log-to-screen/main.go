package main

import log "github.com/jlentink/yaglogger"

func main() {
	log.GetInstance().LogToScreen = true
	log.SetLevelByString("debug")
	log.GetInstance().ShowLogLocation = true
	log.SetLevelByString("all")
	log.GetInstance().LogFilePath = "./yaglogger.log"
	// log.SetLevel(log.Info())
	log.Error("Error")
	log.Warn("Warn")
	log.Debug("Debug")
	log.Info("Info")
	log.Fatal("Fatal")

	// log.PrintMessagef("This is a message")
	log.GetInstance().SetDebug(false)
	// log.PrintDebugMessagef("This is a debug message")
	log.GetInstance().SetDebug(true)
	// log.PrintDebugMessagef("This is a debug message")

	color := log.GetInstance().GetColours()
	log.Print(color.Green("This is a green message"))
	log.Print("This is a white message")
	log.Print(1)
	log.Print(1.1)
	log.Info("Info")
	log.Fatal("dsdsds")
}
