import React from 'react';
import { AppNavigator } from './AppNavigator';
import { SnackbarAlert } from './components/SnackbarAlert/SnackbarAlert';
import { Navbar } from './components/Navbar/Navbar';

export const App: React.FC = () => {
  return (
    <>
      <Navbar />
      <AppNavigator />
      <SnackbarAlert />
    </>
  );
};
