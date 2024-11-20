import React, { useState } from 'react';
import { AuthorizationForm } from '../Form/AuthorizationForm';
import { AuthBody } from '../../../types/authorization';
import { RegisterFormFooter } from './RegisterFormFooter';
import { useRegister } from '../../../queries/user/useRegister';

export const Register: React.FC = () => {
  const [body, setBody] = useState<AuthBody>({
    email: '',
    password: '',
    confirmPassword: '',
  });

  const { register } = useRegister();

  const handleSubmit = async (
    event: React.FormEvent<HTMLFormElement>
  ): Promise<void> => {
    event.preventDefault();
    register(body);
  };

  return (
    <AuthorizationForm
      body={body}
      setBody={setBody}
      handleSubmit={handleSubmit}
      header='Register'
      footer={<RegisterFormFooter />}
    />
  );
};
