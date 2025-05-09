package tasks

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const (
	ChannelLoveID = "1370348253312192582" // #love
)

// onMessageCreate handles incoming messages
func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Check if the bot is mentioned
	for _, user := range m.Mentions {
		if user.ID == s.State.User.ID {
			// Bot was mentioned!
			if ChannelLoveID == m.ChannelID {
				onMentionInLove(s, m)
			} else {
				reply := fmt.Sprintf("主人 <@%s>！我在的哦( ᵒ̴̶̷̤ꈊ˂̶̤̀ )✧", m.Author.ID)
				s.ChannelMessageSend(m.ChannelID, reply)
			}
			break
		}
	}
}

func onMentionInLove(s *discordgo.Session, m *discordgo.MessageCreate) {
	reply := fmt.Sprintf("<@%s> 又来找姐姐聊天了吗？", m.Author.ID)
	s.ChannelMessageSend(m.ChannelID, reply)
}
