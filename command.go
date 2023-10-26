package main

import "fmt"

// command interface
type Command interface {
	execute()
}

// concrete commands
type RunCommand struct {
	character Character
}

func (c *RunCommand) execute() {
	c.character.run()
}

type StayCommand struct {
	character Character
}

func (c *StayCommand) execute() {
	c.character.stay()
}

type SitCommand struct {
	character Character
}

func (c *SitCommand) execute() {
	c.character.sit()
}

//------------------------

// reciever interface
type Character interface {
	run()
	stay()
	sit()
}

// concrete reciever
type TankCharacter struct {
}

func (c *TankCharacter) run() {
	fmt.Println("The Tank Character is running!")
}

func (c *TankCharacter) stay() {
	fmt.Println("The Tank Character is just staying!")
}

func (c *TankCharacter) sit() {
	fmt.Println("The Tank Character is just sitting!")
}

// invoker
type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

func main() {
	tankCharacter := &TankCharacter{}
	runCommand := &RunCommand{
		character: tankCharacter,
	}
	runButton := &Button{
		command: runCommand,
	}
	runButton.press()

	stayCommand := &StayCommand{
		character: tankCharacter,
	}
	stayButton := &Button{
		command: stayCommand,
	}
	stayButton.press()
}
