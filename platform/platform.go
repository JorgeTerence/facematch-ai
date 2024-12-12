package platform

type Account struct {
	Id          interface{}
	PlatformId  string
	Username    string
	Connections []string
}

type TrueAccount struct {
	Id                string
	FirstName         string
	LastName          string
	Headline          string
	Location          string
	industry          string
	Summary           string
	ProfilePictureUrl string
}

const LOGIN_URL_PATTERN = `https://www.linkedin.com/oauth/v2/authorization?response_type=code&state=%d&scope=w_member_social&client_id=%s&redirect_uri=%s`

var (
	accountRepository = []Account{
		Account{Username: "Antonio", Id: ""},
	}
)

// In the future, contact my college's administration so they allow access to the LinkedIn API via their company page.
// For now, I'll be using a mock system
