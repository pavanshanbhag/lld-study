package main

import "mediator"

func main() {
	chatMediator := mediator.NewChatMediator()

	alice := mediator.NewUser("Alice")
	bob := mediator.NewUser("Bob")
	carol := mediator.NewUser("Carol")

	chatMediator.AddUser(alice)
	chatMediator.AddUser(bob)
	chatMediator.AddUser(carol)

	alice.SendMessage("Hello, everyone!")
	bob.SendMessage("Hi Alice!")
	carol.SendMessage("Hey folks!")
}
