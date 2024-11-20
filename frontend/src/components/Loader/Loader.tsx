import { CircularProgress } from '@mui/material';
import React from 'react';

export const Loader: React.FC = () => {
  return <CircularProgress sx={sxStyles} color='info' />;
};

const sxStyles = {
  position: 'absolute',
  top: '50%',
  left: '50%',
};
