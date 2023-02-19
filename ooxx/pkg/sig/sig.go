package sig

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

var serverSigChan_once chan os.Signal //nolint //singletone

func init() { //nolint //whole system will use one of this
	serverSigChan_once = make(chan os.Signal)
}

func ServerNotify() {
	signal.Notify(serverSigChan_once, syscall.SIGINT, syscall.SIGTERM)
	<-serverSigChan_once
}

func ShutdownServer(err error) {
	log.Println(err)
	serverSigChan_once <- os.Interrupt
}
