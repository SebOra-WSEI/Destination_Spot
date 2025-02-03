import axios from 'axios';
import { useAppContextProvider } from '../../AppProvider';
import { endpoints } from '../../utils/routes';
import { CommonResponse, ErrorResponse } from '../../types/response';
import { CookieName, getCookieValueByName } from '../../utils/cookies';
import { SpotResponse } from '../../types/spot';
import { SeverityOption, StatusCode } from '../../utils/consts';
import { reloadPage } from '../../utils/reloadPage';

interface UseRemoveSpotResult {
  remove: (id: number) => Promise<void>;
}

export const useRemoveSpot = (): UseRemoveSpotResult => {
  const { setSeverityText, setSeverity } = useAppContextProvider();

  const token = getCookieValueByName(CookieName.Token);

  const remove = async (id: number) => {
    axios
      .delete(endpoints.spot(String(id)), {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then(({ data, status }: CommonResponse<SpotResponse>) => {
        if (status !== StatusCode.OK) {
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

  return { remove };
};
