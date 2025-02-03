export interface User {
  id: number;
  email: string;
  name: string;
  surname: string;
  role: Role;
}

export enum Role {
  Admin = 'admin',
  User = 'user',
}

export interface UserResponse {
  response: {
    users: Array<User>;
  };
}
