#lang racket

(require racket/date)
(require racket/file)
(require "date-utils.rkt")

(provide already-sent-tonight?
         before-tonight-limit?
         update-date)

(define (already-sent-tonight?)
  (touch-cached-value-file)
  (equal? (cached-value) (today)))

(define (before-tonight-limit? timestamp)
  (equal? (current-night timestamp) (today)))

(define (cached-value)
  (file->string cached-value-path #:mode 'text))

(define cached-value-path
  (build-path (find-system-path 'temp-dir) "last-frost.dat"))

(define (current-night seconds)
  (let ([date (seconds->date seconds)])
    (cond [(< (date-hour date) 8)
           (date->string (shift-day date -1))]
          [else (date->string date)])))

(define (today)
  (current-night (current-seconds)))

(define (touch-cached-value-file)
  (cond [(file-exists? cached-value-path) cached-value-path]
        [else (write-to-file 0 cached-value-path)
              cached-value-path]))

(define (update-date)
  (display-to-file (today) cached-value-path #:mode 'text #:exists 'replace))
