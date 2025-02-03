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

export interface UsersResponse {
  response: {
    users: Array<User>;
  };
}

export interface UserData {
  response: {
    message: string;
    reservation: User;
  };
}
