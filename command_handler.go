package aternos_discord_bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sleeyax/aternos-discord-bot/database/models"
	"github.com/sleeyax/aternos-discord-bot/message"
)

// handleCommands responds to incoming interactive commands on discord.
func (ab *Bot) handleCommands(s *discordgo.Session, i *discordgo.InteractionCreate) {
	command := i.ApplicationCommandData()

	// wrap functions around our utilities to make life easier
	sendText := func(content string) {
		respondWithText(s, i, content)
	}
	sendHiddenText := func(content string) {
		respondWithHiddenText(s, i, content)
	}
	sendErrorText := func(content string, err error) {
		respondWithError(s, i, content, err)
	}

	switch command.Name {
	case HelpCommand:
		sendHiddenText(message.FormatDefault(faq))
	case PingCommand:
		sendHiddenText(message.FormatDefault("Pong!"))
	case ConfigureCommand:
		options := optionsToMap(command.Options)

		err := ab.Database.UpdateServerSettings(&models.ServerSettings{
			GuildID:       i.GuildID,
			SessionCookie: options[SessionOption].StringValue(),
			ServerCookie:  options[ServerOption].StringValue(),
		})
		if err != nil {
			sendErrorText("Failed to save configuration.", err)
			break
		}

		sendText(message.FormatSuccess("Configuration changed successfully."))
	case StatusCommand:
		fallthrough
	case InfoCommand:
		fallthrough
	case PlayersCommand:
		fallthrough
	case StartCommand:
		fallthrough
	default:
		sendText(message.FormatWarning("Command unavailable. Please try again later or refresh your discord client `CTRL + R`"))
	}
}
