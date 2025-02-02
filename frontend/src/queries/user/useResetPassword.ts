import axios from 'axios';
import { useAppContextProvider } from '../../AppProvider';
import {
  ResetPasswordBody,
  ResetPasswordData,
} from '../../types/authorization';
import { endpoints } from '../../utils/routes';
import { SeverityOption } from '../../types/severity';
import { StatusCode } from '../../types/statusCode';
import { CookieName, getCookieValueByName } from '../../utils/cookies';
import { CommonResponse, ErrorResponse } from '../../types/response';

interface UseResetPasswordResult {
  resetPassword: (body: ResetPasswordBody) => void;
}

export const useResetPassword = (
  onCloseModal: () => void
): UseResetPasswordResult => {
  const { setSeverityText, setSeverity } = useAppContextProvider();

  const token = getCookieValueByName(CookieName.Token);

  const resetPassword = async (body: ResetPasswordBody) => {
    axios
      .put(endpoints.resetPassword('1'), body, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then(({ data, status }: CommonResponse<ResetPasswordData>) => {
        if (status !== StatusCode.OK) {
          setSeverity(SeverityOption.Error);
          setSeverityText('Internal Server Error');
          return;
        }

        setSeverity(SeverityOption.Success);
        setSeverityText(data.message);
        onCloseModal();
      })
      .catch((err: ErrorResponse) => {
        setSeverity(SeverityOption.Error);
        setSeverityText(err.response.data.error);
      });
  };

  return { resetPassword };
};
