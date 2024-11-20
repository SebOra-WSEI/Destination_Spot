import React from 'react';
import { Box, Chip } from '@mui/material';
import DirectionsCarIcon from '@mui/icons-material/DirectionsCar';

export const SpotChip: React.FC<{ location: number }> = ({ location }) => (
  <Box sx={sxStyles.iconItem}>
    <Chip
      color='success'
      variant='outlined'
      size='medium'
      icon={<DirectionsCarIcon />}
      label={`Spot ${location}`}
    />
  </Box>
);

const sxStyles = {
  iconItem: {
    width: '20%',
  },
};
