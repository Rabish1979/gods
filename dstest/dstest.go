package main

import (
	"fmt"
    "github.com/Rabish1979/gods/linklist"
)

func main() {
	myPlaylist := createLinkList()
	myPlaylist.insertAtFront(20)
	myPlaylist.insertAtFront(30)
	myPlaylist.insertAtFront(40)
	myPlaylist.Print()
	fmt.Println("hello package")
}
