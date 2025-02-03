import axios from 'axios';
import { useAppContextProvider } from '../../AppProvider';
import { endpoints, routes } from '../../utils/routes';
import { AuthBody, RegisteredUserResponse } from '../../types/authorization';
import { CommonResponse, ErrorResponse } from '../../types/response';
import { useHistory } from 'react-router';
import { SeverityOption, StatusCode } from '../../utils/consts';

interface UseRegisterResult {
  register: (body: AuthBody) => Promise<void>;
}

export const useRegister = (): UseRegisterResult => {
  const { setSeverityText, setSeverity } = useAppContextProvider();

  const history = useHistory();

  const register = async (body: AuthBody) => {
    axios
      .post(endpoints.register, body)
      .then(({ data, status }: CommonResponse<RegisteredUserResponse>) => {
        if (status !== StatusCode.Created) {
          setSeverity(SeverityOption.Error);
          setSeverityText('Internal Server Error');
          return;
        }

        setSeverity(SeverityOption.Success);
        setSeverityText(data.message);

        setTimeout(() => {
          history.push(routes.login);
        }, 500);
      })
      .catch(({ response }: ErrorResponse) => {
        setSeverity(SeverityOption.Error);
        setSeverityText(response.data.error);
      });
  };

  return { register };
};
