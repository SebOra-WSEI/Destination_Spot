import React, { useState } from 'react';
import {
  Avatar,
  Box,
  Divider,
  IconButton,
  Menu,
  MenuItem,
  Tooltip,
  Typography
} from '@mui/material';
import PersonIcon from '@mui/icons-material/Person';
import { FONT_FAMILY } from '../../utils/consts';
import { signOut } from '../../utils/signOut';

export const SettingIcon: React.FC = () => {
  const [anchorEl, setAnchorEl] = useState<HTMLElement | null>(null);

  const handleOpen = (event: React.MouseEvent<HTMLElement>) =>
    setAnchorEl(event.currentTarget);

  const handleClose = () =>
    setAnchorEl(null);

  return (
    <Box>
      <Tooltip title="Settings">
        <IconButton onClick={handleOpen}>
          <Avatar sx={styles.avatar}>
            <PersonIcon />
          </Avatar>
        </IconButton>
      </Tooltip>
      <Menu
        sx={styles.menu}
        anchorOrigin={{
          vertical: 'top',
          horizontal: 'right',
        }}
        anchorEl={anchorEl}
        open={Boolean(anchorEl)}
        onClose={handleClose}
      >
        <MenuItem onClick={() => console.log('reset')}>
          <Typography sx={{ fontFamily: FONT_FAMILY }}>Reset Password</Typography>
        </MenuItem>
        <Divider />
        <MenuItem onClick={signOut}>
          <Typography sx={{ fontFamily: FONT_FAMILY }}>Logout</Typography>
        </MenuItem>
      </Menu>
    </Box >
  )
}

const styles = {
  avatar: {
    width: '2rem',
    height: '2rem',
    marginRight: '1rem'
  },
  menu: {
    marginTop: '2.5rem'
  }
}