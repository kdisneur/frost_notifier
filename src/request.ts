import request from 'request-promise-native'

export interface IOptions {
  auth?: {
    user: string
    pass: string
  },
  formData?: any
  qs?: any
  uri: string
}

export interface IGetter {
  get: (options: IOptions) => Promise<IResponse>
}

export interface IPoster {
  post: (options: IOptions) => Promise<IResponse>
}

export interface IRequester extends IGetter, IPoster {
}

export type IResponse = string

const requester: IRequester = {
  get: (options: IOptions): Promise<IResponse> => {
    return request(options).then(body => body)
  },
  post: (options: IOptions): Promise<IResponse> => {
    return request({ ...options, method: 'POST' }).then(body => body)
  }
}

export default requester
