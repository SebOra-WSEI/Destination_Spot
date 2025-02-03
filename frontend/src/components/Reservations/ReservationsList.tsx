import React from 'react';
import { IconButton, List, Tooltip } from '@mui/material';
import { Reservation } from '../../types/reservation';
import { SpotChip } from './SpotChip/SpotChip';
import { ReservationUserDetails } from './ReservationUserDetails/ReservationUserDetails';
import { ReservationDate } from './ReservationDate/ReservationDate';
import { ReservationListItem } from './ReservationListItem';
import dayjs from 'dayjs';
import DeleteIcon from '@mui/icons-material/Delete';
import { CookieName, getCookieValueByName } from '../../utils/cookies';
import { useRemoveReservation } from '../../queries/reservation/useRemoveReservation';
import { ErrorCard } from '../Error/ErrorCard';
import { routeBuilder } from '../../utils/routes';

interface ReservationsListProps {
  reservations: Array<Reservation> | undefined;
}

export const ReservationsList: React.FC<ReservationsListProps> = ({
  reservations,
}) => {
  const userId = getCookieValueByName(CookieName.UserId);

  const sortedCurrentReservations = reservations?.sort((a, b) => {
    if (isTheSameDay(a, b)) {
      return a.spot.id - b.spot.id;
    }

    return Number(a.details.reservedFrom) - Number(b.details.reservedTo);
  });

  const { remove } = useRemoveReservation();

  if (!reservations?.length) {
    return <ErrorCard isErrorCard text='There are no reservations yet' link={routeBuilder.addReservations} />
  }

  const handleRemove = async (id: number): Promise<void> =>
    await remove(id)

  return (
    <List sx={styles.list}>
      {sortedCurrentReservations?.map(({ details, spot, user }) => (
        <ReservationListItem key={details.id}>
          <ReservationDate reservedFrom={details.reservedFrom} />
          <SpotChip location={spot.location} />
          <ReservationUserDetails user={user} />
          {user.id.toString() === userId && (
            <Tooltip title="Remove reservation">
              <IconButton sx={styles.icon} onClick={() => handleRemove(details.id)}>
                <DeleteIcon />
              </IconButton>
            </Tooltip>
          )}
        </ReservationListItem>
      ))}
    </List>
  );
};

function isTheSameDay(a: Reservation, b: Reservation): boolean {
  return dayjs.unix(Number(a.details.reservedFrom)).format('dddd, D MMM YYYY') ===
    dayjs.unix(Number(b.details.reservedFrom)).format('dddd, D MMM YYYY')
}

const styles = {
  list: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    marginTop: '2.5rem',
  },
  icon: {
    marginLeft: 'auto'
  },
};
