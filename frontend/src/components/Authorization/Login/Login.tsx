import React, { useState } from 'react';
import { AuthBody } from '../../../types/authorization';
import { AuthorizationForm } from '../Form/AuthorizationForm';
import { CreateAccountButton } from '../Form/CreateAccountButton';
import { useGetCurrentUser } from '../../../queries/user/getCurrentUser';
import { UserAlreadyLogged } from '../../Error/UserAlreadyLogged';
import { CookieName, getCookieValueByName } from '../../../utils/cookies';
import { useLogin } from '../../../queries/user/useLogin';

export const Login: React.FC = () => {
  const [body, setBody] = useState<AuthBody>({
    email: '',
    password: '',
  });

  const { data } = useGetCurrentUser({
    skip: !getCookieValueByName(CookieName.UserId),
  });
  const { login } = useLogin();

  const handleSubmit = async (
    event: React.FormEvent<HTMLFormElement>
  ): Promise<void> => {
    event.preventDefault();
    login(body);
  };

  if (data?.id) {
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
