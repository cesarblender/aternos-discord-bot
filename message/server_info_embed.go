package message

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	aternos "github.com/sleeyax/aternos-api"
)

func CreateServerInfoEmbed(info *aternos.ServerInfo) *discordgo.MessageEmbed {
	if info.DynIP == "" {
		info.DynIP = "no disponible"
	}

	return &discordgo.MessageEmbed{
		Title:       "Información del servidor",
		Description: fmt.Sprintf("El servidor '%s' está actualmente **%s**.", info.Name, info.StatusLabel),
		Color:       colorMap[info.Status],
		URL:         "https://aternos.org/server/",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Jugadores en línea",
				Value:  fmt.Sprintf("%d/%d", info.Players, info.MaxPlayers),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Dirección del servidor",
				Value:  fmt.Sprintf("`%s:%d`", info.Address, info.Port),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "IP Dinámica",
				Value:  fmt.Sprintf("`%s`", info.DynIP),
				Inline: true,
			},
		},
	}
}
