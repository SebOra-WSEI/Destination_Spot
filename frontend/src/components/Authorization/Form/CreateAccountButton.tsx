import { Button, CardActions, Divider } from '@mui/material';
import React from 'react';
import { useHistory } from 'react-router';
import { routes } from '../../../utils/routes';
import { BUTTON_RADIUS, FONT_FAMILY } from '../../../utils/consts';

export const CreateAccountButton: React.FC = () => {
  const history = useHistory();

  const handelClick = (): void => history.push(routes.register);

  return (
    <>
      <Divider style={styles.divider}>
        <span style={styles.span}>OR</span>
      </Divider>
      <CardActions>
        <Button
          variant='outlined'
          color='success'
          style={styles.registerButton}
          onClick={handelClick}
        >
          Create a new account
        </Button>
      </CardActions>
    </>
  );
};

const styles = {
  divider: {
    margin: '0.3rem 0 0.5rem 0',
  },
  span: {
    color: '#6c757d',
    fontFamily: FONT_FAMILY,
  },
  registerButton: {
    width: '100%',
    fontFamily: FONT_FAMILY,
    borderRadius: BUTTON_RADIUS,
  },
};
