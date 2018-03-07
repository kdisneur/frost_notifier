#lang racket

(require racket/include)
(require "open-weather.rkt")
(require "timestamp.rkt")
(require "twilio.rkt")

(include (file "/var/run/secrets/frost_notification"))

(define open-weather-config
  (open-weather open-weather-api-key))

(define twilio-config
  (twilio twilio-account-sid twilio-access-token twilio-from-number))

(define message
  "Hey! Pense à protéger ton parebrise! Il semble y avoir un risque de gel ce soir.")

(define (keep-only-tonight-forecast forecast)
  (let ([in-next-hours? (lambda (weather)
                         (before-tonight-limit? (hash-ref weather 'dt)))])
    (filter in-next-hours? forecast)))

(define (risk-of-frost?)
  (let* ([next-days-forecast
           (5-days-3-hours-forecast open-weather-config open-weather-city)]
         [tonight-forecast
           (keep-only-tonight-forecast next-days-forecast)]
         [bad-forecast? (lambda (weather)
                          (let* ([main (hash-ref weather 'main)]
                                 [temperature (hash-ref main 'temp)]
                                 [humidity (hash-ref main 'humidity)])
                            (cond [(< temperature 0) #t]
                                  [(and (< temperature 2) (>= humidity 80)) #t]
                                  [else #f])))]
         [bad-forecast-only (filter bad-forecast? tonight-forecast)])
    (not (empty? bad-forecast-only))))

(define (main)
  (cond [(already-sent-tonight?)
         (println "Skipping. SMS already sent")]
        [(risk-of-frost?)
         (send-message twilio-config twilio-to-number message)
         (update-date)
         (println "SMS sent")]
        [else (println "Skipping. Weather seems good enough")]))

(main)
