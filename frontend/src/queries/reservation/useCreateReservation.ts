import axios from 'axios';
import { useAppContextProvider } from '../../AppProvider';
import { SeverityOption } from '../../types/severity';
import { endpoints } from '../../utils/routes';
import { AuthResponse } from '../../types/authorization';
import { StatusCode } from '../../types/statusCode';
import { ErrorResponse } from '../../types/response';
import { CookieName, getCookieValueByName } from '../../utils/cookies';
import {
  CreatedReservationData,
  ReservationBody,
} from '../../types/reservation';

interface UseCreateReservationResult {
  reserve: (body: ReservationBody) => void;
}

export const useCreateReservation = (
  onCloseModal: () => void
): UseCreateReservationResult => {
  const { setSeverityText, setSeverity } = useAppContextProvider();

  const token = getCookieValueByName(CookieName.Token);

  const reserve = (body: ReservationBody) => {
    axios
      .post(endpoints.reservations, body, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then(({ data, status }: AuthResponse<CreatedReservationData>) => {
        if (status !== StatusCode.Created) {
          setSeverity(SeverityOption.Error);
          setSeverityText('Internal Server Error');
          return;
        }

        setSeverity(SeverityOption.Success);
        setSeverityText(data.response.message);
        onCloseModal();
      })
      .catch(({ response }: ErrorResponse) => {
        setSeverity(SeverityOption.Error);
        setSeverityText(response.data.error);
      });
  };

  return { reserve };
};
