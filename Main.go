package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

const token string = "Njk1NTg2NTE5ODU1MTM2ODA4.XocWSA.qUyB1Z79PyA2C7ehioX7K1N3Vcs"
const color int = 0x00ff00

var start = time.Now()
var httpClient = &http.Client{}

var botID string

func main() {
	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := dg.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	botID = u.ID

	dg.AddHandler(messageHandler)
	err = dg.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")

	dg.UpdateStatusComplex(discordgo.UpdateStatusData{
		Game: &discordgo.Game{
			Name: "with tobio | haikyuhelp",
		},
	})

	<-make(chan struct{})
	return
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "haikyu") {
		s.ChannelMessageSend(m.ChannelID, "teste ;o")
		args := strings.Split(m.Content[6:], " ")
		switch args[0] {
		case "hug":
			s.ChannelMessageSend(m.ChannelID, "Hey! das geht noch nicht >.<")
			break
		case "help":
			s.ChannelMessageSend(m.ChannelID, "Es gibt `haikyuhug`..?")
			break
		}
	}
}
