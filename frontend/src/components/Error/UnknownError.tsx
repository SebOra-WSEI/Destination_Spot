import React from 'react';
import { ErrorCard } from './ErrorCard';

export const UnknownError: React.FC<{ link?: string }> = ({ link }) => (
  <ErrorCard isErrorCard text='Unknown Error' link={link} />
);
