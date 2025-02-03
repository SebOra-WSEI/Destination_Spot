import { useEffect, useState } from 'react';
import { Query } from '../../types/query';
import axios from 'axios';
import { endpoints } from '../../utils/routes';
import { CookieName, getCookieValueByName } from '../../utils/cookies';
import { ErrorResponse } from '../../types/response';
import { Spot, SpotsResponse } from '../../types/spot';
import { StatusCode } from '../../utils/consts';

export const useGetAllSpots = (): Query<Array<Spot>> => {
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [error, setError] = useState<string>('');
  const [spots, setSpots] = useState<Array<Spot>>();

  const token = getCookieValueByName(CookieName.Token);

  useEffect(() => {
    axios
      .get<SpotsResponse>(endpoints.spots, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then((res) => {
        const { status, data } = res;

        if (status === StatusCode.OK) {
          setIsLoading(false);
          setSpots(data.response.spots);
        }
      })
      .catch((err: ErrorResponse) => {
        const { status, data } = err.response;

        if (status !== StatusCode.OK) {
          setError(data.error);
          setIsLoading(false);
        }
      });
  }, [token]);

  return {
    data: spots,
    loading: isLoading,
    error,
  };
};
