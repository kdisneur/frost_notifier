import request from "request-promise-native";

export interface Options {
  auth?: {
    user: string;
    pass: string;
  };
  formData?: object;
  qs?: object;
  uri: string;
}

export interface Getter {
  get: (options: Options) => Promise<Response>;
}

export interface Poster {
  post: (options: Options) => Promise<Response>;
}

export interface Requester extends Getter, Poster {}

export type Response = string;

export const get = (options: Options): Promise<Response> => {
  return request(options).then(body => body);
};

export const post = (options: Options): Promise<Response> => {
  return request({ ...options, method: "POST" }).then(body => body);
};
