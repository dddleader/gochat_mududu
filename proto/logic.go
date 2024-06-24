package proto

// login
type LoginRequest struct {
	Name     string
	Password string
}
type LoginResponse struct {
	Code      int
	AuthToken string
}

// Register
type RegisterRequest struct {
	Name     string
	Password string
}

type RegisterReply struct {
	Code      int
	AuthToken string
}
