import { jwtDecode } from 'jwt-decode';
import { TOKEN_KEY } from './consts';
import { CookieName, eraseCookie, setCookie } from './cookies';
import { routeBuilder } from './routes';
import { useHistory } from 'react-router';

interface SignInArgs {
  url: string;
  token: string | undefined;
  role: string | undefined;
  userId: number | undefined;
}

interface UseAuthResult {
  signIn: ({ url, token, role, userId }: SignInArgs) => void;
  signOut: () => void;
}

export const useAuth = (): UseAuthResult => {
  const history = useHistory();

  const signIn = ({ url, token, role, userId }: SignInArgs): void => {
    if (!token || !role || !userId) {
      return;
    }

    const tokenExpireEpoch = jwtDecode(token)?.exp ?? 0;
    const expireDate = new Date(tokenExpireEpoch).toString();

    window.localStorage.setItem(TOKEN_KEY, token);

    setCookie(CookieName.Role, role);
    setCookie(CookieName.UserId, String(userId));
    setCookie(CookieName.Token, token);
    setCookie(CookieName.Expires, expireDate);

    history.push(url);
  };

  const signOut = (): void => {
    window.localStorage.removeItem(TOKEN_KEY);
    Object.values(CookieName).forEach((val) => eraseCookie(val));
    window.location.assign(routeBuilder.login);
  };

  return {
    signIn,
    signOut,
  };
};
