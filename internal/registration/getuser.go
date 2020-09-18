package registration

import (
	"crypto/sha1"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type (
	discordClaims struct {
		ID            string `json:"id"`
		Username      string `json:"username"`
		Discriminator string `json:"discriminator"`
		jwt.StandardClaims
	}

	bungieClaims struct {
		ID                  string `json:"membershipId" validate:"required"`
		DisplayName         string `json:"displayName" validate:"required"`
		SteamDisplayName    string `json:"steamDisplayName,omitempty"`
		XboxDisplayName     string `json:"xboxDisplayName,omitempty"`
		PSNDisplayName      string `json:"psnDisplayName,omitempty"`
		BlizzardDisplayName string `json:"blizzardDisplayName,omitempty"`
		jwt.StandardClaims
	}

	faceitClaims struct {
		GUID     string `json:"guid"`
		Nickname string `json:"nickname"`
		jwt.StandardClaims
	}
)

func (h *handler) getUser(p *Payload, ip string) (*user, error) {
	dclaims := new(discordClaims)
	if err := h.JWTManager.Decrypt(p.Discord, dclaims); err != nil {
		return nil, err
	}

	bclaims := new(bungieClaims)
	if err := h.JWTManager.Decrypt(p.Bungie, bclaims); err != nil {
		return nil, err
	}

	fclaims := new(faceitClaims)
	if err := h.JWTManager.Decrypt(p.Faceit, fclaims); err != nil {
		return nil, err
	}

	iphash := ""

	hash := sha1.New()
	hash.Write([]byte(ip))
	bs := hash.Sum(nil)
	iphash = fmt.Sprintf("%x", bs)

	u := &user{
		Discord: dclaims.ID,
		Bungie:  bclaims.ID,
		Faceit:  fclaims.GUID,
		IPHash:  iphash,
	}

	return u, nil
}
