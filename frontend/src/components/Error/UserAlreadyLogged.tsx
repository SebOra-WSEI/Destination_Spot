import React from 'react';
import { Button, CardActions, CardContent } from '@mui/material';
import { useHistory } from 'react-router';
import { routeBuilder } from '../../utils/routes';
import { BUTTON_RADIUS, FONT_FAMILY } from '../../utils/consts';
import { CenteredCard } from '../Card/CenteredCard';
import { signOut } from '../../utils/signOut';

export const UserAlreadyLogged: React.FC = () => {
  const history = useHistory();

  const handleClick = (): void => history.push(routeBuilder.profile);

  return (
    <CenteredCard>
      <CardContent>
        <h3 style={styles.header}>User already logged in</h3>
      </CardContent>
      <CardActions>
        <Button
          variant='contained'
          color='error'
          size='small'
          sx={styles.button}
          onClick={signOut}
        >
          Log out
        </Button>
        <Button
          variant='outlined'
          size='small'
          sx={styles.button}
          onClick={handleClick}
        >
          Move to profile page
        </Button>
      </CardActions>
    </CenteredCard>
  );
};

const styles = {
  button: {
    borderRadius: BUTTON_RADIUS,
    fontFamily: FONT_FAMILY,
    marginLeft: 'auto',
  },
  header: {
    fontSize: '1.5rem',
    marginBottom: '-0.5rem',
    fontFamily: FONT_FAMILY,
    color: '#757575',
  },
}
