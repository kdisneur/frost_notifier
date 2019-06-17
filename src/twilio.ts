import { ILogger } from './logger'
import request, { IPoster } from './request'

export interface IConfig {
  accoundSID: string
  accessToken: string
  phoneNumber: string
  requester?: IPoster
  logger: ILogger
}

export interface ITwilio {
  sendMessage: (config: IConfig, recipientPhoneNumber: string, message: string) => Promise<void>
}

const sendMessage = (config: IConfig, recipientPhoneNumber: string, message: string): Promise<void> => {
  const requester = config.requester ? config.requester : request

  config.logger.debug('start posting on twilio')

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
      config.logger.debug('response received from twilio')

      return response
    })
    .then(() => { return })
}

const twilio: ITwilio = { sendMessage }
export default twilio
