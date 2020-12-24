package entities

// Configuration Entity
type Configuration struct {
	AppConfig AppConfiguration
	WA        WASession
}

// AppConfiguration is an entity that stores the app main configuration
type AppConfiguration struct {
	Host string
	Port string
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
