import React from 'react';
import { SnackbarAlert } from './components/SnackbarAlert/SnackbarAlert';
import { Navbar } from './components/Navbar/Navbar';
import { AppNavigator } from './navigators/AppNavigator';

export const App: React.FC = () => (
  <>
    <Navbar />
    <AppNavigator />
    <SnackbarAlert />
  </>
);
