package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

const (
	Hello_Description = `command: hello
	Answers with "Hallo"`
	Hello_Default_Alias = "hello"
)

type HelloCommand struct {
	BasicCommand
}

func (c *HelloCommand) Description() string {
	return Hello_Description
}

func (c *HelloCommand) Aliases() []string {
	return []string{Hello_Default_Alias}
}

func (c *HelloCommand) NameMatch(s string) (match bool, read int) {
	return c.MatchHelper(s, c.Aliases())
}

func (c *HelloCommand) Run(s *discordgo.Session, m *discordgo.MessageCreate, begin int) error {

	left := m.Content[begin:]

	id := m.Author.ID
	msg := "Hallo, <@" + id + ">" + left

	allowMent := &discordgo.MessageAllowedMentions{
		Users: []string{id},
	}

	data := &discordgo.MessageSend{
		Content: msg,
		AllowedMentions: allowMent,
	}
	_, err := s.ChannelMessageSendComplex(m.ChannelID, data)

	if err != nil {
		return fmt.Errorf("Hello Run: %v", err.Error())
	}
	return nil
}
