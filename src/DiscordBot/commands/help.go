package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

const (
	Help_Description = `command: help [command]
	Prints [command]'s help message.
	If can't find the command, prints all available commands.`
	Help_Default_Alias = "help"
	Help_Default_Alias2 = "h"
)

type HelpCommand struct {
	BasicCommand
	Commands	[]Command
}


func (c *HelpCommand) Description() string {
	return Help_Description
}

func (c *HelpCommand) Aliases() []string {
	return []string{Help_Default_Alias, Help_Default_Alias2}
}

func (c *HelpCommand) NameMatch(s string) (match bool, read int) {
	return c.MatchHelper(s, c.Aliases())
}

func (c *HelpCommand) Run(s *discordgo.Session, m *discordgo.MessageCreate, begin int) error {

	if len(m.Content) > begin {
		begin += 1
	}

	input := m.Content[begin:]

	msg := "Those commands exist:\n"
	for _, command := range c.Commands {

		match, _ := command.NameMatch(input)
		if match {
			msg = command.Description()
			break
		}

		msg += "\t\\> "
		for i, alias := range command.Aliases() {
			if i > 0 {
				msg += " | "
			}
			msg += alias
		}
		msg += "\n"
	}

	allowMent := &discordgo.MessageAllowedMentions{}
	data := &discordgo.MessageSend{
		Content: msg,
		AllowedMentions: allowMent,
	}

	_, err := s.ChannelMessageSendComplex(m.ChannelID, data)
	if err != nil {
		return fmt.Errorf("Help Run: %v", err.Error())
	}

	return nil
}
