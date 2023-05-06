package controller

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// Fish handles the /fish command and returns a random fish
func Fish(s *discordgo.Session, i *discordgo.InteractionCreate) {
	response := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Fish!",
		},
	}
	//TODO: Roll a random fish from the available fish, save that caught fish under the user that caught it.

	// send the response to Discord
	err := s.InteractionRespond(i.Interaction, response)
	if err != nil {
		log.Println("Error responding to /fish command: ", err)
	}
}
