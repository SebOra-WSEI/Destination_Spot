const HOST = 'http://localhost:8080';

interface Route {
  default: string;
  login: string;
  profile: string;
  register: string;
  reservations: string;
  addReservations: string;
}

export const endpoints = {
  login: `${HOST}/sign-in`,
  register: `${HOST}/sign-up`,
  reservations: `${HOST}/reservations`,
  spots: `${HOST}/spots`,
  user: (id: string) => `${HOST}/users/${id}`,
  removeReservation: (id: string) => `${HOST}/reservations/${id}`,
  resetPassword: (id: string) => `${HOST}/reset-password/${id}`,
};

export const routeBuilder: Route = {
  default: '/',
  login: '/login',
  profile: '/profile',
  register: '/register',
  reservations: '/reservations',
  addReservations: '/add-reservations',
};
