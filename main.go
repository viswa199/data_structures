package main

import (
	linkedlistgo "DS/linkedList"
)

func main() {
	ll := linkedlistgo.Single{}

	//inserting elements
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Insert(4)

	//display elements
	ll.Display()

	//delete elements
	ll.Delete(2)

	//displaying after deleting elements
	ll.Display()
}
