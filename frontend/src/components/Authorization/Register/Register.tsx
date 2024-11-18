import React, { useState } from 'react';
import { AuthorizationForm } from '../Form/AuthorizationForm';
import { AuthBody, AuthResponse, RegisteredUserData } from '../../../types/authorization';
import { RegisterFormFooter } from './RegisterFormFooter';
import { useAppContextProvider } from '../../../AppProvider';
import axios from 'axios';
import { endpoints, routeBuilder } from '../../../utils/routes';
import { SeverityOption } from '../../../types/severity';
import { ErrorResponse } from '../../../types/response';
import { StatusCode } from '../../../types/statusCode';
import { useHistory } from 'react-router';

export const Register: React.FC = () => {
  const history = useHistory();

  const [body, setBody] = useState<AuthBody>({
    email: '',
    password: '',
    confirmPassword: '',
  });

  const { setSeverityText, setSeverity } = useAppContextProvider();

  const handleSubmit = async (
    event: React.FormEvent<HTMLFormElement>
  ): Promise<void> => {
    event.preventDefault();

    axios.post(endpoints.register, body)
      .then(({ data, status }: AuthResponse<RegisteredUserData>) => {
        if (status == StatusCode.Created) {
          setSeverity(SeverityOption.Success)
          setSeverityText(data.message)

          setTimeout(() => {
            history.push(routeBuilder.login);
          }, 500);
        }
      })
      .catch(({ response }: ErrorResponse) => {
        setSeverity(SeverityOption.Error);
        setSeverityText(response.data.error)
      })
  };

  return (
    <AuthorizationForm
      body={body}
      setBody={setBody}
      handleSubmit={handleSubmit}
      header='Register'
      footer={<RegisterFormFooter />}
    />
  )
}