import axios from 'axios';
import { useEffect, useState } from 'react';
import { endpoints } from '../../utils/routes';
import { CookieName, getCookieValueByName } from '../../utils/cookies';
import { ErrorResponse } from '../../types/response';
import { User, UserResponse } from '../../types/user';
import { Query, QueryVariables } from '../../types/query';
import { StatusCode } from '../../utils/consts';

export const useGetUserById = ({
  skip = false,
  variables,
}: QueryVariables): Query<User> => {
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [error, setError] = useState<string>('');
  const [user, setUser] = useState<User>();

  const token = getCookieValueByName(CookieName.Token);

  useEffect(() => {
    if (!skip) {
      axios
        .get<UserResponse>(endpoints.user(variables.id), {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        })
        .then((res) => {
          const { status, data } = res;

          if (status === StatusCode.OK) {
            setIsLoading(false);
            setUser(data.response.user);
          }
        })
        .catch((err: ErrorResponse) => {
          const { status, data } = err.response;

          if (status !== StatusCode.OK) {
            setError(data.error);
            setIsLoading(false);
          }
        });
    } else {
      setIsLoading(false);
    }
  }, [token, variables.id, skip]);

  return {
    data: user,
    loading: isLoading,
    error,
  };
};
