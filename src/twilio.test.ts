import fs from 'fs'
import request, { IOptions, IResponse } from './request'
import twilio, { IConfig } from './twilio'

it('sends a message using twilio', () => {
  const config: IConfig = {
    accoundSID: 'ACd567b79d7878c78778b898c234fe23ba',
    accessToken: 'a-valid-token',
    phoneNumber: '+330101010101',
    logger: {
      debug: (msg: string): void => { return },
      info: (msg: string): void => { return },
      error: (msg: string): void => { return }
    },
    requester: {
      post: (options: IOptions): Promise<IResponse> => {
        expect(options.uri)
          .toEqual('https://api.twilio.com/2010-04-01/Accounts/ACd567b79d7878c78778b898c234fe23ba/Messages.json')

        expect(options.auth).toEqual({
          user: 'ACd567b79d7878c78778b898c234fe23ba',
          pass: 'a-valid-token'
        })

        expect(options.formData).toEqual({
          From: '+330101010101',
          To: '+330202020202',
          Body: 'Hello World!'
        })

        const content = fs.readFileSync('fixtures/twilio/message-valid-api-key.json', { encoding: 'utf8' })

        return Promise.resolve(content)
      }
    }
  }

  return twilio
    .sendMessage(config, '+330202020202', 'Hello World!')
})
