export interface Logger {
  debug: (msg: string) => void;
  error: (msg: string) => void;
  info: (msg: string) => void;
}
