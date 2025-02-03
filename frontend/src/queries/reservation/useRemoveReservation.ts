import axios from 'axios';
import { useAppContextProvider } from '../../AppProvider';
import { SeverityOption } from '../../types/severity';
import { endpoints, routes } from '../../utils/routes';
import { StatusCode } from '../../types/statusCode';
import { CommonResponse, ErrorResponse } from '../../types/response';
import { CookieName, getCookieValueByName } from '../../utils/cookies';
import { ReservationData } from '../../types/reservation';
import { useHistory } from 'react-router';

interface UseRemoveReservationResult {
  remove: (id: string | undefined) => Promise<void>;
}

export const useRemoveReservation = (): UseRemoveReservationResult => {
  const { setSeverityText, setSeverity } = useAppContextProvider();
  const history = useHistory();

  const token = getCookieValueByName(CookieName.Token);

  const remove = async (id: string | undefined) => {
    axios
      .delete(endpoints.removeReservation(id ?? ''), {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then(({ data, status }: CommonResponse<ReservationData>) => {
        if (status !== StatusCode.OK) {
          setSeverity(SeverityOption.Error);
          setSeverityText('Internal Server Error');
          return;
        }

        setSeverity(SeverityOption.Success);
        setSeverityText(data.response.message);

        setTimeout(() => {
          history.push(routes.reservations);
        }, 500);
      })
      .catch(({ response }: ErrorResponse) => {
        setSeverity(SeverityOption.Error);
        setSeverityText(response.data.error);
      });
  };

  return { remove };
};
