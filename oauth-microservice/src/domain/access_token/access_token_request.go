package access_token

import "github.com/menxqk/rest-microservices-in-go/common/errors"

const (
	GRANT_TYPE_PASSWORD           = "password"
	GRANT_TYPE_CLIENT_CREDENTIALS = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	// Used for client_credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (r *AccessTokenRequest) Validate() errors.RestError {
	switch r.GrantType {
	case GRANT_TYPE_PASSWORD:
		break
	case GRANT_TYPE_CLIENT_CREDENTIALS:
		break
	default:
		return errors.NewBadRequestError("invalid grant_type parameter")
	}

	// TODO: Validate parameter for each grant_type
	return nil
}
