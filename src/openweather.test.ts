import request, { IOptions, IResponse } from './request'
import openweather, { IConfig } from './openweather'
import fs from 'fs'

it('gets a promise with the next forecast', () => {
  const config: IConfig = {
    apiKey: 'valid-api-key',
    logger: {
      debug: (msg: string): void => { return },
      info: (msg: string): void => { return },
      error: (msg: string): void => { return }
    },
    requester: {
      get: (options: IOptions): Promise<IResponse> => {
        expect(options.uri).toEqual('https://api.openweathermap.org/data/2.5/forecast')
        expect(options.qs).toEqual({
          q: 'paris,fr',
          units: 'metric',
          appid: 'valid-api-key'
        })

        const content = fs.readFileSync('fixtures/openweather/paris-fr-valid-api-key.json', { encoding: 'utf8' })

        return Promise.resolve(content)
      }
    }
  }

  return openweather
    .forecast(config, 'paris', 'fr')
    .then(resp => {
      return expect(resp)
        .toEqual([
          { date: new Date('2019-06-02T15:00:00.000Z'),
            dt: 1559487600,
            temperature: 30.37,
            humidity: 44
          },
          { date: new Date('2019-06-02T18:00:00.000Z'),
            dt: 1559498400,
            temperature: 27.18,
            humidity: 54
          },
          { date: new Date('2019-06-02T21:00:00.000Z'),
            dt: 1559509200,
            temperature: 22.8,
            humidity: 64
          },
          { date: new Date('2019-06-03T00:00:00.000Z'),
            dt: 1559520000,
            temperature: 19.33,
            humidity: 75
          },
          { date: new Date('2019-06-03T03:00:00.000Z'),
            dt: 1559530800,
            temperature: 17.01,
            humidity: 79
          },
          { date: new Date('2019-06-03T06:00:00.000Z'),
            dt: 1559541600,
            temperature: 17.04,
            humidity: 73
          },
          { date: new Date('2019-06-03T09:00:00.000Z'),
            dt: 1559552400,
            temperature: 17.39,
            humidity: 79
          },
          { date: new Date('2019-06-03T12:00:00.000Z'),
            dt: 1559563200,
            temperature: 19.97,
            humidity: 66
          },
          { date: new Date('2019-06-03T15:00:00.000Z'),
            dt: 1559574000,
            temperature: 21.81,
            humidity: 54
          },
          { date: new Date('2019-06-03T18:00:00.000Z'),
            dt: 1559584800,
            temperature: 20.05,
            humidity: 60
          },
          { date: new Date('2019-06-03T21:00:00.000Z'),
            dt: 1559595600,
            temperature: 16.75,
            humidity: 69
          },
          { date: new Date('2019-06-04T00:00:00.000Z'),
            dt: 1559606400,
            temperature: 15.21,
            humidity: 76
          },
          { date: new Date('2019-06-04T03:00:00.000Z'),
            dt: 1559617200,
            temperature: 14.36,
            humidity: 79
          },
          { date: new Date('2019-06-04T06:00:00.000Z'),
            dt: 1559628000,
            temperature: 15.45,
            humidity: 78
          },
          { date: new Date('2019-06-04T09:00:00.000Z'),
            dt: 1559638800,
            temperature: 15.56,
            humidity: 84
          },
          { date: new Date('2019-06-04T12:00:00.000Z'),
            dt: 1559649600,
            temperature: 15.13,
            humidity: 90
          },
          { date: new Date('2019-06-04T15:00:00.000Z'),
            dt: 1559660400,
            temperature: 15.75,
            humidity: 85
          },
          { date: new Date('2019-06-04T18:00:00.000Z'),
            dt: 1559671200,
            temperature: 15.65,
            humidity: 77
          },
          { date: new Date('2019-06-04T21:00:00.000Z'),
            dt: 1559682000,
            temperature: 12.95,
            humidity: 78
          },
          { date: new Date('2019-06-05T00:00:00.000Z'),
            dt: 1559692800,
            temperature: 11.85,
            humidity: 80
          },
          { date: new Date('2019-06-05T03:00:00.000Z'),
            dt: 1559703600,
            temperature: 12.05,
            humidity: 82
          },
          { date: new Date('2019-06-05T06:00:00.000Z'),
            dt: 1559714400,
            temperature: 12.75,
            humidity: 85
          },
          { date: new Date('2019-06-05T09:00:00.000Z'),
            dt: 1559725200,
            temperature: 13.05,
            humidity: 87
          },
          { date: new Date('2019-06-05T12:00:00.000Z'),
            dt: 1559736000,
            temperature: 12.85,
            humidity: 89
          },
          { date: new Date('2019-06-05T15:00:00.000Z'),
            dt: 1559746800,
            temperature: 13.05,
            humidity: 91
          },
          { date: new Date('2019-06-05T18:00:00.000Z'),
            dt: 1559757600,
            temperature: 13.85,
            humidity: 88
          },
          { date: new Date('2019-06-05T21:00:00.000Z'),
            dt: 1559768400,
            temperature: 13.25,
            humidity: 87
          },
          { date: new Date('2019-06-06T00:00:00.000Z'),
            dt: 1559779200,
            temperature: 12.1,
            humidity: 84
          },
          { date: new Date('2019-06-06T03:00:00.000Z'),
            dt: 1559790000,
            temperature: 10.42,
            humidity: 90
          },
          { date: new Date('2019-06-06T06:00:00.000Z'),
            dt: 1559800800,
            temperature: 11.2,
            humidity: 84
          },
          { date: new Date('2019-06-06T09:00:00.000Z'),
            dt: 1559811600,
            temperature: 15.55,
            humidity: 61
          },
          { date: new Date('2019-06-06T12:00:00.000Z'),
            dt: 1559822400,
            temperature: 18.46,
            humidity: 46
          },
          { date: new Date('2019-06-06T15:00:00.000Z'),
            dt: 1559833200,
            temperature: 17.75,
            humidity: 52
          },
          { date: new Date('2019-06-06T18:00:00.000Z'),
            dt: 1559844000,
            temperature: 16.45,
            humidity: 60
          },
          { date: new Date('2019-06-06T21:00:00.000Z'),
            dt: 1559854800,
            temperature: 14.15,
            humidity: 70
          },
          { date: new Date('2019-06-07T00:00:00.000Z'),
            dt: 1559865600,
            temperature: 13.25,
            humidity: 73
          },
          { date: new Date('2019-06-07T03:00:00.000Z'),
            dt: 1559876400,
            temperature: 12.1,
            humidity: 82
          },
          { date: new Date('2019-06-07T06:00:00.000Z'),
            dt: 1559887200,
            temperature: 13.25,
            humidity: 80
          },
          { date: new Date('2019-06-07T09:00:00.000Z'),
            dt: 1559898000,
            temperature: 17.06,
            humidity: 61
          },
          { date: new Date('2019-06-07T12:00:00.000Z'),
            dt: 1559908800,
            temperature: 20.35,
            humidity: 50
          }
        ])
    })
})
