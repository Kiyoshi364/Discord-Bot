package commands

import (
	"github.com/bwmarrin/discordgo"
	"strings"
	"fmt"
)

type Command interface {
	// Returns a basic description, what the command does
	Description() string

	// Returns a detailed description, what the command does, arguments...
	Help() string

	// All names accepted for calling this command
	Aliases() []string

	// Returns if it can be called by the input and where the parsing ended
	NameMatch(s string) (match bool, read int)

	// Handler to call when command is called
	Run(s *discordgo.Session, m *discordgo.MessageCreate, begin int) error
}

type BasicCommand struct {
}

func (c *BasicCommand) Description() string {
	return "Description"
}

func (c *BasicCommand) Help() string {
	return "Help"
}

func (c *BasicCommand) Aliases() []string {
	return make([]string, 0)
}

func (c *BasicCommand) NameMatch(s string) (match bool, read int) {
	match, read = c.MatchHelper(s, c.Aliases())
	return
}

func (c *BasicCommand) MatchHelper(s string, aliases []string) (match bool, read int) {
	match, read = false, 0

	for _,command := range aliases {
		if strings.HasPrefix(s, command) {
			match, read = true, len(command)
			return
		}
	}

	fmt.Println("DEGUB " + s + "\n" + aliases[0])

	return
}

func (c *BasicCommand) Run(s *discordgo.Session, m *discordgo.MessageCreate, begin int) error {
	return nil
}
