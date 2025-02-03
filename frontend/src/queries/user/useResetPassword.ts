import axios from 'axios';
import { useAppContextProvider } from '../../AppProvider';
import {
  ResetPasswordBody,
  ResetPasswordResponse,
} from '../../types/authorization';
import { endpoints } from '../../utils/routes';
import { CookieName, getCookieValueByName } from '../../utils/cookies';
import { CommonResponse, ErrorResponse } from '../../types/response';
import { useParams } from 'react-router';
import { SeverityOption, StatusCode } from '../../utils/consts';

interface UseResetPasswordResult {
  resetPassword: (body: ResetPasswordBody) => Promise<void>;
}

export const useResetPassword = (
  osSuccess: () => void
): UseResetPasswordResult => {
  const { setSeverityText, setSeverity } = useAppContextProvider();
  const { id: idParams } = useParams<{ id: string }>();

  const userId = getCookieValueByName(CookieName.UserId);
  const token = getCookieValueByName(CookieName.Token);

  const endpoint = idParams
    ? endpoints.accessControl(idParams)
    : endpoints.resetPassword(userId ?? '');

  const resetPassword = async (body: ResetPasswordBody) => {
    axios
      .put(endpoint, body, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then(({ data, status }: CommonResponse<ResetPasswordResponse>) => {
        if (status !== StatusCode.OK) {
          setSeverity(SeverityOption.Error);
          setSeverityText('Internal Server Error');
          return;
        }

        setSeverity(SeverityOption.Success);
        setSeverityText(data.message);
        osSuccess();
      })
      .catch((err: ErrorResponse) => {
        setSeverity(SeverityOption.Error);
        setSeverityText(err.response.data.error);
      });
  };

  return { resetPassword };
};
