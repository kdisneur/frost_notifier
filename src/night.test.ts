import { isCurrentNight, night } from './night'

it('night: returns the same date before midnight', () => {
  const beforeMidnight = 'August 4, 1988 22:35:24'

  expect(night(new Date(beforeMidnight))).toEqual(new Date(beforeMidnight))
})

it('night: returns yesterday date after midnight', () => {
  const yesterday = 'August 4, 1988 05:12:24'
  const afterMidnight = 'August 5, 1988 05:12:24'

  expect(night(new Date(afterMidnight))).toEqual(new Date(yesterday))
})

it('isCurrentNight: returns true when the dates are in the same night', () => {
  const now = 'August 4, 1988 22:35:24'
  const other = 'August 5, 1988 02:21:24'

  expect(isCurrentNight(new Date(now), new Date(other))).toBeTruthy()
})

it('isCurrentNight: returns true when the dates are in two different nights', () => {
  const now = 'August 4, 1988 22:35:24'
  const other = 'August 5, 1988 21:21:24'

  expect(isCurrentNight(new Date(now), new Date(other))).toBeFalsy()
})
