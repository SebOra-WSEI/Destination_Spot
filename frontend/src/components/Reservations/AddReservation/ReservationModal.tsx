import React, { useState } from 'react';
import { CenteredModal } from '../../Modal/CenteredModal';
import dayjs from 'dayjs';
import {
  Button,
  DialogActions,
  DialogContent,
  FormControl,
  InputLabel,
  MenuItem,
  Select,
  SelectChangeEvent,
} from '@mui/material';
import { BUTTON_RADIUS, FONT_FAMILY } from '../../../utils/consts';
import { Reservation } from '../../../types/reservation';
import { Spot } from '../../../types/spot';
import { useCreateReservation } from '../../../queries/reservation/useCreateReservation';
import { isSameDay } from '../../../utils/isSameDay';
import { CookieName, getCookieValueByName } from '../../../utils/cookies';

interface ReservationModalProps {
  isModalOpen: boolean;
  setModalOpen: (isModalOpen: boolean) => void;
  selectedDate: dayjs.Dayjs | null;
  reservations?: Array<Reservation>;
  spots?: Array<Spot>;
}

export const ReservationModal: React.FC<ReservationModalProps> = ({
  isModalOpen,
  selectedDate,
  setModalOpen,
  reservations,
  spots,
}) => {
  const [selectedSpotId, setSelectedSpotId] = useState<string>('0');

  const date = selectedDate as dayjs.Dayjs;
  const userId = getCookieValueByName(CookieName.UserId);

  const onCloseModal = (): void => setModalOpen(false);

  const handleChange = (evt: SelectChangeEvent) =>
    setSelectedSpotId(evt.target.value);

  const reservationsForSelectedDay = filterReservationsForSelectedDay(
    selectedDate,
    reservations
  );

  const availableSpots =
    spots?.filter(
      (spot) =>
        !reservationsForSelectedDay.map((r) => r.spot.id)?.includes(spot.id)
    ) ?? [];

  const spotReservedByUser = reservationsForSelectedDay.find(
    (reservation) => reservation.user.id.toString() === userId
  )?.spot.location;

  const { reserve } = useCreateReservation();

  const spotId =
    Number(selectedSpotId) > 0 ? selectedSpotId : availableSpots[0].id;

  const handleSubmit = async (
    evt: React.FormEvent<HTMLFormElement>
  ): Promise<void> => {
    evt.preventDefault();
    await reserve({
      userId: Number(getCookieValueByName(CookieName.UserId)),
      spotId: Number(spotId),
      reservedFrom: String(createDate(date, 0, 0, 0)),
      reservedTo: String(createDate(date, 23, 59, 59)),
    });
  };

  return (
    <CenteredModal isModalOpen={isModalOpen} handleSubmit={handleSubmit}>
      <DialogContent>
        <h3 style={styles.header}>
          {date.format('dddd, D MMM YYYY').toString()}
        </h3>
        {availableSpots.length ? (
          <>
            {spotReservedByUser ? (
              <p>
                You have already reserved spot number{' '}
                <strong>{spotReservedByUser}</strong>.
              </p>
            ) : (
              <FormControl fullWidth>
                <InputLabel style={styles.label}>Spot Number</InputLabel>
                <Select
                  label='Spot Number'
                  style={styles.label}
                  value={spotId.toString()}
                  onChange={handleChange}
                >
                  {availableSpots.map(({ id, location }) => (
                    <MenuItem key={id} style={styles.label} value={String(id)}>
                      {location}
                    </MenuItem>
                  ))}
                </Select>
              </FormControl>
            )}
          </>
        ) : (
          <p>All parking spots are already reserved.</p>
        )}
      </DialogContent>
      <DialogActions>
        <Button variant='outlined' onClick={onCloseModal} style={styles.button}>
          Close
        </Button>
        <Button
          variant='contained'
          color='success'
          type='submit'
          style={styles.button}
          disabled={!availableSpots.length || !!spotReservedByUser}
        >
          Reserve
        </Button>
      </DialogActions>
    </CenteredModal>
  );
};

const filterReservationsForSelectedDay = (
  selectedDate: dayjs.Dayjs | null,
  reservations?: Array<Reservation>
) => {
  if (!reservations || !selectedDate) {
    return [];
  }

  const filteredReservations = reservations?.filter((reservation) =>
    isSameDay(reservation.details.reservedFrom, selectedDate)
  );

  return filteredReservations;
};

const createDate = (
  date: dayjs.Dayjs,
  h: number,
  min: number,
  sec: number
): number => {
  const day = date.date();
  const month = date.month();
  const year = date.year();

  return new Date(year, month, day, h, min, sec).getTime() / 1000;
};

const styles = {
  header: {
    display: 'flex',
    justifyContent: 'center',
  },
  label: {
    fontFamily: FONT_FAMILY,
  },
  button: {
    fontFamily: FONT_FAMILY,
    borderRadius: BUTTON_RADIUS,
  },
};
