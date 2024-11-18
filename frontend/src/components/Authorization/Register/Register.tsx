import React, { useState } from 'react';
import { AuthorizationForm } from '../Form/AuthorizationForm';
import { AuthBody } from '../../../types/authorization';
import { RegisterFormFooter } from './RegisterFormFooter';

export const Register: React.FC = () => {
  const [body, setBody] = useState<AuthBody>({
    email: '',
    password: '',
    confirmPassword: '',
  });

  const handleSubmit = async (
    event: React.FormEvent<HTMLFormElement>
  ): Promise<void> => {
    event.preventDefault();
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