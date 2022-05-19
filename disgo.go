package disgo

import "runtime"

// The main hub for interacting with the Discord API, and the starting point for any bot
func New(botToken string) (s *Session) {
	s = &Session{
		Identify: Identify{
			Token:          botToken,
			Intents:        0,
			ShardID:        0,
			ShardCount:     1,
			Compress:       false,
			LargeThreshold: 50,
			Properties: Properties{
				Browser: "disgo",
				Device:  "disgo",
				Os:      runtime.GOOS,
			},
		},
	}

	return s
}
