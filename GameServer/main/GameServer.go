package main

import (
	C "code.google.com/p/sgemu/Core"
	D "code.google.com/p/sgemu/Data"
	. "code.google.com/p/sgemu/GameServer"
	"code.google.com/p/sgemu/SG"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	//"WebAdmin"
)

var (
	Closing = false
)

func main() {
	defer OnClose()

	runtime.GOMAXPROCS(5)

	log.SetFlags(log.Ltime | log.Lshortfile)
	log.SetPrefix("[Log]")

	D.InitializeDatabase()
	D.CreateDatabase()

	D.LoadData()

	SG.ReadConfig()

	Server = new(GServer)
	Server.Start("GameServer", SG.Config.GSConfig.IP, SG.Config.GSConfig.Port, SG.Config.GSConfig.WANIP)

	//go WebAdmin.Start(GS.Server.Log) 
	go ListenSignals()

	CMD()
}

func ListenSignals() {
	signals := make(chan os.Signal, 10)
	signal.Notify(signals)
	for sig := range signals {
		if sig == os.Interrupt || sig == os.Kill {
			OnClose()
			return
		}
	}
}

func OnClose() {
	if Closing {
		return
	}
	Closing = true

	if x := recover(); x != nil {
		log.Printf("%v %s\n", x, C.PanicPath())
	}

	defer func() {
		if x := recover(); x != nil {
			log.Printf("%v %s\n", x, C.PanicPath())
		}
		cmd := ""
		log.Println("Press enter to quit...")
		fmt.Scanln(&cmd)
		os.Exit(0)
	}()

	if Server != nil {
		Server.OnShutdown()
	}
}

func CMD() {
	for {
		cmd := ""
		fmt.Scanln(&cmd)
		switch cmd {
		case "exit":
			return
		case "cleardb":
			D.ClearDatabase()
			log.Println("Database has been cleared!")
		default:
			if Closing {
				return
			}
		}
	}
}
