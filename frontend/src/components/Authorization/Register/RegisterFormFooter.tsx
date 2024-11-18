import { Button, CardActions, Divider, Typography } from '@mui/material';
import React from 'react';
import { useHistory } from 'react-router';
import { routeBuilder } from '../../../utils/routes';
import { FONT_FAMILY } from '../../../utils/consts';


export const RegisterFormFooter: React.FC = () => {
  const history = useHistory()

  const handleClick = (): void => history.push(routeBuilder.login)

  return (
    <>
      <Divider style={styles.divider} />
      <CardActions>
        <Typography style={styles.loginMessage}>
          Already have account?
          <Button
            onClick={handleClick}
            size='small'
            style={styles.loginMessageLink}
          >
            Login
          </Button>
        </Typography>
      </CardActions>
    </>
  );
}

const styles = {
  divider: {
    margin: '0.3rem 0 0.5rem 0',
  },
  loginMessage: {
    fontFamily: FONT_FAMILY,
    marginLeft: 'auto',
    fontSize: '0.9rem',
  },
  loginMessageLink: {
    fontFamily: FONT_FAMILY,
    fontSize: '0.9rem',
    padding: 0,
  },
}