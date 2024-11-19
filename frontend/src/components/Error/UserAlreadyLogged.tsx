import React from 'react';
import { Box, Button, Card, CardActions, CardContent } from '@mui/material';
import { useHistory } from 'react-router';
import { routeBuilder } from '../../utils/routes';
import { TOKEN_KEY } from '../../utils/consts';
import { CookieName, eraseCookie } from '../../utils/cookies';

export const UserAlreadyLogged: React.FC = () => {
  const history = useHistory();

  const handleClick = (): void => history.push(routeBuilder.profile);

  return (
    <Box sx={{}}>
      <Card>
        <CardContent>
          <h3 style={{}}>User already logged in</h3>
        </CardContent>
        <CardActions>
          <Button
            variant='contained'
            color='error'
            size='small'
            sx={{}}
            onClick={signOut}
          >
            Log out
          </Button>
          <Button variant='outlined' size='small' sx={{}} onClick={handleClick}>
            Move to profile page
          </Button>
        </CardActions>
      </Card>
    </Box>
  );
};

const signOut = (): void => {
  window.localStorage.removeItem(TOKEN_KEY);
  Object.values(CookieName).forEach((val) => eraseCookie(val));
  window.location.replace(routeBuilder.login);
};
