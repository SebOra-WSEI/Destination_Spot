export interface PasswordRule {
  rule: boolean;
  label: string;
}

export const getPasswordValidationRules = (
  password: string,
  confirmPassword?: string
): Array<PasswordRule> => [
  {
    rule: /[A-Z]/.test(password),
    label: 'One uppercase latter',
  },
  {
    rule: /[0-9]/.test(password),
    label: 'One number',
  },
  {
    rule: /[^A-Za-z0-9]/.test(password),
    label: 'One special character',
  },
  {
    rule: password.length >= 8,
    label: 'Min. 8 characters',
  },
  {
    rule:
      password !== '' && confirmPassword !== '' && password === confirmPassword,
    label: 'Passwords match',
  },
];
