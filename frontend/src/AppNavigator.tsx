import React from 'react';
import { Route, Switch } from 'react-router';
import { routes } from './utils/routes';
import { Login } from './components/Authorization/Login/Login';
import { Register } from './components/Authorization/Register/Register';
import { UserView } from './components/UserView/UserView';
import { PageNotFound } from './components/Error/PageNotFound';
import { ReservationNavigator } from './ReservationNavigator';

export const AppNavigator: React.FC = () => (
  <Switch>
    <Route exact path={routes.default} component={Login} />
    <Route path={routes.login} component={Login} />
    <Route path={routes.register} component={Register} />
    <Route path={routes.profile} component={UserView} />
    <Route
      path={[
        routes.reservations,
        routes.addReservations
      ]}
      component={ReservationNavigator}
    />
    <Route component={PageNotFound} />
  </Switch>
);
