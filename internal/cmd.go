package internal

import (
	"context"
	"time"
)

func Main(ctx context.Context, config *Config) error {
	translator := config.Translations()

	frostNotifier := &FrostNotifier{
		LocalCache:  config.NewLocalCache(),
		Message:     translator.Translate(TranslationKeyCoverYourWindscreen),
		OpenWeather: config.NewOpenWeather(),
		Twilio:      config.NewTwilio(),
		Logger:      NewDefaultLogger(config.Debug),
	}

	return frostNotifier.Run(ctx, time.Now(), config.Recipient)
}
