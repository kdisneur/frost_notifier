package internal

import (
	"encoding/json"
	"fmt"
	"frostnotifier/internal/localcache"
	"frostnotifier/internal/openweather"
	"frostnotifier/internal/twilio"
	"io/ioutil"
)

type Config struct {
	ConfigFile  string
	CountryCode string
	Debug       bool
	ZipCode     string
	Language    string
	Credentials struct {
		LocalCache struct {
			Path string `json:"path"`
		} `json:"localcache"`
		OpenWeather struct {
			APIKey string `json:"api_key"`
		} `json:"openweather"`
		Twilio struct {
			AccoundSID string `json:"account_sid"`
			Token      string `json:"token"`
			Sender     string `json:"virual_phone_number"`
		} `json:"twilio"`
	}
	Recipient string
}

func (c *Config) NewLocalCache() *localcache.LocalCache {
	return localcache.New(c.Credentials.LocalCache.Path)
}

func (c *Config) NewOpenWeather() *openweather.OpenWeather {
	return openweather.New(c.Credentials.OpenWeather.APIKey, c.CountryCode, c.ZipCode)
}

func (c *Config) NewTwilio() *twilio.Twilio {
	return twilio.New(c.Credentials.Twilio.AccoundSID, c.Credentials.Twilio.Token, c.Credentials.Twilio.Sender)
}

func (c *Config) Load() error {
	rawCredentials, err := ioutil.ReadFile(c.ConfigFile)
	if err != nil {
		return fmt.Errorf("can't read %s: %w", c.ConfigFile, err)
	}

	if json.Unmarshal(rawCredentials, &c.Credentials) != nil {
		return fmt.Errorf("can't parse credentials (%s): %w", c.ConfigFile, err)
	}

	return nil
}

func (c *Config) Translations() Translations {
	return i18n.Locale(Locale(c.Language))
}
