import React from 'react';
import { AppNavigator } from './AppNavigator';
import { SnackbarAlert } from './components/SnackbarAlert/SnackbarAlert';

export const App: React.FC = () => {
  return (
    <>
      <AppNavigator />
      <SnackbarAlert />
    </>
  );
};
