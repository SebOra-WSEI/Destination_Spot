import React from 'react';
import { Button, CardActions, CardContent } from '@mui/material';
import { useHistory } from 'react-router';
import { routes } from '../../utils/routes';
import { BUTTON_RADIUS, FONT_FAMILY } from '../../utils/consts';
import { CenteredCard } from '../Card/CenteredCard';
import { signOut } from '../../utils/signOut';

export const UserAlreadyLogged: React.FC = () => {
  const history = useHistory();

  const handleClick = (): void => history.push(routes.profile);

  return (
    <CenteredCard isErrorCard>
      <CardContent>
        <h3 style={styles.header}>User already logged in</h3>
      </CardContent>
      <CardActions>
        <Button
          size='small'
          style={styles.button()}
          onClick={handleClick}
        >
          Move to profile page
        </Button>
        <Button
          color='error'
          size='small'
          style={styles.button(true)}
          onClick={signOut}
        >
          Sign out
        </Button>
      </CardActions>
    </CenteredCard>
  );
};

const styles = {
  button: (isLastButton?: boolean) => ({
    borderRadius: BUTTON_RADIUS,
    fontFamily: FONT_FAMILY,
    ...(isLastButton ? { marginLeft: 'auto' } : {}),
  }),
  header: {
    fontSize: '1.5rem',
    marginBottom: '-0.5rem',
    marginTop: '-0.5rem',
    fontFamily: FONT_FAMILY,
    color: '#757575',
    display: 'flex',
    justifyContent: 'center',
  },
};
