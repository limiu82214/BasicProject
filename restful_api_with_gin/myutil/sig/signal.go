package sig

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

var ServerSigChan chan os.Signal

func init() {
	ServerSigChan = make(chan os.Signal)
}

func ServerNotify() {
	signal.Notify(ServerSigChan, syscall.SIGINT, syscall.SIGTERM)
	<-ServerSigChan
}

func ShutdownServer(err error) {
	log.Println(err)
	ServerSigChan <- os.Interrupt
}
