import React from 'react';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import { SettingIcon } from './SettingIcon';
import { NavbarElements } from './NavbarElements';
import { SmallNavbar } from './SmallNavbar';
import { routeBuilder } from '../../utils/routes';

export const Navbar: React.FC = () => {
  const pages = [
    routeBuilder.profile,
    routeBuilder.reservations
  ].map((page) => page.slice(1))

  return (
    <AppBar position="static">
      <Toolbar disableGutters>
        <SmallNavbar pages={pages} />
        <NavbarElements pages={pages} />
        <SettingIcon />
      </Toolbar>
    </AppBar>
  );
};