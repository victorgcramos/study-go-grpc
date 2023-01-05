package user

const (
	RoutePrefixUser = "/user"

	RouteVersion = "/"
	RouteUserNew = "/new"
)

// NewUser defines the request body for user/new api handler
type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewUserReply defines the response body for user/new api handler
type NewUserReply struct {
	Token string `json:"token"` // Server verification token
}

type VersionReply struct {
	Version string `json:"version"`
	Token   string `json:"token"`
}
