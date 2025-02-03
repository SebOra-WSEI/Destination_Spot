import axios from 'axios';
import { useAppContextProvider } from '../../AppProvider';
import { endpoints } from '../../utils/routes';
import { CommonResponse, ErrorResponse } from '../../types/response';
import { CookieName, getCookieValueByName } from '../../utils/cookies';
import { ReservationResponse, ReservationBody } from '../../types/reservation';
import { SeverityOption, StatusCode } from '../../utils/consts';
import { reloadPage } from '../../utils/reloadPage';

interface UseCreateReservationResult {
  reserve: (body: ReservationBody) => Promise<void>;
}

export const useCreateReservation = (): UseCreateReservationResult => {
  const { setSeverityText, setSeverity } = useAppContextProvider();

  const token = getCookieValueByName(CookieName.Token);

  const reserve = async (body: ReservationBody) => {
    axios
      .post(endpoints.reservations, body, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then(({ data, status }: CommonResponse<ReservationResponse>) => {
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
