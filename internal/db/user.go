package db

// User mirrors the users row shape. Email added in migration 003.
type User struct {
	ID    int
	Name  string
	Email string
}

func (u *User) HasEmail() bool {
	return u.Email != ""
}
