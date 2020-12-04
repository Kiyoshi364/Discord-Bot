package discordbot

import (
	"github.com/bwmarrin/discordgo"
	"./commands"
	"strings"
	"fmt"
)

const token string = "NDUyNjQwMzA1MTQxOTA3NDU3.WxNAIg._tChkfwo80J9UM_w8Vm8UB8XsoY"
const prefix string = ":>"

/* link https://discord.com/oauth2/authorize?client_id=452640305141907457&scope=bot&permissions=0 */

func Init() (err error) {
	s, er := discordgo.New("Bot " + token)
	if er != nil {
		err = er
		return
	}

	s.AddHandler(messageHandler)

	er = s.Open()
	if er != nil {
		err = er
		return
	}
	return
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, prefix) {
		return
	}

	hello := commands.HelloCommand{}

	if strings.HasPrefix(m.Content, prefix + hello.Aliases()[0]) {
		err := hello.Run(s, m, len(prefix))
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	wakeup := commands.WakeupCommand{}

	if strings.HasPrefix(m.Content, prefix + wakeup.Aliases()[0]) {
		err := wakeup.Run(s, m, len(prefix))
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
