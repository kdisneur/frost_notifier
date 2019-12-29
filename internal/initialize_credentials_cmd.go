package internal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func InitializeCredentialsCmd(destination string) error {
	reader := bufio.NewReader(os.Stdin)

	var config CredentialsConfig
	config.LocalCache.Path = "/usr/local/var/frost-notifier/caching/last-text-message"

	var err error

	openweatherConfig := config.OpenWeather
	err = readUserInput(reader, "Openweather APIKey (can be found https://home.openweathermap.org/api_keys)", &openweatherConfig.APIKey)
	if err != nil {
		return fmt.Errorf("can't read openweather api-key input: %w", err)
	}

	twilioConfig := config.Twilio
	err = readUserInput(reader, "Twilio Account SID (can be found https://www.twilio.com/console)", &twilioConfig.AccoundSID)
	if err != nil {
		return fmt.Errorf("can't read twilio account sid input: %w", err)
	}

	err = readUserInput(reader, "Twilio Token (can be found https://www.twilio.com/console)", &twilioConfig.Token)
	if err != nil {
		return fmt.Errorf("can't read twilio token input: %w", err)
	}

	err = readUserInput(reader, "Twilio Virtual Phone Number (can be found https://www.twilio.com/console/phone-numbers/incoming)", &twilioConfig.Sender)
	if err != nil {
		return fmt.Errorf("can't read twilio virtual phone number input: %w", err)
	}

	localCacheConfig := config.LocalCache
	err = readUserInputKeepOnEmpty(reader, fmt.Sprintf("Cache folder (default: %s)\n", config.LocalCache.Path), &localCacheConfig.Path)
	if err != nil {
		return fmt.Errorf("can't read cache folder input: %w", err)
	}

	config.OpenWeather = openweatherConfig
	config.Twilio = twilioConfig
	config.LocalCache = localCacheConfig

	file, err := json.MarshalIndent(&config, "", "  ")
	if err != nil {
		return fmt.Errorf("can't generate configuration JSON: %w", err)
	}

	if err := ioutil.WriteFile(destination, file, 0644); err != nil {
		return fmt.Errorf("can't generate %s file: %w", destination, err)
	}

	return nil
}

func readUserInputKeepOnEmpty(reader *bufio.Reader, sentence string, receiver *string) error {
	value := *receiver

	if err := readUserInput(reader, sentence, &value); err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	*receiver = value

	return nil
}

func readUserInput(reader *bufio.Reader, sentence string, receiver *string) error {
	fmt.Println(sentence)
	inputRaw, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	value := strings.TrimSuffix(inputRaw, "\n")
	*receiver = value

	return nil
}
