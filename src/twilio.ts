import * as logger from "./logger";
import * as request from "./request";

export interface Config {
  accoundSID: string;
  accessToken: string;
  phoneNumber: string;
  requester?: request.Poster;
  logger: logger.Logger;
}

export interface Twilio {
  sendMessage: (
    config: Config,
    recipientPhoneNumber: string,
    message: string
  ) => Promise<void>;
}

export const sendMessage = (
  config: Config,
  recipientPhoneNumber: string,
  message: string
): Promise<void> => {
  const requester = config.requester ? config.requester : request;

  config.logger.debug("start posting on twilio");

  return requester
    .post({
      uri: `https://api.twilio.com/2010-04-01/Accounts/${config.accoundSID}/Messages.json`,
      auth: {
        user: config.accoundSID,
        pass: config.accessToken
      },
      formData: {
        From: config.phoneNumber,
        To: recipientPhoneNumber,
        Body: message
      }
    })
    .then(body => JSON.parse(body))
    .then(response => {
      config.logger.debug("response received from twilio");

      return response;
    })
    .then(() => {});
};
