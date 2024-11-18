export interface User {
  id: number;
  email: string;
  name: string;
  surname: string;
  role: Role;
}

enum Role {
  Admin = 'admin',
  User = 'user',
}
