import React from 'react';
import { List } from '@mui/material';
import { Reservation } from '../../types/reservation';
import { SpotChip } from './SpotChip/SpotChip';
import { ReservationUserDetails } from './ReservationUserDetails/ReservationUserDetails';
import { ReservationDate } from './ReservationDate/ReservationDate';
import { ReservationListItem } from './ReservationListItem';
import dayjs from 'dayjs';

interface ReservationsListProps {
  reservations: Array<Reservation> | undefined;
}

export const ReservationsList: React.FC<ReservationsListProps> = ({
  reservations,
}) => {
  const sortedCurrentReservations = reservations?.sort((a, b) => {
    if (isTheSameDay(a, b)) {
      return a.spot.id - b.spot.id;
    }

    return Number(a.details.reservedFrom) - Number(b.details.reservedTo);
  });


  return (
    <List sx={sxStyles.list}>
      {sortedCurrentReservations?.map(({ details, spot, user }) => (
        <ReservationListItem key={details.id}>
          <ReservationDate reservedFrom={details.reservedFrom} />
          <SpotChip location={spot.location} />
          <ReservationUserDetails user={user} />
        </ReservationListItem>
      ))}
    </List>
  );
};

function isTheSameDay(a: Reservation, b: Reservation): boolean {
  return dayjs.unix(Number(a.details.reservedFrom)).format('dddd, D MMM YYYY') ===
    dayjs.unix(Number(b.details.reservedFrom)).format('dddd, D MMM YYYY')
}

const sxStyles = {
  list: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    marginTop: '2.5rem',
  },
};
