import React from 'react';
import { Route, Switch } from 'react-router';
import { routeBuilder } from './utils/routes';
import { Login } from './components/Authorization/Login/Login';

export const AppNavigator: React.FC = () => (
  <Switch>
    <Route exact path={routeBuilder.default} component={Login} />
    <Route path={routeBuilder.login} component={Login} />
  </Switch>
);
