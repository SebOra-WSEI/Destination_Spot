import { useEffect, useState } from 'react';
import { EmptyQueryVariables, Query } from '../../types/query';
import axios from 'axios';
import { endpoints } from '../../utils/routes';
import { CookieName, getCookieValueByName } from '../../utils/cookies';
import { ErrorResponse, } from '../../types/response';
import { StatusCode } from '../../types/statusCode';
import { Spot, SpotResponse } from '../../types/spot';

export const useGetAllSpots = ({
  skip = false,
}: EmptyQueryVariables = {}): Query<Array<Spot>> => {
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [error, setError] = useState<string>('');
  const [spots, setSpots] = useState<Array<Spot>>();

  const token = getCookieValueByName(CookieName.Token);

  useEffect(() => {
    if (!skip) {
      axios
        .get<SpotResponse>(endpoints.spots, {
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
    }
  }, [skip, token]);

  return {
    data: spots,
    loading: isLoading,
    error,
  };
};
