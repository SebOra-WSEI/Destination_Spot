import React from 'react';
import { CookieName, getCookieValueByName } from '../../utils/cookies';
import { UserNotLogged } from '../Error/UserNotLogged';
import { ReservationsList } from './ReservationsList';
import { useGetAllReservations } from '../../queries/reservation/useGetAllReservations';
import { ErrorCard } from '../Error/ErrorCard';
import { Loader } from '../Loader/Loader';

export const ReservationsView: React.FC = () => {
  const userId = getCookieValueByName(CookieName.UserId);

  const { data, loading, error } = useGetAllReservations();

  if (!userId) {
    return <UserNotLogged />;
  }

  if (loading) {
    return <Loader />;
  }

  if (error) {
    return <ErrorCard text={error} />;
  }

  return <ReservationsList reservations={data} />;
};
