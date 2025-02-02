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

  const enabledSpots = spots?.filter(
    (spot) =>
      !getSpotsIdsForEachDay(reservations)[date]?.includes(spot.id)
  );

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

const getSpotsIdsForEachDay = (reservations?: Array<Reservation>): ReservationsGroupedByDay => {
  const reservationsGroupedByDay: ReservationsGroupedByDay = {};

  reservations?.forEach((reservation) => {
    const date = new Date(
      Number(reservation.details.reservedFrom) * 1000
    ).toDateString();

    if (!reservationsGroupedByDay[date]) {
      reservationsGroupedByDay[date] = [];
    }

    reservationsGroupedByDay[date].push(reservation.spot.id);
  });

  return reservationsGroupedByDay;
};