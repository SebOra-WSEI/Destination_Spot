import { Visibility, VisibilityOff } from '@mui/icons-material';
import {
  FormControl,
  IconButton,
  Input,
  InputAdornment,
  InputLabel,
} from '@mui/material';
import React, { useState } from 'react';

interface PasswordInputProps {
  password: string;
  handlePasswordChange: (
    evt: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => void;
  label?: string;
}

export const PasswordInput: React.FC<PasswordInputProps> = ({
  password,
  handlePasswordChange,
  label,
}) => {
  const [showPassword, setShowPassword] = useState<boolean>(false);

  const handleClickShowPassword = (): void => setShowPassword(!showPassword);

  return (
    <FormControl fullWidth variant='standard'>
      <InputLabel htmlFor={label ?? 'password'}>
        {label ?? 'Password'} *
      </InputLabel>
      <Input
        autoComplete='password'
        value={password}
        id={label ?? 'password'}
        type={showPassword ? 'text' : 'password'}
        onChange={handlePasswordChange}
        endAdornment={
          <InputAdornment position='end'>
            <IconButton onClick={handleClickShowPassword}>
              {showPassword ? <VisibilityOff /> : <Visibility />}
            </IconButton>
          </InputAdornment>
        }
      />
    </FormControl>
  );
};
