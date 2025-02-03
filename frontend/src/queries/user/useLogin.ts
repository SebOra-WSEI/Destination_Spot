import axios from 'axios';
import { useAppContextProvider } from '../../AppProvider';
import { endpoints, routes } from '../../utils/routes';
import { AuthBody, LoggedUserData } from '../../types/authorization';
import { CommonResponse, ErrorResponse } from '../../types/response';
import { jwtDecode } from 'jwt-decode';
import { SeverityOption, StatusCode, TOKEN_KEY } from '../../utils/consts';
import { CookieName, setCookie } from '../../utils/cookies';

interface UseLoginResult {
  login: (body: AuthBody) => Promise<void>;
}

export const useLogin = (): UseLoginResult => {
  const { setSeverityText, setSeverity } = useAppContextProvider();

  const login = async (body: AuthBody) => {
    axios
      .post(endpoints.login, body)
      .then(({ data, status }: CommonResponse<LoggedUserData>) => {
        if (status !== StatusCode.OK) {
          setSeverity(SeverityOption.Error);
          setSeverityText('Internal Server Error');
          return;
        }

        const {
          token,
          user: { role, id },
        } = data;

        const tokenExpireEpoch = jwtDecode(token)?.exp ?? 0;
        const expireDate = new Date(tokenExpireEpoch).toString();

        window.localStorage.setItem(TOKEN_KEY, token);

        setCookie(CookieName.Role, role);
        setCookie(CookieName.UserId, String(id));
        setCookie(CookieName.Token, token);
        setCookie(CookieName.Expires, expireDate);

        window.location.replace(routes.profile);
      })
      .catch(({ response }: ErrorResponse) => {
        setSeverity(SeverityOption.Error);
        setSeverityText(response.data.error);
      });
  };

  return { login };
};
