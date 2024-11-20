const HOST = 'http://localhost:8080';

interface Route {
  default: string;
  login: string;
  profile: string;
  register: string;
  reservations: string;
}

export const endpoints = {
  login: `${HOST}/sign-in`,
  register: `${HOST}/sign-up`,
  user: (id: string) => `${HOST}/users/${id}`,
};

export const routeBuilder: Route = {
  default: '/',
  login: '/login',
  profile: '/profile',
  register: '/register',
  reservations: '/reservations',
};
