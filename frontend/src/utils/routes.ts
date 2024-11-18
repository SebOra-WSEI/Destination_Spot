const HOST = 'http://localhost:8080';

interface Route {
  default: string;
  login: string;
  parking: string;
  register: string;
}

export const endpoints = {
  login: `${HOST}/sign-in`,
  register: `${HOST}/sign-up`,
};

export const routeBuilder: Route = {
  default: '/',
  login: '/login',
  parking: '/parking',
  register: '/register',
};
