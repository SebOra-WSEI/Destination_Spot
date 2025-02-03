import React from 'react';
import { ErrorCard } from './ErrorCard';
import { routes } from '../../utils/routes';

export const UserNotLogged: React.FC = () => (
  <ErrorCard isErrorCard text='Please log in' link={routes.login} />
);
