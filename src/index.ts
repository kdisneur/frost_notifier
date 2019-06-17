import database, { IConfig as IDatabaseConfig } from './database'
import { isCurrentNight } from './night'
import openweather, { IConfig as IOpenweatherConfig, IProbe } from './openweather'
import twilio, { IConfig as ITwilioConfig } from './twilio'

export interface IPlace {
  city: string
  country: string
}

export interface IConfig {
  database: IDatabaseConfig
  openweather: IOpenweatherConfig
  twilio: ITwilioConfig
}

const Message = 'Hey! Pense à protéger ton parebrise! Il semble y avoir un risque de gel ce soir.'

const keepValidProbe = (date: Date, probe: IProbe): boolean => {
  return isCurrentNight(date, probe.date) && hasBadForecast(probe)
}

const hasBadForecast = (probe: IProbe): boolean => {
  return probe.temperature < 0 || (probe.temperature <= 2 && probe.humidity >= 80)
}

const sendFrostNotification = (config: IConfig, place: IPlace, recipient: string): Promise<string> => {
  const now = new Date()

  return database
    .hasItem(config.database, now)
    .then(alreadySent => {
      if (alreadySent) {
        return 'skipping... notification already sent'
      }

      return openweather
        .forecast(config.openweather, place.city, place.country)
        .then(probes => probes.filter(probe => keepValidProbe(now, probe)))
        .then(probes => sendMessage(config, recipient, now, probes))
    })
}

const sendMessage = (config: IConfig, recipient: string, date: Date, probes: IProbe[]): Promise<string> => {
  if (probes.length === 0) {
    return Promise.resolve('skipping... weather seems good enough')
  }

  return twilio
    .sendMessage(config.twilio, recipient, Message)
    .then(() => database.putItem(config.database, date, true))
    .then(() => 'SMS sent')
}

export default sendFrostNotification
