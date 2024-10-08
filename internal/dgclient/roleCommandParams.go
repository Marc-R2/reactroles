package dgclient

import (
	"github.com/bwmarrin/discordgo"
	"github.com/zaptross/reactroles/internal/pgdb"
)

type RoleCommandParams struct {
	Server      *pgdb.ServerConfiguration
	Session     *discordgo.Session
	Message     *discordgo.MessageCreate
	Interaction *discordgo.InteractionCreate
	Client      *DiscordGoClient
	Action      string
	Rest        []string
}

func (rc *RoleCommandParams) GuildID() string {
	if rc.Message != nil {
		return rc.Message.GuildID
	}

	return rc.Interaction.GuildID
}

func (rc *RoleCommandParams) ChannelID() string {
	if rc.Message != nil {
		return rc.Message.ChannelID
	}

	return rc.Interaction.ChannelID
}

func (rc *RoleCommandParams) AuthorID() string {
	if rc.Message != nil {
		return rc.Message.Author.ID
	}

	return rc.Interaction.Member.User.ID
}

func (rcp *RoleCommandParams) Reply(message string) {
	if rcp.Interaction != nil {
		/* rcp.Session.InteractionResponseEdit(rcp.Interaction.Interaction, &discordgo.WebhookEdit{
			Content: message,
		}) */
		iRespEdit(message, rcp.Session, rcp.Interaction)
	} else {
		rcp.Session.ChannelMessageSend(rcp.ChannelID(), message)
	}
}
