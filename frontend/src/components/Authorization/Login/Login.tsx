import React, { useState } from 'react';
import { AuthBody } from '../../../types/authorization';
import { AuthorizationForm } from '../Form/AuthorizationForm';
import { CreateAccountButton } from '../Form/CreateAccountButton';
import { UserAlreadyLogged } from '../../Error/UserAlreadyLogged';
import { CookieName, getCookieValueByName } from '../../../utils/cookies';
import { useLogin } from '../../../queries/user/useLogin';

export const Login: React.FC = () => {
  const [body, setBody] = useState<AuthBody>({
    email: '',
    password: '',
  });

  const { login } = useLogin();

  const userId = getCookieValueByName(CookieName.UserId);

  const handleSubmit = async (
    event: React.FormEvent<HTMLFormElement>
  ): Promise<void> => {
    event.preventDefault();
    await login(body);
  };

  if (userId) {
    return <UserAlreadyLogged />;
  }

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
