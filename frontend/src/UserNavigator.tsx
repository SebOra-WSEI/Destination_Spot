import React from "react";
import { Route, Switch } from "react-router";
import { routes } from "./utils/routes";
import { UserView } from "./components/User/UserView";
import { UsersList } from "./components/User/UsersList";
import { CookieName, getCookieValueByName } from "./utils/cookies";
import { Role } from "./types/user";
import { PermissionDenied } from "./components/Error/PermissionDenied";

export const UserNavigator: React.FC = () => {
  const role = getCookieValueByName(CookieName.Role);

  return (
    <Switch>
      {role === Role.Admin && (
        <Route path={routes.userDetails} component={() => <>ccc</>} />
      )}
      {role === Role.Admin && (
        <Route path={routes.users} component={UsersList} />
      )}
      <Route path={routes.profile} component={UserView} />
      <Route component={PermissionDenied} />
    </Switch>
  )
}