package discordbot

import (
	"github.com/bwmarrin/discordgo"
)

/* link to add o server
https://discord.com/oauth2/authorize?client_id=452640305141907457&scope=bot&permissions=0 */

func Init() (err error) {

	bot, er := newBot("config.json")
	if er != nil {
		err = er
		return
	}

	s, er := discordgo.New("Bot " + bot.Token)
	if er != nil {
		err = er
		return
	}

	s.AddHandler(bot.MessageHandler)

	er = s.Open()
	if er != nil {
		err = er
		return
	}
	return
}
