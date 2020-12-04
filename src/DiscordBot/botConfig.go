package discordbot

import (
	"github.com/bwmarrin/discordgo"
	"./commands"
	"strings"
	"fmt"
)

const (
	Num_Commands	= 2
)

type BotConfig struct {
	Token		string `json:"Token"`
	Prefix		string `json:"Prefix"`
	Commands	[]commands.Command
}

func newBot(filename string) (bot *BotConfig, err error) {
	bot = &BotConfig{}

	er := ReadConfig(filename, bot)
	if er != nil {
		err = er
		return
	}

	addCommands(bot)

	return
}

func addCommands(bot *BotConfig) {
	bot.Commands = make([]commands.Command, Num_Commands, Num_Commands)

	bot.Commands[0] = &commands.HelloCommand{}
	bot.Commands[1] = &commands.WakeupCommand{}
}

func (bot *BotConfig) MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, bot.Prefix) {
		return
	}

	found := false
	prefixLen := len(bot.Prefix)
	for _, command := range bot.Commands {
		match, read := command.NameMatch(m.Content[prefixLen:])
		if match {
			err := command.Run(s, m, prefixLen + read)
			if err != nil {
				fmt.Println(err.Error())
			}
			found = true
			break
		}
	}

	if !found {
		id := m.Author.ID
		msg := "<@" + id + "> Command not found: " +
			m.Content[len(bot.Prefix):]

		allowMent := &discordgo.MessageAllowedMentions{
			Users: []string{id},
		}

		data := &discordgo.MessageSend{
			Content: msg,
			AllowedMentions: allowMent,
		}

		_, err := s.ChannelMessageSendComplex(m.ChannelID, data)
		if err != nil {
			fmt.Printf("MessageHandler Command Not Found: %s", err.Error())
		}

	}

	return
}
