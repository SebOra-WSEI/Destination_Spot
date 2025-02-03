const HOST = 'http://localhost:8080';

interface Route {
  default: string;
  login: string;
  profile: string;
  register: string;
  addReservations: string;
  reservations: string;
  reservationDetails: string;
}

export const endpoints = {
  login: `${HOST}/sign-in`,
  register: `${HOST}/sign-up`,
  reservation: (id: string) => `${HOST}/reservations/${id}`,
  reservations: `${HOST}/reservations`,
  spots: `${HOST}/spots`,
  user: (id: string) => `${HOST}/users/${id}`,
  removeReservation: (id: string) => `${HOST}/reservations/${id}`,
  resetPassword: (id: string) => `${HOST}/reset-password/${id}`,
};

export const routes: Route = {
  default: '/',
  login: '/login',
  profile: '/profile',
  register: '/register',
  addReservations: '/add-reservations',
  reservations: '/reservations',
  reservationDetails: '/reservations/:id',
};

export const routeBuilder = {
  reservationDetails: (id: string) =>
    routes.reservationDetails.replace(':id', id),
};
