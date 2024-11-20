import React from 'react';
import { ErrorCard } from './ErrorCard';
import { routeBuilder } from '../../utils/routes';

export const UserNotLogged: React.FC = () => (
  <ErrorCard
    isErrorCard
    text='Please log in'
    link={routeBuilder.login}
  />
);