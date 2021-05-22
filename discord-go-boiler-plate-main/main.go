package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	BotID string
	NameofBot *discordgo.Session
	// Public variables
	BotToken  string
	BotPrefix string
	// Private variables
	config *configStruct
)

func main() {

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())

	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())

	}

	BotToken = config.BotToken
	BotPrefix = config.BotPrefix

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	NameofBot, err := discordgo.New("bot " + config.BotToken)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := NameofBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	NameofBot.AddHandler(messageHandler)

	err = NameofBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("bot is online")

	<-make(chan struct{})
}

type configStruct struct {
	BotToken  string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func messageHandler(x *discordgo.Session, y *discordgo.MessageCreate) {

	if strings.HasPrefix(y.Content, config.BotPrefix) {
		if y.Author.ID == BotID {
			return
		}

		if m.Content == "!d" {
			_, _ = x.ChannelMessageSend(y.ChannelID, "https://tenor.com/view/dababy-rapper-hip-hop-rap-digibyte-gif-17582117") // example command for testing purposes
		}
	}
}
