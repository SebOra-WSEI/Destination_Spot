import axios from 'axios';
import { useAppContextProvider } from '../../AppProvider';
import { SeverityOption } from '../../types/severity';
import { endpoints } from '../../utils/routes';
import { StatusCode } from '../../types/statusCode';
import { CommonResponse, ErrorResponse } from '../../types/response';
import { CookieName, getCookieValueByName } from '../../utils/cookies';
import { ReservationData, ReservationBody } from '../../types/reservation';
import { reloadPage } from '../../utils/reloadPage';

interface UseCreateReservationResult {
  reserve: (body: ReservationBody) => void;
}

export const useCreateReservation = (): UseCreateReservationResult => {
  const { setSeverityText, setSeverity } = useAppContextProvider();

  const token = getCookieValueByName(CookieName.Token);

  const reserve = (body: ReservationBody) => {
    axios
      .post(endpoints.reservations, body, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then(({ data, status }: CommonResponse<ReservationData>) => {
        if (status !== StatusCode.Created) {
          setSeverity(SeverityOption.Error);
          setSeverityText('Internal Server Error');
          return;
        }

        setSeverity(SeverityOption.Success);
        setSeverityText(data.response.message);
        reloadPage();
      })
      .catch(({ response }: ErrorResponse) => {
        setSeverity(SeverityOption.Error);
        setSeverityText(response.data.error);
      });
  };

  return { reserve };
};
