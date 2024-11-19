import { User } from './user';

interface ErrorData {
  error: string;
}

export interface ErrorResponse {
  response: {
    status: number;
    data: ErrorData;
  };
}

export interface UserResponse {
  response: {
    user: User;
  };
}
