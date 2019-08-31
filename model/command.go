package model

//Command is the base struct for bot commands
type Command struct {
	Name        string
	Description string
	Params      []string
	Action      Action
}

//NewCommand returns a pointer for new instance of Command
func NewCommand(name string) *Command {
	if name == "" {
		return nil
	}

	return &Command{
		Name: name,
	}
}
