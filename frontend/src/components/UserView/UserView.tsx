import React from 'react';
import { useGetCurrentUser } from '../../queries/user/getCurrentUser';

export const UserView: React.FC = () => {
  const { data, loading, error } = useGetCurrentUser();

  return (
    <>
      {data?.id}
      {data?.email}
      {data?.name}
      {data?.surname}
      {data?.role}

      {loading}
      {error}
    </>
  );
};
