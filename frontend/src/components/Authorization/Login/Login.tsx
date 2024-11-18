import React, { useState } from 'react';
import { useAppContextProvider } from '../../../AppProvider';
import {
  AuthErrorResponse,
  AuthorizationBody,
  AuthResponse,
} from '../../../types/authorization';
import { AuthorizationForm } from '../Form/AuthorizationForm';
import axios from 'axios';
import { endpoints, routeBuilder } from '../../../utils/routes';
import { SeverityOption } from '../../../types/severity';
import { useAuth } from '../../../utils/authorization';

export const Login: React.FC = () => {
  const [body, setBody] = useState<AuthorizationBody>({
    email: '',
    password: '',
  });

  const { setSeverityText, setSeverity } = useAppContextProvider();

  const { signIn } = useAuth();

  const handleSubmit = async (
    event: React.FormEvent<HTMLFormElement>
  ): Promise<void> => {
    event.preventDefault();

    axios
      .post(endpoints.login, body)
      .then((res: AuthResponse) => {
        if (res.status === 200) {
          const { token, user } = res.data;

          signIn({
            url: routeBuilder.parking,
            token,
            role: user.role,
            userId: user.id,
          });
        }
      })
      .catch(({ response }: AuthErrorResponse) => {
        setSeverity(SeverityOption.Error);
        setSeverityText(response.data.error);
      });
  };

  return (
    <AuthorizationForm
      body={body}
      setBody={setBody}
      handleSubmit={handleSubmit}
      header='Login'
    />
  );
};
