import axios from 'axios';
import { useAppContextProvider } from '../../AppProvider';
import { SeverityOption } from '../../types/severity';
import { endpoints, routeBuilder } from '../../utils/routes';
import { AuthBody, RegisteredUserData } from '../../types/authorization';
import { StatusCode } from '../../types/statusCode';
import { CommonResponse, ErrorResponse } from '../../types/response';
import { useHistory } from 'react-router';

interface UseRegisterResult {
  register: (body: AuthBody) => Promise<void>;
}

export const useRegister = (): UseRegisterResult => {
  const { setSeverityText, setSeverity } = useAppContextProvider();

  const history = useHistory();

  const register = async (body: AuthBody) => {
    axios
      .post(endpoints.register, body)
      .then(({ data, status }: CommonResponse<RegisteredUserData>) => {
        if (status !== StatusCode.Created) {
          setSeverity(SeverityOption.Error);
          setSeverityText('Internal Server Error');
          return;
        }

        setSeverity(SeverityOption.Success);
        setSeverityText(data.message);

        setTimeout(() => {
          history.push(routeBuilder.login);
        }, 500);
      })
      .catch(({ response }: ErrorResponse) => {
        setSeverity(SeverityOption.Error);
        setSeverityText(response.data.error);
      });
  };

  return { register };
};
