import React from 'react';
import { Route, Switch } from 'react-router';
import { routeBuilder } from './utils/routes';
import { Login } from './components/Authorization/Login/Login';
import { Register } from './components/Authorization/Register/Register';
import { UserView } from './components/UserView/UserView';

export const AppNavigator: React.FC = () => (
  <Switch>
    <Route exact path={routeBuilder.default} component={Login} />
    <Route path={routeBuilder.login} component={Login} />
    <Route path={routeBuilder.register} component={Register} />
    <Route path={routeBuilder.profile} component={UserView} />
  </Switch>
);