import axios from 'axios';
import { useAppContextProvider } from '../../AppProvider';
import { endpoints } from '../../utils/routes';
import { CommonResponse, ErrorResponse } from '../../types/response';
import { CookieName, getCookieValueByName } from '../../utils/cookies';
import { SpotResponse } from '../../types/spot';
import { SeverityOption, StatusCode } from '../../utils/consts';
import { reloadPage } from '../../utils/reloadPage';

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
      .then(({ data, status }: CommonResponse<SpotResponse>) => {
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

  return { create };
};
