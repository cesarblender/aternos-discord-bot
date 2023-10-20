package aternos_discord_bot

import "github.com/bwmarrin/discordgo"

const (
	HelpCommand      = "help"
	PingCommand      = "pinga"
	ConfigureCommand = "configure"
	StartCommand     = "start"
	StopCommand      = "stop"
	StatusCommand    = "status"
	InfoCommand      = "info"
	PlayersCommand   = "players"
	SessionOption    = "session"
	ServerOption     = "server"
)

var (
	adminPermissions int64 = discordgo.PermissionManageServer
	userPermissions  int64 = discordgo.PermissionUseSlashCommands
	dmPermission           = false
)

// List of available discord commands.
var commands = []*discordgo.ApplicationCommand{
	{
		Name:        ConfigureCommand,
		Description: "Guardar la configuración",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:         SessionOption,
				Description:  "Establecer el valor de la cookie ATERNOS_SESSION",
				Type:         discordgo.ApplicationCommandOptionString,
				Required:     true,
				ChannelTypes: []discordgo.ChannelType{discordgo.ChannelTypeGuildText},
			},
			{
				Name:         ServerOption,
				Description:  "Establecer el valor de la cookie ATERNOS_SERVER",
				Type:         discordgo.ApplicationCommandOptionString,
				Required:     true,
				ChannelTypes: []discordgo.ChannelType{discordgo.ChannelTypeGuildText},
			},
		},
		DefaultMemberPermissions: &adminPermissions,
		DMPermission:             &dmPermission,
	},
	{
		Name:                     StartCommand,
		Description:              "Iniciar el servidor de Minecraft",
		DefaultMemberPermissions: &adminPermissions,
		DMPermission:             &dmPermission,
	},
	{
		Name:                     StopCommand,
		Description:              "Detener el servidor de Minecraft",
		DefaultMemberPermissions: &adminPermissions,
		DMPermission:             &dmPermission,
	},
	{
		Name:                     PingCommand,
		Description:              "Comprobar si el bot de Discord está activo",
		DefaultMemberPermissions: &userPermissions,
		DMPermission:             &dmPermission,
	},
	{
		Name:                     StatusCommand,
		Description:              "Obtener el estado del servidor de Minecraft",
		DefaultMemberPermissions: &userPermissions,
		DMPermission:             &dmPermission,
	},
	{
		Name:                     InfoCommand,
		Description:              "Obtener información detallada sobre el estado del servidor de Minecraft",
		DefaultMemberPermissions: &userPermissions,
		DMPermission:             &dmPermission,
	},
	{
		Name:                     PlayersCommand,
		Description:              "Listar jugadores activos",
		DefaultMemberPermissions: &userPermissions,
		DMPermission:             &dmPermission,
	},
	{
		Name:                     HelpCommand,
		Description:              "Obtener ayuda",
		DefaultMemberPermissions: &adminPermissions,
		DMPermission:             &dmPermission,
	},
}
