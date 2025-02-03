import {
  Box,
  Button,
  Card,
  CardActions,
  CardContent,
  TextField,
} from '@mui/material';
import React from 'react';
import {
  BUTTON_RADIUS,
  FONT_FAMILY,
  MARGIN_TOP_CONTENT,
} from '../../../utils/consts';
import { AuthBody } from '../../../types/authorization';
import { PasswordInput } from './PasswordInput';
import { useLocation } from 'react-router';
import { routes } from '../../../utils/routes';
import { PasswordCheckList } from '../PasswordValidation/PasswordCheckList';
import { getPasswordValidationRules } from '../../../utils/getPasswordValidationRules';

interface AuthorizationFormProps {
  body: AuthBody;
  setBody: (body: AuthBody) => void;
  handleSubmit: (event: React.FormEvent<HTMLFormElement>) => Promise<void>;
  header: string;
  footer: React.ReactNode;
}

export const AuthorizationForm: React.FC<AuthorizationFormProps> = ({
  body,
  setBody,
  handleSubmit,
  header,
  footer,
}) => {
  const { pathname } = useLocation();
  const { email, password, confirmPassword } = body;

  const isRegisterView = pathname === routes.register;

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
          {isRegisterView && (
            <TextField
              required
              label='Confirm Password'
              variant='standard'
              type='password'
              autoComplete='password'
              value={confirmPassword}
              fullWidth
              onChange={(evt) =>
                setBody({
                  ...body,
                  confirmPassword: evt.target.value,
                })
              }
            />
          )}
          {isRegisterView && Boolean(password) && (
            <PasswordCheckList
              rules={getPasswordValidationRules(password, confirmPassword)}
            />
          )}
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
        {footer}
      </Card>
    </Box>
  );
};

const styles = {
  box: {
    display: 'flex',
    justifyContent: 'center',
    marginTop: MARGIN_TOP_CONTENT,
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
    borderRadius: BUTTON_RADIUS,
  },
};
