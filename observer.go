package main

import "fmt"

type Observer interface {
	update(string)
}

type Observable interface {
	subscribe(observer Observer)
	notifyAll()
}

type User struct {
	id string
}

type Streamer struct {
	observerList []Observer
	name         string
}

func (u *User) update(streamerName string) {
	fmt.Printf("%s has started the stream\n", streamerName)
}

func newStreamer(name string) *Streamer {
	return &Streamer{name: name}
}

func (s *Streamer) subscribe(o Observer) {
	s.observerList = append(s.observerList, o)
}

func (s *Streamer) notifyAll() {
	for _, observer := range s.observerList {
		observer.update(s.name)
	}
}

func main() {
	Bratishkinoff := newStreamer("Bratishkinoff")

	blisstod := &User{id: "blisstod@gmail.com"}
	skits := &User{id: "skits@gmail.com"}

	Bratishkinoff.subscribe(blisstod)
	Bratishkinoff.subscribe(skits)

	Bratishkinoff.notifyAll()
}
