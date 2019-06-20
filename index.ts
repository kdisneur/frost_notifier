import "source-map-support/register";
import dotenv from "dotenv";
import * as frostNotification from "./src/frostNotification";
import winston from "winston";

dotenv.config();

const readEnvVariable = (name: string): string => {
  const value = process.env[name];
  if (!value) {
    throw new Error(`missing environment variable: ${name}`);
  }

  return value;
};

const logger = winston.createLogger({
  level: readEnvVariable("FROST_NOTIFICATION_LOGGER"),
  format: winston.format.json(),
  transports: [
    new winston.transports.Console({ format: winston.format.simple() })
  ]
});

const config = {
  database: {
    tableName: readEnvVariable("FROST_NOTIFICATION_TABLE_NAME"),
    logger: logger
  },
  openweather: {
    apiKey: readEnvVariable("OPENWEATHER_API_KEY"),
    logger: logger
  },
  twilio: {
    accoundSID: readEnvVariable("TWILIO_ACCOUND_SID"),
    accessToken: readEnvVariable("TWILIO_ACCESS_TOKEN"),
    phoneNumber: readEnvVariable("TWILIO_PHONE_NUMBER"),
    logger: logger
  }
};

export const triggerNotification = (): Promise<void> => {
  const place = {
    city: readEnvVariable("FROST_NOTIFICATION_CITY"),
    country: readEnvVariable("FROST_NOTIFICATION_COUNTRY")
  };

  logger.info("starting frost analysis");

  return frostNotification
    .run(config, place, readEnvVariable("FROST_NOTIFICATION_PHONE_NUMBER"))
    .then(message => {
      logger.info(message);
    })
    .catch(message => {
      logger.error(message);
    });
};
