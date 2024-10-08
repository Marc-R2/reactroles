package dgclient

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func (client *DiscordGoClient) GetOnInteractionHandler() func(*discordgo.Session, *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		log.Printf("[dgclient] Interaction received: %s\n", i.Interaction.Data)

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: ":gear: Working...",
				// Flags:   uint64(discordgo.MessageFlagsEphemeral),
				Flags: discordgo.MessageFlagsEphemeral,
			},
		})

		defer client.updateRoleSelectorMessage(i.GuildID)

		if i.Member.User.ID == s.State.User.ID {
			return
		}

		server := client.db.ServerConfigurationGet(i.GuildID)

		switch i.ApplicationCommandData().Options[0].Name {
		case Actions.Add:
			handleAddRoleSlashCommand(client, s, i, server)
		case Actions.Remove:
			handleRemoveRoleSlashCommand(client, s, i, server)
		case Actions.Update:
			handleUpdateRoleSlashCommand(client, s, i, server)
		case Actions.Help:
			handleHelpSlashCommand(client, s, i, server)
		case ActionConfigure:
			handleConfigureSlashCommand(client, s, i)
		case Actions.CreateChannel:
			handleCreateChannelSlashCommand(client, s, i, server)
		case Actions.LinkChannel:
			handleLinkChannelSlashCommand(client, s, i, server)
		case Actions.RemoveChannel:
			handleRemoveChannelSlashCommand(client, s, i, server)
		}
	}
}

func (client *DiscordGoClient) GetSlashCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "role",
		Description: "Manage roles",
		Options: []*discordgo.ApplicationCommandOption{
			addRoleSlashCommand(),
			removeRoleSlashCommand(),
			updateRoleSlashCommand(),
			helpSlashCommand(),
			configureServerSlashCommand(),
			createChannelSlashCommand(),
			linkChannelSlashCommand(),
			removeChannelSlashCommand(),
		},
	}
}
