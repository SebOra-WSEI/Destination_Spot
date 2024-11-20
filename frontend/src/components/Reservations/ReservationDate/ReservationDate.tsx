import { Box, ListItemText } from '@mui/material';
import React from 'react';
import { FONT_FAMILY } from '../../../utils/consts';
import dayjs from 'dayjs';

export const ReservationDate: React.FC<{ reservedFrom: string }> = ({
  reservedFrom,
}) => (
  <Box sx={sxStyles.dateItem}>
    <ListItemText
      primaryTypographyProps={{
        fontFamily: FONT_FAMILY,
      }}
      primary={dayjs.unix(Number(reservedFrom)).format('dddd, D MMM YYYY')}
    />
  </Box>
);

const sxStyles = {
  dateItem: {
    width: '30%',
    wordWrap: 'break-word',
  },
};
