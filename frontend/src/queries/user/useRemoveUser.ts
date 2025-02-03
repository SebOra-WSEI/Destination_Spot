import axios from 'axios';
import { useAppContextProvider } from '../../AppProvider';
import { endpoints } from '../../utils/routes';
import { CommonResponse, ErrorResponse } from '../../types/response';
import { CookieName, getCookieValueByName } from '../../utils/cookies';
import { UserResponse } from '../../types/user';
import { SeverityOption, StatusCode } from '../../utils/consts';
import { reloadPage } from '../../utils/reloadPage';

interface UseRemoveUserResult {
  remove: (id: string | undefined) => Promise<void>;
}

export const useRemoveUser = (onSuccess?: () => void): UseRemoveUserResult => {
  const { setSeverityText, setSeverity } = useAppContextProvider();

  const token = getCookieValueByName(CookieName.Token);

  const remove = async (id: string | undefined) => {
    axios
      .delete(endpoints.user(id ?? ''), {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then(({ data, status }: CommonResponse<UserResponse>) => {
        if (status !== StatusCode.OK) {
          setSeverity(SeverityOption.Error);
          setSeverityText('Internal Server Error');
          return;
        }

        setSeverity(SeverityOption.Success);
        setSeverityText(data.response.message);

        onSuccess
          ? setTimeout(() => {
              onSuccess();
            }, 500)
          : reloadPage();
      })
      .catch(({ response }: ErrorResponse) => {
        setSeverity(SeverityOption.Error);
        setSeverityText(response.data.error);
      });
  };

  return { remove };
};
