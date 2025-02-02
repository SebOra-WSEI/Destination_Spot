import React from 'react';
import { ErrorCard } from './ErrorCard';

export const PageNotFound: React.FC = () => (
  <ErrorCard isErrorCard text='Page not found' />
);
