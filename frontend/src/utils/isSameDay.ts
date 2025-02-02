import dayjs from 'dayjs';

export const isSameDay = (reservedFrom: string, selectedDay: dayjs.Dayjs) => {
  const unix = new Date(reservedFrom).getTime();
  const reservationDate = dayjs.unix(unix);

  return (
    reservationDate.date() === selectedDay.date() &&
    reservationDate.month() === selectedDay.month() &&
    reservationDate.year() === selectedDay.year()
  );
};
