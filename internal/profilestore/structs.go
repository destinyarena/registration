package profilestore

// User is the database user struct
type User struct {
	Discord string
	Bungie  string
	Faceit  string
	Banned  bool
	IPHash  string
}
