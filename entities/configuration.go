package entities

// Configuration Entity
type Configuration struct {
	WA WASession
}

// WASession is an entity that stores a whatsapp session
type WASession struct {
	ClientID    string
	ClientToken string
	ServerToken string
	EncKey      []byte
	MacKey      []byte
	Wid         string
}
