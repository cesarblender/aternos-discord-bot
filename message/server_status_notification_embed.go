package message

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	aternos "github.com/sleeyax/aternos-api"
)

func CreateServerStatusNotificationEmbed(info *aternos.ServerInfo) (*discordgo.MessageEmbed, error) {
	switch info.Status {
	case aternos.Online:
		return &discordgo.MessageEmbed{
			Title:       "El servidor está en línea :v",
			Description: fmt.Sprintf("papu dale rapido :''v el servidor se va a cerrar en %d segundos pendejo, te quiero mucho de verdad eres lo mejor que existe dios te bendiga te quiero te amo te aprecio mucho papu, de verdad, te quiero.", info.Countdown),
			Color:       colorMap[aternos.Online],
			Fields: []*discordgo.MessageEmbedField{
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
		}, nil
	case aternos.Offline:
		return &discordgo.MessageEmbed{
			Title:       "El servidor está fuera de línea",
			Description: "El servidor está actualmente fuera de línea.",
			Color:       colorMap[aternos.Offline],
		}, nil
	default:
		return nil, fmt.Errorf("código de estado del servidor desconocido '%d'", info.Status)
	}
}
