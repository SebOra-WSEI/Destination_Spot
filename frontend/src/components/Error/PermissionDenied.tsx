import React from "react";
import { ErrorCard } from "./ErrorCard";
import { routes } from "../../utils/routes";

export const PermissionDenied: React.FC = () => (
  <ErrorCard isErrorCard text='Permission denied' link={routes.profile} />
)