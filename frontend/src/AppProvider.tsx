import React, {
  PropsWithChildren,
  createContext,
  useContext,
  useState,
} from 'react';
import { SeverityOption } from './utils/consts';

interface AppContextState {
  severityText: string;
  severity: SeverityOption;
  setSeverityText: (value: string) => void;
  setSeverity: (value: SeverityOption) => void;
}

const defaultAppState = {
  severityText: '',
  severity: SeverityOption.Error,
  setSeverityText: (_: string) => {},
  setSeverity: (_: SeverityOption) => {},
};

const AppContext = createContext(defaultAppState);

export const useAppContextProvider = (): AppContextState =>
  useContext(AppContext);

export const AppContextProvider: React.FC<PropsWithChildren> = ({
  children,
}) => {
  const [severityText, setSeverityText] = useState<string>(
    defaultAppState.severityText
  );
  const [severity, setSeverity] = useState<SeverityOption>(
    defaultAppState.severity
  );

  return (
    <AppContext.Provider
      value={{
        setSeverity,
        setSeverityText,
        severity,
        severityText,
      }}
    >
      {children}
    </AppContext.Provider>
  );
};
