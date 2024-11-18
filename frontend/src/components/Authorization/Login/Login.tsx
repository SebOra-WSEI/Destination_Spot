import React, { useState } from 'react';
import { useAppContextProvider } from '../../../AppProvider';
import {
  AuthBody,
  AuthResponse,
  LoggedUserData,
} from '../../../types/authorization';
import { AuthorizationForm } from '../Form/AuthorizationForm';
import axios from 'axios';
import { endpoints, routeBuilder } from '../../../utils/routes';
import { SeverityOption } from '../../../types/severity';
import { useAuth } from '../../../utils/authorization';
import { CreateAccountButton } from '../Form/CreateAccountButton';
import { ErrorResponse } from '../../../types/response';
import { StatusCode } from '../../../types/statusCode';

export const Login: React.FC = () => {
  const [body, setBody] = useState<AuthBody>({
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
      .then(({ data, status }: AuthResponse<LoggedUserData>) => {
        if (status === StatusCode.OK) {
          const { token, user } = data;

          signIn({
            url: routeBuilder.parking,
            token,
            role: user.role,
            userId: user.id,
          });
        }
      })
      .catch(({ response }: ErrorResponse) => {
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
      footer={<CreateAccountButton />}
    />
  );
};
