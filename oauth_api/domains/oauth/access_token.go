package oauth

import "time"

// AccessToken for a user
type AccessToken struct {
	UserID      int64  `json:"user_id"`
	AccessToken string `json:"access_token"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).UTC().Before(time.Now().UTC())
}
