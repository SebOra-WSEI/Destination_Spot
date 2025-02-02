import React, { useState } from 'react';
import { DateCalendar } from '@mui/x-date-pickers/DateCalendar';
import { dayPickerClasses, LocalizationProvider, pickersDayClasses } from '@mui/x-date-pickers';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import dayjs from 'dayjs';
import { badgeClasses, } from '@mui/material';
import { pickersToolbarClasses } from '@mui/x-date-pickers/internals';
import { FONT_FAMILY } from '../../../utils/consts';
import { CookieName, getCookieValueByName } from '../../../utils/cookies';
import { UserNotLogged } from '../../Error/UserNotLogged';
import { CalendarBadge } from './CalendarBadge';
import { useGetAllSpots } from '../../../queries/spots/useGetAllSpots';
import { useGetAllReservations } from '../../../queries/reservation/useGetAllReservations';
import { Loader } from '../../Loader/Loader';
import { ErrorCard } from '../../Error/ErrorCard';
import { ReservationModal } from './ReservationModal';

export const AddReservationView: React.FC = () => {
  const userId = getCookieValueByName(CookieName.UserId);
  const today = dayjs();

  const [selectedDate, setSelectedDate] = useState<dayjs.Dayjs | null>(today);
  const [isModalOpen, setModalOpen] = useState<boolean>(false);

  const {
    data: spots,
    error: spotsError,
    loading: spotsLoading
  } = useGetAllSpots()

  const {
    data: reservations,
    error: reservationsError,
    loading: reservationsLoading
  } = useGetAllReservations()

  const handleDateClick = (newDate: dayjs.Dayjs | null) => {
    setSelectedDate(newDate);
    setModalOpen(true);
  };

  if (!userId) {
    return <UserNotLogged />;
  }

  const isLoading = spotsLoading || reservationsLoading;
  const isError = spotsError || reservationsError;

  if (isLoading) {
    return <Loader />
  }

  if (isError) {
    return <ErrorCard text={isError} />;
  }

  return (
    <>
      <LocalizationProvider dateAdapter={AdapterDayjs}>
        <DateCalendar
          sx={styles.calendar}
          onChange={handleDateClick}
          value={selectedDate}
          shouldDisableDate={(date) => shouldDisableDate(date, today)}
          slots={{
            day: (props) => (
              <CalendarBadge
                dayProps={props}
                spots={spots}
                reservations={reservations}
                isDateDisabled={shouldDisableDate(dayjs(props.day), today)}
              />
            )
          }}
        />
      </LocalizationProvider>
      <ReservationModal
        isModalOpen={isModalOpen}
        setModalOpen={setModalOpen}
        selectedDate={selectedDate}
        reservations={reservations}
        spots={spots}
      />
    </>
  )
}

function shouldDisableDate(date: dayjs.Dayjs, today: dayjs.Dayjs): boolean {
  const dayOfWeek = date.day();

  return (
    !(dayOfWeek >= 1 && dayOfWeek <= 5) ||
    date.isBefore(today, 'day') ||
    date.diff(today, 'day') > 6
  );
}

const styles = {
  calendar: {
    boxShadow: '0 0 20px 10px rgba(0, 0, 0, 0.1)',
    borderRadius: '1rem',
    padding: '2rem',
    width: '23rem',
    marginTop: '8rem',
    [`& .${pickersDayClasses.root}`]: {
      width: 40,
      height: 40,
    },
    [`& .${pickersDayClasses.dayWithMargin}`]: {
      margin: 0.2,
      fontSize: 12,
    },
    [`& .${badgeClasses.badge}`]: {
      height: 18,
      minWidth: 18,
      fontSize: 10,
    },
    [`& .${dayPickerClasses.weekContainer}`]: {
      padding: 0.1,
    },
    [`& .${pickersToolbarClasses.content}`]: {
      fontFamily: FONT_FAMILY,
      textAlign: 'center',
      color: '#428bca',
    }
  },
}