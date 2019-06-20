import * as logger from "./logger";
import * as request from "./request";

export interface Config {
  apiKey: string;
  requester?: request.Getter;
  logger: logger.Logger;
}

export interface Openweather {
  forecast: (config: Config, city: string, country: string) => Promise<Probe[]>;
}

export interface Probe {
  dt: number;
  date: Date;
  temperature: number;
  humidity: number;
}

interface OpenweatherProbeAPI {
  dt: number;
  main: {
    temp: number;
    humidity: number;
  };
}

const probeFromAPI = (rawProbe: OpenweatherProbeAPI): Probe => {
  const date = new Date();
  date.setTime(rawProbe.dt * 1000);

  return {
    date: date,
    dt: rawProbe.dt,
    temperature: rawProbe.main.temp,
    humidity: rawProbe.main.humidity
  };
};

export const forecast = (
  config: Config,
  city: string,
  country: string
): Promise<Probe[]> => {
  const requester = config.requester ? config.requester : request;

  const query = {
    q: `${city},${country}`,
    units: "metric",
    appid: config.apiKey
  };

  config.logger.debug("start searching on openweather");

  return requester
    .get({
      uri: "https://api.openweathermap.org/data/2.5/forecast",
      qs: query
    })
    .then(body => JSON.parse(body))
    .then(response => {
      config.logger.debug("response received from openweather");

      return response;
    })
    .then(response => response.list.map(probeFromAPI));
};
