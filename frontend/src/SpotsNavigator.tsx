import React from 'react';
import { Route, Switch } from 'react-router';
import { routes } from './utils/routes';
import { PermissionDenied } from './components/Error/PermissionDenied';
import { CookieName, getCookieValueByName } from './utils/cookies';
import { Role } from './types/user';
import { SpotsList } from './components/Spots/SpotsList';

export const SpotsNavigator: React.FC = () => {
  const role = getCookieValueByName(CookieName.Role);

  return (
    <Switch>
      {role === Role.Admin && (
        <Route path={routes.spots} component={SpotsList} />
      )}
      <Route component={PermissionDenied} />
    </Switch>
  );
}