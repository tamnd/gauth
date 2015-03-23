package gauth

// User contains information about a user
type User struct {
	ID          string
	Name        string
	Email       string
	Username    string
	Firstname   string
	Lastname    string
	Location    string
	Description string
	Avatar      string
	Phone       string
	Urls        map[string]string
	Raw         interface{}
}
