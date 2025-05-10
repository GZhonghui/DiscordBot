package tasks

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"local.dev/DiscordBot/tools"
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
	reply := fmt.Sprintf("<@%s> 又来找姐姐聊天了吗？需要先配置一下OpenAI的密钥呢！", m.Author.ID)

	if tools.OpenaiKey != "" {
		prompt := "现在请假装你是G的姐姐，你的性格温柔，宠爱自己的弟弟，G现在需要你的鼓励，请你夸夸他，让他感到姐姐的温暖吧！"
		prompt += "G对你说的话是：「" + m.Content + "」"
		prompt += "（请注意输出格式：直接输出纯文本，不要用格式控制符号）"
		result, err := tools.CallChatGPT(tools.OpenaiKey, "gpt-4.1", prompt)
		if err != nil {
			reply = fmt.Sprintf("<@%s> 调用失败了哦\n%s", m.Author.ID, err)
		} else {
			reply = fmt.Sprintf("<@%s> %s", m.Author.ID, result)
		}
	}

	s.ChannelMessageSend(m.ChannelID, reply)
}
