package main

import "fmt"

type MockNotifier struct {
}

func (s MockNotifier) Notify() {
	fmt.Println("mock notifier")
}

func TestNotifier() {
	notify(MockNotifier{})
}
