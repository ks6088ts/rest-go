package setting

import (
	"fmt"

	"github.com/spf13/viper"
)

// DatabaseSettings ...
type DatabaseSettings struct {
	Database Database `yaml:"database"`
}

// Database ...
type Database struct {
	Dbms    string  `yaml:"dbms"`
	Connect Connect `yaml:"connect"`
}

// Connect ...
type Connect struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Protocol string `yaml:"protocol"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Dbargs   string `yaml:"dbargs"`
}

// NewDatabase ...
func NewDatabase(file string) (*DatabaseSettings, error) {
	viper.SetConfigFile(file)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	var settings DatabaseSettings
	if err := viper.Unmarshal(&settings); err != nil {
		return nil, err
	}
	return &settings, nil
}

// ConnectString ...
func (d *DatabaseSettings) ConnectString() string {
	c := &d.Database.Connect
	return fmt.Sprintf("%s:%s@%s(%s:%s)/%s?%s", c.User, c.Password, c.Protocol, c.Host, c.Port, c.Dbname, c.Dbargs)
}

// DbmsString ...
func (d *DatabaseSettings) DbmsString() string {
	return d.Database.Dbms
}
