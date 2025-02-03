import React from "react";
import { Route, Switch } from "react-router";
import { routeBuilder } from "./utils/routes";
import { ReservationsView } from "./components/Reservations/ReservationsView";
import { AddReservationView } from "./components/Reservations/AddReservation/AddReservationView";
import { ReservationDetailsView } from "./components/Reservations/Details/ReservationDetailsView";

export const ReservationNavigator: React.FC = () => (
  <Switch>
    <Route path={routeBuilder.reservationDetails} component={ReservationDetailsView} />
    <Route path={routeBuilder.reservations} component={ReservationsView} />
    <Route path={routeBuilder.addReservations} component={AddReservationView} />
  </Switch>
)