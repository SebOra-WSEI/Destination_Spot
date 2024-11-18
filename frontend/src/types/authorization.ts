import { User } from './user';

export interface AuthorizationBody {
  email: string;
  password: string;
}

interface ErrorData {
  error: string;
}

export interface AuthErrorResponse {
  response: {
    status: number;
    data: ErrorData;
  };
}

interface LoggedUserData {
  token: string;
  user: User;
}

export interface AuthResponse {
  status: number;
  data: LoggedUserData;
}
