package gravatar

import (
	"fmt"

	"github.com/Vaelatern/monitoror/pkg/hash"
)

const GravatarURL = "https://www.gravatar.com/avatar/%s?d=blank"

func GetGravatarURL(email string) string {
	return fmt.Sprintf(GravatarURL, hash.GetMD5Hash(email))
}
