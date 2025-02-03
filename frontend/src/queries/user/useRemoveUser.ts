import axios from 'axios';
import { useAppContextProvider } from '../../AppProvider';
import { SeverityOption } from '../../types/severity';
import { endpoints } from '../../utils/routes';
import { StatusCode } from '../../types/statusCode';
import { CommonResponse, ErrorResponse } from '../../types/response';
import { CookieName, getCookieValueByName } from '../../utils/cookies';
import { UserData } from '../../types/user';

interface UseRemoveUserResult {
  remove: (id: string | undefined) => Promise<void>;
}

export const useRemoveUser = (): UseRemoveUserResult => {
  const { setSeverityText, setSeverity } = useAppContextProvider();

  const token = getCookieValueByName(CookieName.Token);

  const remove = async (id: string | undefined) => {
    axios
      .delete(endpoints.user(id ?? ''), {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then(({ data, status }: CommonResponse<UserData>) => {
        if (status !== StatusCode.OK) {
          setSeverity(SeverityOption.Error);
          setSeverityText('Internal Server Error');
          return;
        }

        setSeverity(SeverityOption.Success);
        setSeverityText(data.response.message);

        setTimeout(() => {
          window.location.reload();
        }, 500);
      })
      .catch(({ response }: ErrorResponse) => {
        setSeverity(SeverityOption.Error);
        setSeverityText(response.data.error);
      });
  };

  return { remove };
};
