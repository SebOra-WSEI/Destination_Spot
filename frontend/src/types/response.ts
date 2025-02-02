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

export interface CommonResponse<T> {
  status: number;
  data: T;
}
