package user

// User holds information about the user.
type User struct {
	id       string
	username string
	password string
}

// SetUsername will set the username.
func (u *User) SetUsername(username string) {
	u.username = username
}

// Username returns the username.
func (u *User) Username() string {
	return u.username
}

// SetPassword will set the password.
func (u *User) SetPassword(password string) {
	u.password = password
}

// Password returns the password.
func (u *User) Password() string {
	return u.password
}

// SetID will set the id.
func (u *User) SetID(id string) {
	u.id = id
}

// ID returns the id.
func (u *User) ID() string {
	return u.id
}
