package config

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"
	"strconv"
)

type Account struct {
	Name   string `xml:"name"`
	Id     string `xml:"client_id"`
	Secret string `xml:"client_secret"`
}

type Config struct {
	Accounts []*Account `xml:"account"`
}

func New() *Config {
	defaultAccount := Account{
		Name:   "ENTER NAME",
		Id:     "ENTER CLIENT ID",
		Secret: "ENTER CLIENT SECRET",
	}
	accountSlice := make([]*Account, 0)
	accountSlice = append(accountSlice, &defaultAccount)
	cfg := Config{
		Accounts: accountSlice,
	}
	return &cfg
}

func Parse(data []byte) (*Config, error) {
	var cfg *Config
	err := xml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, errors.New("cfg parsing error, fill cfg. error: " + err.Error())
	}
	if &cfg == nil {
		return nil, errors.New("nil cfg parsed from config input: \r\n" + string(data))
	}

	for num, acc := range cfg.Accounts {
		if len(acc.Secret) < 20 || len(acc.Id) < 20 || len(acc.Name) < 1 {
			return nil, errors.New("account number " + strconv.Itoa(num+1) + " have incorrect filled fields. length conditions: secret>20 id>20 name>0")
		}
	}
	return cfg, nil
}

func GetConfig(fileName string) (*Config, error) {
	file, err := os.Open(fileName)
	if err != nil {
		file, err = os.Create(fileName)
		if err != nil {
			return nil, err
		}
		cfg := New()
		cfg.Save(fileName)
		return cfg, errors.New("fill config please")
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		cfg := New()
		cfg.Save(fileName)
		return cfg, errors.New("fill config please")
	}
	cfg, err := Parse(data)
	if err != nil {
		cfg = New()
	}
	return cfg, err
}

func (c *Config) Save(fileName string) error {
	data, err := c.Serialize()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, data, os.ModeAppend)
}

func (c *Config) Serialize() ([]byte, error) {
	return xml.MarshalIndent(c, "", "\t")
}
