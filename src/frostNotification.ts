import * as database from "./database";
import * as night from "./night";
import * as openweather from "./openweather";
import * as twilio from "./twilio";

export interface Place {
  city: string;
  country: string;
}

export interface Config {
  database: database.Config;
  openweather: openweather.Config;
  twilio: twilio.Config;
}

const Message =
  "Hey! Pense à protéger ton parebrise! Il semble y avoir un risque de gel ce soir.";

const hasBadForecast = (probe: openweather.Probe): boolean => {
  return (
    probe.temperature < 0 || (probe.temperature <= 2 && probe.humidity >= 80)
  );
};

const keepValidProbe = (date: Date, probe: openweather.Probe): boolean => {
  return night.isCurrentNight(date, probe.date) && hasBadForecast(probe);
};

const sendTextMessage = (
  config: Config,
  recipient: string,
  date: Date,
  probes: openweather.Probe[]
): Promise<string> => {
  if (probes.length === 0) {
    return Promise.resolve("skipping... weather seems good enough");
  }

  return twilio
    .sendMessage(config.twilio, recipient, Message)
    .then(() => database.putItem(config.database, date, true))
    .then(() => "SMS sent");
};

export const run = (
  config: Config,
  place: Place,
  recipient: string
): Promise<string> => {
  const now = new Date();

  return database.hasItem(config.database, now).then(alreadySent => {
    if (alreadySent) {
      return "skipping... notification already sent";
    }

    return openweather
      .forecast(config.openweather, place.city, place.country)
      .then(probes => probes.filter(probe => keepValidProbe(now, probe)))
      .then(probes => sendTextMessage(config, recipient, now, probes));
  });
};
