import { useEffect, useState } from 'react';
import { Query } from '../../types/query';
import axios from 'axios';
import { endpoints } from '../../utils/routes';
import { CookieName, getCookieValueByName } from '../../utils/cookies';
import { ErrorResponse } from '../../types/response';
import { User, UsersResponse } from '../../types/user';
import { StatusCode } from '../../utils/consts';

export const useGetAllUsers = (): Query<Array<User>> => {
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [error, setError] = useState<string>('');
  const [users, setUsers] = useState<Array<User>>();

  const token = getCookieValueByName(CookieName.Token);

  useEffect(() => {
    axios
      .get<UsersResponse>(endpoints.users, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then((res) => {
        const { status, data } = res;

        if (status === StatusCode.OK) {
          setIsLoading(false);
          setUsers(data.response.users);
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
    data: users,
    loading: isLoading,
    error,
  };
};
