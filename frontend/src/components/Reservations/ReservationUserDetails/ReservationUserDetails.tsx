import React from 'react';
import { Box, ListItemText, Typography } from '@mui/material';
import PersonIcon from '@mui/icons-material/Person';
import { User } from '../../../types/user';
import { FONT_FAMILY } from '../../../utils/consts';

export const ReservationUserDetails: React.FC<{ user: User }> = ({ user }) => (
  <Box sx={styles.additionalItem}>
    <ListItemText
      primary={
        <Box sx={sxStyles.userName}>
          <PersonIcon />
          <Typography variant='body2'>
            <strong>{user.name + ' ' + user.surname}</strong>
          </Typography>
        </Box>
      }
      secondary={<Typography style={styles.userEmail}>{user.email}</Typography>}
    />
  </Box>
);

const sxStyles = {
  userName: {
    display: 'flex',
    fontSize: '13px',
    alignItems: 'center',
    fontFamily: FONT_FAMILY,
    wordWrap: 'break-word',
  },
};

const styles = {
  additionalItem: {
    wordWrap: 'break-word',
  },
  userEmail: {
    fontSize: '16px',
    fontFamily: FONT_FAMILY,
  },
};
