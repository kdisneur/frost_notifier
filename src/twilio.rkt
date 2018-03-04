#lang racket

(require net/base64
         net/http-client
         net/uri-codec)

(provide twilio
         send-message)

(define-struct twilio (account-sid access-token from-number))

(define (send-message config to-number text)
  (let ([credentials (user-authentication config)]
        [encoded-from-number (form-urlencoded-encode
                               (twilio-from-number config))]
        [encoded-to-number (form-urlencoded-encode to-number)])
    (http-sendrecv "api.twilio.com"
                   (twilio-path config)
                   #:ssl? #t
                   #:version "1.1"
                   #:method "POST"
                   #:headers
                   (list (format "Authorization: Basic ~a" credentials)
                     "Content-Type: application/x-www-form-urlencoded")
                   #:data (alist->form-urlencoded
                            (list (cons 'Body text)
                                  (cons 'To encoded-to-number)
                                  (cons 'From encoded-from-number))))))

(define (twilio-path config)
  (format "/2010-04-01/Accounts/~a/Messages.json" (twilio-account-sid config)))

(define (user-authentication config)
  (bytes->string/utf-8
    (base64-encode
      (string->bytes/utf-8
        (format "~a:~a"
                (twilio-account-sid config)
                (twilio-access-token config)))
      #"")))
