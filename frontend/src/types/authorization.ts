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

export interface RegisteredUserData {
  message: string;
  user: User;
}

export interface AuthResponse<T> {
  status: number;
  data: T;
}

export interface ResetPasswordBody {
  currentPassword: string;
  newPassword: string;
  confirmNewPassword: string;
}

export type ResetPasswordData = RegisteredUserData;
