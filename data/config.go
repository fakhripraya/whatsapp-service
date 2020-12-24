package data

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fakhripraya/whatsapp-service/entities"

	whatsapp "github.com/Rhymen/go-whatsapp"
	"github.com/hashicorp/go-hclog"
	qrcode "github.com/skip2/go-qrcode"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// Whatsapp is a struct for whatsapp variable
type Whatsapp struct {
	Wac    *whatsapp.Conn
	logger hclog.Logger
}

// config is a variable that holds the application new session in yaml form
var config []byte

// WAconfig is an application current whatsapp configuration from config.whatsapp.yaml file
var WAconfig *entities.WASession

// CurrWASession is an entity that holds application current whatsapp session
var CurrWASession whatsapp.Session

// NewWA creates a new Whatsapp handler
func NewWA(logger hclog.Logger) (*Whatsapp, error) {

	// setting the working directory path
	logger.Info("Looking for the WhatsApp API configuration working directory")
	workingDirectiory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// set the config name to config.whatsapp and add the config path
	viper.SetConfigName("config." + "whatsapp")
	viper.AddConfigPath(workingDirectiory + "/config")
	viper.AutomaticEnv()

	// Change _ underscore in env to . dot notation in viper
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// Read config
	if err = viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// unmarshal the configuration from config.whatsapp.yaml
	logger.Info("Unmarshalling WhatsApp API configuration")
	err = viper.Unmarshal(&WAconfig)
	if err != nil {
		return nil, err
	}

	// translate the configuration into whatsapp.Session type
	CurrWASession := whatsapp.Session{
		ClientId:    WAconfig.ClientID,
		ClientToken: WAconfig.ClientToken,
		ServerToken: WAconfig.ServerToken,
		EncKey:      WAconfig.EncKey,
		MacKey:      WAconfig.MacKey,
		Wid:         WAconfig.Wid,
	}

	// establish new whatsapp connection, timeout in 20 seconds
	logger.Info("Establishing new WhatsApp connection")
	wac, err := whatsapp.NewConn(20 * time.Second)

	if err != nil {
		return nil, err
	}

	// make the signal that interrupt while login
	qrChan := make(chan string)

	if CurrWASession.ClientId != "" {
		// if session available in the configuration, restoring the old session

		logger.Info("Restoring WhatsApp session")
		newSess, err := wac.RestoreWithSession(CurrWASession)

		if err != nil {

			// if the session expired, generate a new QR code
			logger.Info("Session expired, generating QR code")

			// if login success, get the config
			config, err = GenerateQR(logger, qrChan, wac)
			if err != nil {
				return nil, err
			}
		} else {

			// if restoring old session success, get the config
			config, err = yaml.Marshal(&newSess)
			if err != nil {
				return nil, err
			}
		}
	} else {
		// if session unavailable in the configuration, create a new session

		config, err = GenerateQR(logger, qrChan, wac)
		if err != nil {
			return nil, err
		}
	}

	// write the new whatsapp session into the config.whatsapp.yaml file
	logger.Info("Writing new WhatsApp session in yaml")
	err = ioutil.WriteFile("./config/config.whatsapp.yaml", config, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return &Whatsapp{
		Wac:    wac,
		logger: logger,
	}, nil
}

// GenerateQR is a function to generate whatsapp session QR
func GenerateQR(logger hclog.Logger, qrChan chan string, wac *whatsapp.Conn) ([]byte, error) {
	// trigger the asynchronous function if the signal interrupt
	go func() {
		qr := <-qrChan
		var err error

		// remove an existing QR code
		os.Remove("./QR/WA_QRCode.png")

		//Show qr code or save it somewhere to scan
		err = qrcode.WriteFile(qr, qrcode.Medium, 256, "./QR/WA_QRCode.png")
		// if error, exit the application
		if err != nil {
			log.Fatal(err)
		} else {
			logger.Info("QR Code successfully generated, you can use the QR code to login your WhatsApp by scanning the code through your device")
		}

	}()

	// generating QR code
	logger.Info("Generating login QR Code")
	sess, err := wac.Login(qrChan)

	if err != nil {
		return nil, err
	}

	// marshaling the new session into the configuration type
	logger.Info("Marshalling WhatsApp session")
	config, err = yaml.Marshal(&sess)
	if err != nil {
		return nil, err
	}

	return config, nil
}
