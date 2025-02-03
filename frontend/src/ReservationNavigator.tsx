import React from "react";
import { Route, Switch } from "react-router";
import { routes } from "./utils/routes";
import { ReservationsView } from "./components/Reservations/ReservationsView";
import { AddReservationView } from "./components/Reservations/AddReservation/AddReservationView";
import { ReservationDetailsView } from "./components/Reservations/Details/ReservationDetailsView";

export const ReservationNavigator: React.FC = () => (
  <Switch>
    <Route path={routes.reservationDetails} component={ReservationDetailsView} />
    <Route path={routes.reservations} component={ReservationsView} />
    <Route path={routes.createReservation} component={AddReservationView} />
  </Switch>
)