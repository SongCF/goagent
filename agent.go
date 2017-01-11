package main

import (
	"time"
	"fmt"
)


func agent(sess *Session) {
	// minute timer
	min_timer := time.After(time.Minute)

	for {
		select {
		case msg := <- sess.in:
			sess.packet_count++
			// TODO game rpc
			// ret
			ret := "ack_" + string(msg[:])
			send(sess, []byte(ret))
		case <- sess.ntf:
			// TODO
			push := "ntf"
			send(sess, []byte(push))
		case <- min_timer:
			timer_work()
			min_timer = time.After(time.Minute)
		case <- die:
			close(sess.die)
			return
		}
	}
}


func timer_work() {
	fmt.Println("on timer")
}