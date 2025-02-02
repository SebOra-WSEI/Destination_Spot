import React, { PropsWithChildren } from 'react';
import { Box, ListItem } from '@mui/material';

export const ReservationListItem: React.FC<PropsWithChildren> = ({
  children,
}) => (
  <ListItem sx={sxStyles.listItem}>
    <Box sx={sxStyles.reservationsDetails}>{children}</Box>
  </ListItem>
);

const sxStyles = {
  listItem: {
    boxShadow: '0.7rem 0.7rem 0.7rem rgba(0, 0, 0, 0.1)',
    borderRadius: '1rem',
    marginTop: '0.8rem',
    width: '45rem',
    padding: '0.5rem',
    '&:nth-of-type(odd)': {
      backgroundColor: '#FFFFFF',
    },
    '&:nth-of-type(even)': {
      backgroundColor: '#F3F3F3',
    },
  },
  reservationsDetails: {
    fontWeight: 'bold',
    width: '100%',
    display: 'flex',
    alignItems: 'center',
    flexWrap: 'wrap',
    padding: '1px',
    marginLeft: '1rem'
  },
};
