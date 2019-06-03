export const isCurrentNight = (today: Date, date: Date): boolean => {
  const currentNight = night(today)
  const otherNight = night(date)

  return currentNight.getFullYear() === otherNight.getFullYear() &&
    currentNight.getMonth() === otherNight.getMonth() &&
    currentNight.getDate() === otherNight.getDate()
}

export const night = (today: Date): Date => {
  if (today.getHours() < 8) {
    const yesterday = new Date(today)
    yesterday.setDate(today.getDate() - 1)

    return yesterday
  }

  return today
}
