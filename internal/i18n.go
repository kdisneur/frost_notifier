package internal

var i18n I18n

const (
	TranslationKeyCoverYourWindscreen TranslationKey = "cover_your_windscreen"
)

func init() {
	i18n.translations = make(map[Locale]Translations)
	i18n.translations[Locale("en")] = Translations(map[TranslationKey]string{
		TranslationKeyCoverYourWindscreen: "You should cover your windscreen, it's going to be freezing cold!",
	})

	i18n.translations[Locale("fr")] = Translations(map[TranslationKey]string{
		TranslationKeyCoverYourWindscreen: "Tu devrais couvrir ton parebrise, il va cailler 🥶",
	})
}

type Locale string

type TranslationKey string

type Translations map[TranslationKey]string

type I18n struct {
	translations map[Locale]Translations
}

func (i *I18n) Locale(locale Locale) Translations {
	return i.translations[locale]
}

func (t Translations) Translate(key TranslationKey) string {
	return t[key]
}
