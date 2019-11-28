package internal

import (
	"bufio"
	"fmt"
	"github.com/google/go-jsonnet"
	"io/ioutil"
	"os"
	"strings"
)

const credentialsJsonnetTemplate = `
{
  localcache: {
    path: std.extVar("CACHE_FOLDER"),
  },
  openweather: {
    api_key: std.extVar("OPENWEATHER_APIKEY"),
  },
  twilio: {
    account_sid: std.extVar("TWILIO_ACCOUNTSID"),
    token: std.extVar("TWILIO_TOKEN"),
    virual_phone_number: std.extVar("TWILIO_VIRTUALPHONENUMBER"),
  }
}
`

func InitializeCredentialsCmd(destination string) error {
	reader := bufio.NewReader(os.Stdin)

	openWeatherAPIKey, err := readUserInput(reader, "Openweather APIKey (can be found https://home.openweathermap.org/api_keys)")
	if err != nil {
		return fmt.Errorf("can't read openweather api-key input: %w", err)
	}

	twilioAccountSID, err := readUserInput(reader, "Twilio Account SID (can be found https://www.twilio.com/console)")
	if err != nil {
		return fmt.Errorf("can't read twilio account sid input: %w", err)
	}

	twilioToken, err := readUserInput(reader, "Twilio Token (can be found https://www.twilio.com/consolE)")
	if err != nil {
		return fmt.Errorf("can't read twilio token input: %w", err)
	}

	twilioPhoneNumber, err := readUserInput(reader, "Twilio Virtual Phone Number (can be found https://www.twilio.com/console/phone-numbers/incoming)")
	if err != nil {
		return fmt.Errorf("can't read twilio virtual phone number input: %w", err)
	}

	defaultCacheFolder := "/usr/local/var/frost-notifier/caching/last-text-message"
	cacheFolder, err := readUserInputWithDefault(reader, fmt.Sprintf("Cache folder (default: %s)\n", defaultCacheFolder), defaultCacheFolder)
	if err != nil {
		return fmt.Errorf("can't read cache folder input: %w", err)
	}

	vm := jsonnet.MakeVM()
	vm.ExtVar("CACHE_FOLDER", cacheFolder)
	vm.ExtVar("OPENWEATHER_APIKEY", openWeatherAPIKey)
	vm.ExtVar("TWILIO_ACCOUNTSID", twilioAccountSID)
	vm.ExtVar("TWILIO_TOKEN", twilioToken)
	vm.ExtVar("TWILIO_VIRTUALPHONENUMBER", twilioPhoneNumber)
	credentialsContent, _ := vm.EvaluateSnippet("credentials.jsonnet", credentialsJsonnetTemplate)
	if err := ioutil.WriteFile(destination, []byte(credentialsContent), 0644); err != nil {
		return fmt.Errorf("can't generate credentials.json file: %w", err)
	}

	return nil
}

func readUserInputWithDefault(reader *bufio.Reader, sentence string, defaultValue string) (string, error) {
	value, err := readUserInput(reader, sentence)
	if err != nil {
		return "", err
	}

	if value == "" {
		return defaultValue, nil
	}

	return value, nil
}

func readUserInput(reader *bufio.Reader, sentence string) (string, error) {
	fmt.Println(sentence)
	inputRaw, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(inputRaw, "\n"), nil
}
