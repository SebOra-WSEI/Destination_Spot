import { TOKEN_KEY } from './consts';
import { CookieName, eraseCookie } from './cookies';
import { routeBuilder } from './routes';

export const signOut = (): void => {
  window.localStorage.removeItem(TOKEN_KEY);
  Object.values(CookieName).forEach((val) => eraseCookie(val));
  window.location.replace(routeBuilder.login);
};
