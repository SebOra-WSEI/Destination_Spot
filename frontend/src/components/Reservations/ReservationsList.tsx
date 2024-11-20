import React from 'react';
import { List } from '@mui/material';
import { Reservation } from '../../types/reservation';
import { SpotChip } from './SpotChip/SpotChip';
import { ReservationUserDetails } from './ReservationUserDetails/ReservationUserDetails';
import { ReservationDate } from './ReservationDate/ReservationDate';
import { ReservationListItem } from './ReservationListItem';

interface ReservationsListProps {
  reservations: Array<Reservation> | undefined;
}

export const ReservationsList: React.FC<ReservationsListProps> = ({
  reservations,
}) => {
  return (
    <List sx={sxStyles.list}>
      {reservations?.map(({ details, spot, user }) => (
        <ReservationListItem key={details.id}>
          <ReservationDate reservedFrom={details.reservedFrom} />
          <SpotChip location={spot.location} />
          <ReservationUserDetails user={user} />
        </ReservationListItem>
      ))}
    </List>
  );
};

const sxStyles = {
  list: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    marginTop: '2.5rem',
  },
};
