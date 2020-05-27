package config

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type Config struct {
	UUID     string
	UUIDFile string
	Endpoint string
	Token    string
	LogFile  string
	PID      int
	PidFile  string

	Heartbeat chan interface{}
	Register  chan interface{}
	Log       chan interface{}
}

func NewConfig(configReader *viper.Viper) (*Config, error) {
	UUIDFile := configReader.GetString("uuidfile")
	if UUIDFile == "" {
		UUIDFile = "agentd.uuid"
	}

	PidFile := configReader.GetString("pidfile")
	if UUIDFile == "" {
		PidFile = "agentd.pid"
	}

	LogFile := configReader.GetString("logfile")
	if UUIDFile == "" {
		LogFile = "logs/agent.log"
	}

	Endpoint := configReader.GetString("endpoint")
	if UUIDFile == "" {
		Endpoint = "http://localhost:8080/v1/api"
	}

	Token := configReader.GetString("token")

	UUID := ""
	if cxt, err := ioutil.ReadFile(UUIDFile); err == nil {
		UUID = string(cxt)
	} else if os.IsNotExist(err) {
		//strings.Replace(uuid.New().String(), "-", "", -1)
		UUID = strings.ReplaceAll(uuid.New().String(), "-", "")
		ioutil.WriteFile(UUIDFile, []byte(UUID), os.ModePerm)
	} else {
		return nil, err
	}

	// 获取PID
	PID := os.Getgid()
	ioutil.WriteFile(PidFile, []byte(strconv.Itoa(PID)), os.ModePerm)

	return &Config{
		Endpoint:  Endpoint,
		Token:     Token,
		UUID:      UUID,
		UUIDFile:  UUIDFile,
		LogFile:   LogFile,
		PID:       PID,
		PidFile:   PidFile,
		Heartbeat: make(chan interface{}, 64),
		Register:  make(chan interface{}, 64),
		Log:       make(chan interface{}, 10240),
	}, nil
}
