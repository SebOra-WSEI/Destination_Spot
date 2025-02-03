const HOST = 'http://localhost:8080';

interface Route {
  default: string;
  login: string;
  profile: string;
  register: string;
  createReservation: string;
  reservations: string;
  reservationDetails: string;
  spots: string;
  users: string;
  userDetails: string;
}

export const endpoints = {
  // Authorization
  accessControl: (id: string) => `${HOST}/access-control/${id}`,
  login: `${HOST}/sign-in`,
  register: `${HOST}/sign-up`,
  resetPassword: (id: string) => `${HOST}/reset-password/${id}`,
  // Reservation
  reservation: (id: string) => `${HOST}/reservations/${id}`,
  reservations: `${HOST}/reservations`,
  // Spot
  spot: (id: string) => `${HOST}/spots/${id}`,
  spots: `${HOST}/spots`,
  // User
  user: (id: string) => `${HOST}/users/${id}`,
  users: `${HOST}/users`,
};

export const routes: Route = {
  default: '/',
  login: '/login',
  profile: '/profile',
  register: '/register',
  createReservation: '/add-reservation',
  reservations: '/reservations',
  reservationDetails: '/reservations/:id',
  spots: '/spots',
  users: '/users',
  userDetails: '/users/:id',
};

export const routeBuilder = {
  reservationDetails: (id: string) =>
    routes.reservationDetails.replace(':id', id),
  userDetails: (id: string) => routes.userDetails.replace(':id', id),
};
