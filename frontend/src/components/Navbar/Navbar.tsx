import React from 'react';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import { SettingIcon } from './SettingIcon';
import { NavbarElements } from './NavbarElements';
import { SmallNavbar } from './SmallNavbar';
import { routeBuilder } from '../../utils/routes';
import { CookieName, getCookieValueByName } from '../../utils/cookies';
import { useLocation } from 'react-router';
import { Role } from '../../types/user';

export const Navbar: React.FC = () => {
  const location = useLocation();

  const role = getCookieValueByName(CookieName.Role);

  const pages = role
    ? loggedUserPages(role as Role)
    : getNotLoggedUserPages(location.pathname);

  return (
    <AppBar position='static'>
      <Toolbar disableGutters>
        <SmallNavbar pages={pages} />
        <NavbarElements pages={pages} />
        {!!role && <SettingIcon />}
      </Toolbar>
    </AppBar>
  );
};

function loggedUserPages(role: Role) {
  const loggedUserPages = mapPageName([
    routeBuilder.profile,
    routeBuilder.reservations,
    routeBuilder.addReservations
  ]);

  if (role === Role.User) {
    return loggedUserPages;
  }
  return [...loggedUserPages];
}

function getNotLoggedUserPages(path: string): Array<string> {
  if (path === routeBuilder.login) {
    return mapPageName([routeBuilder.register]);
  }
  return mapPageName([routeBuilder.login]);
}

function mapPageName(pages: Array<string>): Array<string> {
  return pages.map((page) => page.slice(1));
}
