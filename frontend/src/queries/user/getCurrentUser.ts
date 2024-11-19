import { CookieName, getCookieValueByName } from '../../utils/cookies';
import { User } from '../../types/user';
import { EmptyQueryVariables, Query } from '../../types/query';
import { useGetUserById } from './getUserById';

export const useGetCurrentUser = ({
  skip = false,
}: EmptyQueryVariables = {}): Query<User> => {
  const userId = getCookieValueByName(CookieName.UserId);

  const { data, loading, error } = useGetUserById({
    variables: { id: userId ?? '' },
    skip,
  });

  return { data, loading, error };
};
