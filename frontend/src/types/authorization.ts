import { User } from './user';

export interface AuthBody {
  email: string;
  password: string;
  confirmPassword?: string;
}

export interface LoggedUserData {
  token: string;
  user: User;
}

export interface RegisteredUserResponse {
  message: string;
  user: User;
}

export interface ResetPasswordBody {
  currentPassword?: string;
  newPassword: string;
  confirmNewPassword: string;
}

export type ResetPasswordResponse = RegisteredUserResponse;
