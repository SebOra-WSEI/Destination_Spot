import { useEffect, useState } from 'react';
import { Query, QueryVariables } from '../../types/query';
import axios from 'axios';
import { endpoints } from '../../utils/routes';
import { CookieName, getCookieValueByName } from '../../utils/cookies';
import { ErrorResponse } from '../../types/response';
import { Reservation, ReservationResponse } from '../../types/reservation';
import { StatusCode } from '../../utils/consts';

export const useGetReservationById = ({
  variables,
}: QueryVariables): Query<Reservation> => {
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [error, setError] = useState<string>('');
  const [reservation, setReservation] = useState<Reservation>();

  const token = getCookieValueByName(CookieName.Token);

  useEffect(() => {
    axios
      .get<ReservationResponse>(endpoints.reservation(variables.id), {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then((res) => {
        const { status, data } = res;
        if (status === StatusCode.OK) {
          setIsLoading(false);
          setReservation(data.response.reservation);
        }
      })
      .catch((err: ErrorResponse) => {
        const { status, data } = err.response;

        if (status !== StatusCode.OK) {
          setError(data.error);
          setIsLoading(false);
        }
      });
  }, [token, variables.id]);

  return {
    data: reservation,
    loading: isLoading,
    error,
  };
};
