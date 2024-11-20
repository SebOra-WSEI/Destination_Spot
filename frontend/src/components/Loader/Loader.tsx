import { CircularProgress } from '@mui/material';
import React from 'react';

export const Loader: React.FC = () => {
  return <CircularProgress sx={styles} color='info' />;
};

const styles = {
  position: 'absolute',
  top: '50%',
}