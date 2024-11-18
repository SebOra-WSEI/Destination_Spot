import {
  Box,
  Button,
  Card,
  CardActions,
  CardContent,
  TextField,
} from '@mui/material';
import React from 'react';
import { FONT_FAMILY, MARGIN_TOP_CONTENT } from '../../../utils/consts';
import { AuthorizationBody } from '../../../types/authorization';
import { PasswordInput } from './PasswordInput';
import { useLocation } from 'react-router';
import { routeBuilder } from '../../../utils/routes';
import { CreateAccountButton } from './CreateAccountButton';

interface AuthorizationFormProps {
  body: AuthorizationBody;
  setBody: (body: AuthorizationBody) => void;
  handleSubmit: (event: React.FormEvent<HTMLFormElement>) => Promise<void>;
  header: string;
}

export const AuthorizationForm: React.FC<AuthorizationFormProps> = ({
  body,
  setBody,
  handleSubmit,
  header,
}) => {
  const { pathname } = useLocation();
  const { email, password } = body;

  const isLoginView =
    pathname === routeBuilder.login || pathname === routeBuilder.default;

  const handlePasswordChange = (
    evt: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ): void => {
    setBody({
      ...body,
      password: evt.target.value,
    });
  };

  return (
    <Box onSubmit={handleSubmit} component='form' style={styles.box}>
      <Card style={styles.card}>
        <CardContent>
          <h3 style={styles.header}>{header}</h3>
          <TextField
            label='Email'
            variant='standard'
            type='text'
            autoComplete='email'
            autoFocus
            required
            value={email}
            fullWidth
            onChange={(evt) =>
              setBody({
                ...body,
                email: evt.target.value,
              })
            }
          />
          <PasswordInput
            password={password}
            handlePasswordChange={handlePasswordChange}
          />
        </CardContent>
        <CardActions>
          <Button
            fullWidth
            type='submit'
            variant='contained'
            style={styles.button}
          >
            {header}
          </Button>
        </CardActions>
        {isLoginView && <CreateAccountButton />}
      </Card>
    </Box>
  );
};

const styles = {
  box: {
    display: 'flex',
    justifyContent: 'center',
    marginTop: MARGIN_TOP_CONTENT,
    alignItems: 'center',
  },
  card: {
    width: '23rem',
    borderRadius: '0.5rem',
    boxShadow: '0.5rem 1rem 1rem rgba(0, 0, 0, 0.1)',
  },
  header: {
    display: 'flex',
    justifyContent: 'center',
  },
  button: {
    marginBottom: '0.2rem',
    fontFamily: FONT_FAMILY,
    borderRadius: '0.5rem',
  },
};
