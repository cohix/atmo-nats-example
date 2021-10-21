package main

import (
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/suborbital/grav/grav"
	"github.com/suborbital/grav/transport/nats"
	"github.com/suborbital/vektor/vlog"
)

func main() {
	logger := vlog.Default(vlog.Level(vlog.LogLevelDebug))

	gnats, err := nats.New("nats://localhost:4222")
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to nats.New"))
	}

	g := grav.New(
		grav.UseLogger(logger),
		grav.UseTransport(gnats),
	)

	if err := g.ConnectBridgeTopic("com.suborbital.atmo/hello"); err != nil {
		log.Fatal(errors.Wrap(err, "failed to ConnectTopic"))
	}

	if err := g.ConnectBridgeTopic("com.suborbital.atmo/respond"); err != nil {
		log.Fatal(errors.Wrap(err, "failed to ConnectTopic"))
	}

	pod := g.Connect()
	pod.On(func(msg grav.Message) error {
		fmt.Println("received something:", string(msg.Data()))
		return nil
	})

	go func() {
		<-time.After(time.Second)
		pod.Send(grav.NewMsg("com.suborbital.atmo/hello", []byte("world")))

		<-time.After(time.Second)
		pod.Send(grav.NewMsg("com.suborbital.atmo/hello", []byte("again")))
	}()

	<-time.After(time.Second * 5)
}
