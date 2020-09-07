package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go CheckSignal(cancelFunc)

	if err := ServerStart(ctx); err != nil {
		log.Fatal(err)
	}

}

func CheckSignal(cancelFunc context.CancelFunc) {
	signalChannel := make(chan os.Signal) // Ctrl + C
	signal.Notify(signalChannel, os.Interrupt)
	for {
		currentSignal := <-signalChannel
		switch currentSignal {
		case os.Interrupt:
			//В случае если пришел сигнал Ctrl+C - вызываем функцию отмены контекста
			cancelFunc() //-> ctx.Done()
			return
		}
	}
}

func ServerStart(ctx context.Context) error {
	loader, err := net.ResolveTCPAddr("tcp", ":8000")
	if err != nil {
		return err
	}

	listener, err := net.ListenTCP("tcp", loader)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		select {
		case <-ctx.Done():
			log.Println("ContextFunc() done!")
			log.Println("Connection closed!")
			return nil
		default:
			if err := listener.SetDeadline(time.Now().Add(time.Second * 2)); err != nil {
				return err
			}
			_, err := listener.Accept()
			if err != nil {
				if os.IsTimeout(err) {
					continue
				}
				return err
			}
		}
	}

}
