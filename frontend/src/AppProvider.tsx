import React, {
  PropsWithChildren,
  createContext,
  useContext,
  useState,
} from 'react';
import { SeverityOption } from './types/severity';

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
  const [severityText, setSeverityText] = useState<string>('');
  const [severity, setSeverity] = useState<SeverityOption>(
    SeverityOption.Error
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
