package bridge

import "fmt"

type OS interface {
	Install()
}

type Android struct{}

func (a Android) Install() {
	fmt.Println("Install OS Android")
}

type IOS struct{}

func (i IOS) Install() {
	fmt.Println("Install OS IOS")
}
