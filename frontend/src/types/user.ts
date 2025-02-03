import { Role } from '../utils/consts';

export interface User {
  id: number;
  email: string;
  name: string;
  surname: string;
  role: Role;
}

export interface UsersResponse {
  response: {
    users: Array<User>;
  };
}

export interface UserResponse {
  response: {
    message: string;
    user: User;
  };
}
