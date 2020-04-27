package setting

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
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
	DbName   string `yaml:"db_name"`
	DbArgs   string `yaml:"db_args"`
}

// NewDatabase ...
func NewDatabase(file string) (*DatabaseSettings, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var d DatabaseSettings
	err = yaml.Unmarshal(buf, &d)
	if err != nil {
		return nil, err
	}

	return &d, nil
}

// ConnectString ...
func (d *DatabaseSettings) ConnectString() string {
	c := &d.Database.Connect
	return fmt.Sprintf("%s:%s@%s(%s:%s)/%s?%s", c.User, c.Password, c.Protocol, c.Host, c.Port, c.DbName, c.DbArgs)
}

// DbmsString ...
func (d *DatabaseSettings) DbmsString() string {
	return d.Database.Dbms
}
