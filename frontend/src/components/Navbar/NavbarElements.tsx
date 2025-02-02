import React from 'react';
import DirectionsCarIcon from '@mui/icons-material/DirectionsCar';
import { Box, Button } from '@mui/material';
import { FONT_FAMILY } from '../../utils/consts';
import { useHistory } from 'react-router';

export const NavbarElements: React.FC<{ pages: Array<string> }> = ({
  pages,
}) => {
  const history = useHistory();

  return (
    <>
      <DirectionsCarIcon sx={{ ...styles.bigScreen, ...styles.icon }} />
      <Box sx={{ ...styles.box, ...styles.bigScreen }}>
        {pages.map((page) => (
          <Button
            key={page}
            onClick={() => history.push(page)}
            style={styles.button(pages.length)}
          >
            {page.split('-').join(' ')}
          </Button>
        ))}
      </Box>
    </>
  );
};

const styles = {
  bigScreen: {
    display: {
      xs: 'none',
      md: 'flex',
    },
  },
  icon: {
    margin: '0 1rem 0 1rem',
  },
  box: {
    flexGrow: 1,
  },
  button: (pageCount: number) => ({
    color: 'white',
    fontFamily: FONT_FAMILY,
    ...(pageCount === 1
      ? {
        marginLeft: 'auto',
        marginRight: '1.5rem',
      }
      : {}),
  }),
};
