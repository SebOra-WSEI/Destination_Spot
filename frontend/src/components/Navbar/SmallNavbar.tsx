import React, { useState } from 'react';
import { Box, IconButton, Menu, MenuItem, Typography } from '@mui/material';
import { FONT_FAMILY } from '../../utils/consts';
import DirectionsCarIcon from '@mui/icons-material/DirectionsCar';
import MenuIcon from '@mui/icons-material/Menu';
import { useHistory } from 'react-router';

export const SmallNavbar: React.FC<{ pages: Array<string> }> = ({ pages }) => {
  const [anchorEl, setAnchorEl] = useState<HTMLElement | null>(null);

  const history = useHistory();

  const handleOpen = (event: React.MouseEvent<HTMLElement>) =>
    setAnchorEl(event.currentTarget);

  const handleClose = () => {
    setAnchorEl(null);
  };

  const handleOnClick = (page: string) => {
    history.push('/' + page.toLocaleLowerCase());
    handleClose();
  };

  const mappedPagesNames = pages.map(
    (page) => page[0].toUpperCase() + page.slice(1)
  );

  return (
    <>
      <Box sx={sxStyles.box}>
        <IconButton onClick={handleOpen} style={styles.menuIcon}>
          <MenuIcon />
        </IconButton>
        <Menu
          anchorOrigin={{
            vertical: 'top',
            horizontal: 'left',
          }}
          anchorEl={anchorEl}
          open={Boolean(anchorEl)}
          onClose={handleClose}
          sx={sxStyles.menu}
        >
          {mappedPagesNames.map((page) => (
            <MenuItem key={page} onClick={() => handleOnClick(page)}>
              <Typography style={styles.text}>{page}</Typography>
            </MenuItem>
          ))}
        </Menu>
      </Box>
      <DirectionsCarIcon sx={sxStyles.carIcon} />
      <Typography variant='h6' sx={sxStyles.header}>
        Destination Spot
      </Typography>
    </>
  );
};

const sxStyles = {
  box: {
    flexGrow: 1,
    display: {
      xs: 'flex',
      md: 'none',
    },
  },
  menu: {
    display: {
      xs: 'block',
      md: 'none',
    },
    marginTop: '2rem',
  },
  carIcon: {
    display: {
      xs: 'flex',
      md: 'none',
    },
    marginRight: '1rem',
  },
  header: {
    display: { xs: 'flex', md: 'none' },
    flexGrow: 1,
    fontFamily: FONT_FAMILY,
    fontWeight: 700,
    letterSpacing: '0.1rem',
    color: 'inherit',
  },
};

const styles = {
  menuIcon: {
    marginLeft: '1rem',
    color: 'inherit',
  },
  text: {
    fontFamily: FONT_FAMILY,
  },
};
