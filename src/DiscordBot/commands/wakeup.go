package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

const (
	Wakeup_Description = `command: wakeup <user> <channel 1> <channel 2>
	Repete 3 vezes
		Move <user> para <channel 1>
		Move <user> para <channel 2>
	Move <user> para o canal que estava antes`
	Wakeup_Help = Wakeup_Description
	Wakeup_Default_Alias = "wakeup"
)

type WakeupCommand struct {
	BasicCommand
}


func (c *WakeupCommand) Description() string {
	return Wakeup_Description
}

func (c *WakeupCommand) Help() string {
	return Wakeup_Help
}

func (c *WakeupCommand) Aliases() []string {
	return []string{Wakeup_Default_Alias}
}

func (c *WakeupCommand) NameMatch(s string) (match bool, read int) {
	match, read = c.MatchHelper(s, c.Aliases())
	return
}

// TODO: find a way to get user's initial channel and return him there
func (c *WakeupCommand) Run(s *discordgo.Session, m *discordgo.MessageCreate, begin int) (err error) {

	// Parse info
	in := m.Content[begin:]

	words := strings.Split(in, " ")
	words = words[1:]
	flags := byte(0)
	var userID, chan1, chan2, chan0 string

	for _, word := range(words) {
		if flags == 0 {
			if ok := validateUserID(word); !ok {
				continue
			}

			word = word[3:len(word)-1]
			userID = word
			flags += 1
		} else if flags > 3 {
			break
		} else {
			if ok := validateChannelID(word); !ok{
				continue
			}
			word = word[2:len(word)-1]

			ch, err := s.Channel(word)
			if err != nil || ch != nil && ch.Type != discordgo.ChannelTypeGuildVoice {
				msg := fmt.Sprintf("%s is not a valid voice channel in this server", word)
				_, er := s.ChannelMessageSend(m.ChannelID, msg)
				if er != nil {
					return fmt.Errorf("Wakeup Run CMSend: %s :> %v", msg, er.Error())
				}
				if err != nil {
					return fmt.Errorf("Wakeup Run Channel: %s :> %v", word, err.Error())
				}

			}

			if flags == 1 {
				chan1 = word
			} else if flags == 2 {
				chan2 = word
			} else {
				chan0 = word
			}
			flags += 1
		}
	}
	if userID == "" {
		msg := "Couldn't find a User Mention"
		_, err := s.ChannelMessageSend(m.ChannelID, msg)
		if err != nil {
			return fmt.Errorf("Wakeup Run CMSend: %s :> %v", msg,  err.Error())
		}
		return nil
	} else if chan1 == "" {
		msg := "Couldn't find a Channel Mention after the User Mention"
		_, err := s.ChannelMessageSend(m.ChannelID, msg)
		if err != nil {
			return fmt.Errorf("Wakeup Run CMSend: %s :> %v", msg,  err.Error())
		}
		return nil
	} else if chan2 == "" {
		msg := "Couldn't find a second Channel Mention after the User Mention"
		_, err := s.ChannelMessageSend(m.ChannelID, msg)
		if err != nil {
			return fmt.Errorf("Wakeup Run CMSend: %s :> %v", msg,  err.Error())
		}
		return nil
	} else if chan0 == "" {
		chan0 = chan1
	}

	user, err := s.User(userID)
	username := ""
	if err != nil {
		return fmt.Errorf("Wakeup Run User: %s :> %v", userID, err.Error())
	} else {
		username = user.Username + " "
	}

	msg := "Waking " + username + "up Process Iniciated"
	_, err = s.ChannelMessageSend(m.ChannelID, msg)
	if err != nil {
		return fmt.Errorf("Wakeup Run CMSend: %s :> %v", msg,  err.Error())
	}

	// Wakeup
	for i := 0; i < 3; i+=1 {
		err = s.GuildMemberMove(m.GuildID, userID, &chan1)
		if err != nil {
			return fmt.Errorf("Wakeup Run GMMove: %s, %s, %v :> %v", m.GuildID, userID, &chan1, err.Error())
		}

		for j := int32(1); j != 0; j+=1 {}

		err = s.GuildMemberMove(m.GuildID, userID, &chan2)
		if err != nil {
			return fmt.Errorf("Wakeup Run GMMove: %s, %s, %v :> %v", m.GuildID, userID, &chan2, err.Error())
		}

		for j := int32(1); j != 0; j+=1 {}
	}
	err = s.GuildMemberMove(m.GuildID, userID, &chan0)
	if err != nil {
		return fmt.Errorf("Wakeup Run GMMove: %s, %s, %v :> %v", m.GuildID, userID, &chan0, err.Error())
	}

	msg = "Waking " + username + "up Process Finished"
	_, err = s.ChannelMessageSend(m.ChannelID, msg)
	if err != nil {
		return fmt.Errorf("Wakeup Run CMSend: %s :> %v", msg,  err.Error())
	}
	return nil
}

func validateUserID(str string) bool {
	if len(str) < 4 {
		return false
	}
	for i, c := range(str) {
		switch i {
		case 0:
			if c != '<' {
				return false
			}
		case 1:
			if c != '@' {
				return false
			}
		case len(str)-1:
			if c != '>' {
				return false
			}
		case 2:
			if c != '!' {
				return false
			}
		default:
			if c < '0' || c > '9' {
				return false
			}
		}
	}
	return true
}

func validateChannelID(str string) bool {
	if len(str) < 3 {
		return false
	}
	for i, c := range(str) {
		switch i {
		case 0:
			if c != '<' {
				return false
			}
		case 1:
			if c != '#' {
				return false
			}
		case len(str)-1:
			if c != '>' {
				return false
			}
		default:
			if c < '0' || c > '9' {
				return false
			}
		}
	}
	return true
}
