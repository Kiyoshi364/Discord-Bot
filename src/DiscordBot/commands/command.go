package commands

import (
	"github.com/bwmarrin/discordgo"
)

type Command interface {
	// Returns a basic description, what the command does
	Description() string

	// Returns a detailed description, what the command does, arguments
	Help() string

	// All names accepted for calling this command
	Aliases() []string

	// Handler to call when command is called
	Run(s *discordgo.Session, m *discordgo.MessageCreate, begin int) error
}

type BasicCommand struct {
	Alias	[]string
}

func (c *BasicCommand) Description() string {
	return "Description"
}

func (c *BasicCommand) Help() string {
	return "Help"
}

func (c *BasicCommand) Aliases() []string {
	return c.Alias
}

func (c *BasicCommand) Run(s *discordgo.Session, m *discordgo.MessageCreate, begin int) error {
	return nil
}
