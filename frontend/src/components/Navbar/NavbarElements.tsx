import React from 'react';
import DirectionsCarIcon from '@mui/icons-material/DirectionsCar';
import { Box, Button } from '@mui/material';
import { FONT_FAMILY } from '../../utils/consts';
import { useHistory } from 'react-router';

export const NavbarElements: React.FC<{ pages: Array<string> }> = ({ pages }) => {
  const history = useHistory();

  return (
    <>
      <DirectionsCarIcon sx={{ ...styles.bigScreen, ...styles.icon }} />
      <Box sx={{ ...styles.box, ...styles.bigScreen }}>
        {pages.map((page) => (
          <Button
            key={page}
            onClick={() => history.push(page)}
            sx={styles.button}
          >
            {page}
          </Button>
        ))}
      </Box >
    </>
  )
}

const styles = {
  bigScreen: {
    display: {
      xs: 'none',
      md: 'flex'
    }
  },
  icon: {
    margin: '0 1rem 0 1rem'
  },
  box: {
    flexGrow: 1
  },
  button: {
    color: 'white',
    fontFamily: FONT_FAMILY
  }
}