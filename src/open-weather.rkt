#lang racket

(require json
         net/http-client)

(provide open-weather
         5-days-3-hours-forecast)

(define-struct open-weather (api-key))

(define (5-days-3-hours-forecast config city)
  (define-values (status headers payload)
    (http-sendrecv "api.openweathermap.org"
                   (path config city)))
  (hash-ref (read-json payload) 'list))

(define (path config city)
  (format "/data/2.5/forecast?units=~a&id=~a&APPID=~a"
          "metric"
          city
          (open-weather-api-key config)))
