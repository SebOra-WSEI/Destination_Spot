import { useEffect, useState } from 'react';
import { EmptyQueryVariables, Query } from '../../types/query';
import axios from 'axios';
import { endpoints } from '../../utils/routes';
import { CookieName, getCookieValueByName } from '../../utils/cookies';
import { ErrorResponse } from '../../types/response';
import { StatusCode } from '../../types/statusCode';
import { Reservation, ReservationResponse } from '../../types/reservation';

export const useGetAllReservations = ({
  skip = false,
}: EmptyQueryVariables = {}): Query<Array<Reservation>> => {
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [error, setError] = useState<string>('');
  const [reservations, setReservations] = useState<Array<Reservation>>();

  const token = getCookieValueByName(CookieName.Token);

  useEffect(() => {
    if (!skip) {
      axios
        .get<ReservationResponse>(endpoints.reservations, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        })
        .then((res) => {
          const { status, data } = res;
          if (status === StatusCode.OK) {
            setIsLoading(false);
            setReservations(data.response.reservations);
          }
        })
        .catch((err: ErrorResponse) => {
          const { status, data } = err.response;

          if (status !== StatusCode.OK) {
            setError(data.error);
            setIsLoading(false);
          }
        });
    }
  }, [skip, token]);

  return {
    data: reservations,
    loading: isLoading,
    error,
  };
};
