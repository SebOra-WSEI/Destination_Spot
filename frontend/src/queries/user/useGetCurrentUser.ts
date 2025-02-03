import { CookieName, getCookieValueByName } from '../../utils/cookies';
import { User } from '../../types/user';
import { Query } from '../../types/query';
import { useGetUserById } from './useGetUserById';

export const useGetCurrentUser = ({ skip = false }): Query<User> => {
  const userId = getCookieValueByName(CookieName.UserId);

  const { data, loading, error } = useGetUserById({
    variables: { id: userId ?? '' },
    skip,
  });

  return { data, loading, error };
};
