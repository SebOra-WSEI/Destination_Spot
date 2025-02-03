import axios from 'axios';
import { useAppContextProvider } from '../../AppProvider';
import { SeverityOption } from '../../types/severity';
import { endpoints } from '../../utils/routes';
import { StatusCode } from '../../types/statusCode';
import { CommonResponse, ErrorResponse } from '../../types/response';
import { CookieName, getCookieValueByName } from '../../utils/cookies';
import { SpotData } from '../../types/spot';

interface UseCreateSpotResult {
  create: (location: number) => Promise<void>;
}

export const useCreateSpot = (): UseCreateSpotResult => {
  const { setSeverityText, setSeverity } = useAppContextProvider();

  const token = getCookieValueByName(CookieName.Token);

  const create = async (location: number) => {
    axios
      .post(
        endpoints.spots,
        { location },
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      )
      .then(({ data, status }: CommonResponse<SpotData>) => {
        if (status !== StatusCode.Created) {
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

  return { create };
};
