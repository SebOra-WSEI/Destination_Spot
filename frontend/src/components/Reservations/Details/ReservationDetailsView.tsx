import React from "react";
import { useGetReservationById } from "../../../queries/reservation/useGetReservationById";
import { useHistory, useParams } from "react-router";
import { UserNotLogged } from "../../Error/UserNotLogged";
import { Loader } from "../../Loader/Loader";
import { CenteredCard } from "../../Card/CenteredCard";
import { Button, CardActions, CardContent, Divider, IconButton, Tooltip, Typography } from "@mui/material";
import { ErrorCard } from "../../Error/ErrorCard";
import { CardElement } from "./CardElement";
import { CookieName, getCookieValueByName } from "../../../utils/cookies";
import { useRemoveReservation } from "../../../queries/reservation/useRemoveReservation";
import { Role } from "../../../types/user";
import { routes } from "../../../utils/routes";
import ArrowBackIcon from '@mui/icons-material/ArrowBack';

export const ReservationDetailsView: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const history = useHistory();

  const userId = getCookieValueByName(CookieName.UserId);
  const role = getCookieValueByName(CookieName.Role);

  const { data, loading, error } = useGetReservationById({
    variables: { id }
  });

  const { remove } = useRemoveReservation()

  if (!id) {
    return <UserNotLogged />;
  }

  if (loading) {
    return <Loader />;
  }

  if (error) {
    return <ErrorCard text={error} />;
  }

  const { details, user, spot } = data ?? {};

  const isRemoveButtonEnabled = role === Role.Admin || userId === user?.id.toString()

  const handleRemove = async (id: string | undefined) => {
    await remove(id ?? '');
    history.push(routes.reservations);
  }

  return (
    <CenteredCard>
      <CardContent>
        <h3 style={styles.header}>Reservation Details</h3>
        <Typography gutterBottom variant="h5" component="div" textAlign='center'>
          {convertEpochToDate(details?.reservedFrom).toDateString()}
        </Typography>
        <Divider variant='middle' />
        <Typography gutterBottom />
        <CardElement prefix='Spot' text={spot?.location} />
        <CardElement
          prefix='Reserved from'
          text={convertEpochToDate(details?.reservedFrom).toLocaleTimeString()}
        />
        <CardElement
          prefix='Reserved to'
          text={convertEpochToDate(details?.reservedTo).toLocaleTimeString()}
          gutterBottom
        />
        <CardElement
          prefix='Reserved By'
          text={user?.name + ' ' + user?.surname}
        />
        <CardElement prefix='Email' text={user?.email} />
      </CardContent>
      <CardActions>
        <Tooltip title="Go Back">
          <IconButton onClick={() => history.goBack()}>
            <ArrowBackIcon />
          </IconButton>
        </Tooltip>
        {isRemoveButtonEnabled && (
          <>
            <Button
              size="small"
              color='error'
              style={styles.button}
              onClick={() => handleRemove(String(details?.id))}
            >
              Remove
            </Button>
            <Button
              size="small"
              style={styles.button}
            >
              Edit
            </Button>
          </>
        )}
      </CardActions>
    </CenteredCard>
  )
}

function convertEpochToDate(epoch: string | undefined): Date {
  return new Date(Number(epoch) * 1000)
}

const styles = {
  header: {
    display: 'flex',
    justifyContent: 'center',
  },
  button: {
    marginLeft: 'auto'
  }
};