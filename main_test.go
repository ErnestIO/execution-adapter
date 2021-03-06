package main

import (
	"errors"
	"os"
	"testing"
	"time"

	ecc "github.com/ernestio/ernest-config-client"
	"github.com/nats-io/nats"

	. "github.com/smartystreets/goconvey/convey"
)

func wait(ch chan bool) error {
	return waitTime(ch, 500*time.Millisecond)
}

func waitTime(ch chan bool, timeout time.Duration) error {
	select {
	case <-ch:
		return nil
	case <-time.After(timeout):
	}
	return errors.New("timeout")
}

func TestBasicRedirections(t *testing.T) {

	Convey("Given this service is fully set up", t, func() {
		n := ecc.NewConfig(os.Getenv("NATS_URI")).Nats()
		chfak := make(chan bool)
		cherr := make(chan bool)
		chvcl := make(chan bool)

		n.Subscribe("config.get.connectors", func(msg *nats.Msg) {
			n.Publish(msg.Reply, []byte(`{"executions":["fake","salt"]}`))
		})

		go main()

		n.Subscribe("execution.create.fake", func(msg *nats.Msg) {
			chfak <- true
		})
		n.Subscribe("execution.create.error", func(msg *nats.Msg) {
			cherr <- true
		})
		n.Subscribe("execution.create.salt", func(msg *nats.Msg) {
			chvcl <- true
		})
		Convey("When it receives an invalid fake message", func() {
			n.Publish("execution.create", []byte(`{"service":"aaa","type":"a"}`))
			Convey("Then it should redirect to execution error creation", func() {
				So(wait(cherr), ShouldNotBeNil)
			})
		})
		Convey("When it receives a valid fake message", func() {
			n.Publish("execution.create", []byte(`{"service":"aaa","type":"fake"}`))
			Convey("Then it should redirect it to a fake connector", func() {
				So(wait(chfak), ShouldBeNil)
			})
		})
		Convey("When it receives a valid salt message", func() {
			n.Publish("execution.create", []byte(`{"service":"aaa","type":"salt"}`))
			Convey("Then it should redirect it to a fake connector", func() {
				So(wait(chvcl), ShouldBeNil)
			})
		})
	})
}
