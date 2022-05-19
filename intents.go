package disgo

import (
	"errors"
)

// Reference: https://canary.discord.com/developers/docs/topics/gateway#gateway-intents
var IntentsFlags = map[string]uint64{
	"Guilds":                 1 << 0,
	"GuildMembers":           1 << 1,
	"GuildBans":              1 << 2,
	"GuildEmojisAndStickers": 1 << 3,
	"GuildIntegrations":      1 << 4,
	"GuildWebhooks":          1 << 5,
	"GuildInvites":           1 << 6,
	"GuildVoiceStates":       1 << 7,
	"GuildPresences":         1 << 8,
	"GuildMessages":          1 << 9,
	"GuildMessageReactions":  1 << 10,
	"GuildMessageTyping":     1 << 11,
	"DirectMessages":         1 << 12,
	"DirectMessageReactions": 1 << 13,
	"DirectMessageTyping":    1 << 14,
	"MessageContent":         1 << 15,
	"GuildScheduledEvents":   1 << 16,
}

func CalcIntents(intents ...string) (error, uint64) {
	var err error
	var intentsBitField uint64 = 0

	for _, intent := range intents {
		if IntentsFlags[intent] == 0 {
			err = errors.New("invalid intent: " + intent)
		} else {
			intentsBitField |= IntentsFlags[intent]
		}

	}

	return err, intentsBitField
}