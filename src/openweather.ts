import request, { IGetter } from './request'
import { ILogger } from './logger'

export interface IConfig {
  apiKey: string
  requester?: IGetter
  logger: ILogger
}

export interface IOpenweather {
  forecast: (config: IConfig, city: string, country: string) => Promise<IProbe[]>
}

export interface IProbe {
  dt: number
  date: Date
  temperature: number
  humidity: number
}

interface IOpenweatherProbeAPI {
  dt: number
  main: {
    temp: number
    humidity: number
  }
}

const forecast = (config: IConfig, city: string, country: string): Promise<IProbe[]> => {
  const requester = config.requester ? config.requester : request

  const query = {
    q: `${city},${country}`,
    units: 'metric',
    appid: config.apiKey
  }

  config.logger.debug('start searching on openweather')

  return requester
    .get({
      uri: 'https://api.openweathermap.org/data/2.5/forecast',
      qs: query
    })
    .then(body => JSON.parse(body))
    .then(response => {
      config.logger.debug('response received from openweather')

      return response
    })
    .then(response => response.list.map(probeFromAPI))
}

const probeFromAPI = (rawProbe: IOpenweatherProbeAPI): IProbe => {
  const date = new Date()
  date.setTime(rawProbe.dt * 1000)

  return {
    date: date,
    dt: rawProbe.dt,
    temperature: rawProbe.main.temp,
    humidity: rawProbe.main.humidity
  }
}

const openweather: IOpenweather = { forecast }
export default openweather
