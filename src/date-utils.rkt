#lang racket

(require racket/date)

(provide shift-day)

(define (shift-day date number-days)
  (let ([timestamp (date->seconds date)]
         [seconds   (* number-days 24 60 60)])
    (seconds->date (+ timestamp seconds))))
