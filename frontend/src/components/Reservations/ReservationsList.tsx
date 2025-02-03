import React from 'react';
import { IconButton, List, Tooltip } from '@mui/material';
import { Reservation } from '../../types/reservation';
import { SpotChip } from './SpotChip/SpotChip';
import { ReservationUserDetails } from './ReservationUserDetails/ReservationUserDetails';
import { ReservationDate } from './ReservationDate/ReservationDate';
import dayjs from 'dayjs';
import InfoIcon from '@mui/icons-material/Info';
import { ErrorCard } from '../Error/ErrorCard';
import { routeBuilder, routes } from '../../utils/routes';
import { useHistory } from 'react-router';
import { CommonListItem } from '../List/CommonListItem';

interface ReservationsListProps {
  reservations: Array<Reservation> | undefined;
}

export const ReservationsList: React.FC<ReservationsListProps> = ({
  reservations,
}) => {
  const history = useHistory();

  const sortedCurrentReservations = reservations?.sort((a, b) => {
    if (isTheSameDay(a, b)) {
      return a.spot.id - b.spot.id;
    }

    return Number(a.details.reservedFrom) - Number(b.details.reservedTo);
  });

  if (!reservations?.length) {
    return <ErrorCard isErrorCard text='There are no reservations yet' link={routes.addReservations} />
  }

  const handleClick = (id: number): void =>
    history.push(routeBuilder.reservationDetails(String(id)))

  return (
    <List sx={styles.list}>
      {sortedCurrentReservations?.map(({ details, spot, user }) => (
        <CommonListItem key={details.id}>
          <ReservationDate reservedFrom={details.reservedFrom} />
          <SpotChip location={spot.location} />
          <ReservationUserDetails user={user} />
          <Tooltip title="Reservation details">
            <IconButton sx={styles.icon} onClick={() => handleClick(details.id)}>
              <InfoIcon />
            </IconButton>
          </Tooltip>
        </CommonListItem>
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
