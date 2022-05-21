package disgo

// Represents a Discord Unavailable Guild object
type APIUnavailableGuild struct {
	// The ID of the guild
	Id string `json:"id"`
	// Whether the guild is unavailable
	Unavailable bool `json:"unavailable"`
}

// Represents a Discord User object
type APIUser struct {
	// The user's ID
	Id string `json:"id"`
	// The user's username
	Username string `json:"username"`
	// The user's discriminator
	Discriminator string `json:"discriminator"`
	// The user's avatar hash
	Avatar string `json:"avatar"`
	// Whether the user is a bot
	Bot bool `json:"bot"`
	// whether the user is an Official Discord System user (part of the urgent message system)
	System bool `json:"system"`
	// Whether the user has two factor enabled on their account
	MfaEnabled bool `json:"mfa_enabled"`
	// The user's banner hash
	Banner string `json:"banner"`
	// The user's banner color encoded as an integer representation of hexadecimal color code
	AccentColor int `json:"accent_color"`
	// Whether the email on this account has been verified
	Verified bool `json:"verified"`
	// The user's email
	Email string `json:"email"`
	// The public flags on the user's account
	PublicFlags int `json:"public_flags"`
}
