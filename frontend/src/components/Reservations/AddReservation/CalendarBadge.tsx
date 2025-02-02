import React from "react";
import { Badge } from "@mui/material";
import { PickersDay, PickersDayProps } from '@mui/x-date-pickers';
import { Dayjs } from "dayjs";
import { Spot } from "../../../types/spot";
import { Reservation } from "../../../types/reservation";

interface ReservationsGroupedByDay {
  [day: string]: Array<number>;
}

interface CalendarBadgeProps {
  dayProps: PickersDayProps<Dayjs>;
  spots?: Array<Spot>;
  reservations?: Array<Reservation>;
  isDateDisabled: boolean
}

export const CalendarBadge: React.FC<CalendarBadgeProps> = ({
  dayProps,
  spots,
  reservations,
  isDateDisabled
}) => {
  const date = dayProps.day.toDate().toDateString();

  const enabledSpots = getEnabledSpots(date, reservations, spots);

  return (
    <Badge
      key={date}
      overlap='circular'
      color={enabledSpots?.length ? 'success' : 'error'}
      badgeContent={isDateDisabled ? undefined : enabledSpots?.length.toString()}
    >
      <PickersDay {...dayProps} />
    </Badge>
  );
}

const getEnabledSpots = (
  date: string,
  reservations?: Array<Reservation>,
  spots?: Array<Spot>,
): Array<Spot> => {
  const reservationsGroupedByDay: ReservationsGroupedByDay = {};

  reservations?.forEach((reservation) => {
    const reservationDate = new Date(
      Number(reservation.details.reservedFrom) * 1000
    ).toDateString();

    if (!reservationsGroupedByDay[reservationDate]) {
      reservationsGroupedByDay[reservationDate] = [];
    }

    reservationsGroupedByDay[reservationDate].push(reservation.spot.id);
  });

  return spots?.filter(
    (spot) =>
      !reservationsGroupedByDay[date]?.includes(spot.id)
  ) ?? [];
};