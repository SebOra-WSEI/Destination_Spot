import React from "react";
import { useGetReservationById } from "../../../queries/reservation/useGetReservationById";
import { useParams } from "react-router";
import { UserNotLogged } from "../../Error/UserNotLogged";
import { Loader } from "../../Loader/Loader";
import { CenteredCard } from "../../Card/CenteredCard";
import { CardContent, Divider, Typography } from "@mui/material";
import { ErrorCard } from "../../Error/ErrorCard";
import { CardElement } from "./CardElement";

export const ReservationDetailsView: React.FC = () => {
  const { id } = useParams<{ id: string }>();

  const { data, loading, error } = useGetReservationById({
    variables: { id }
  });

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

  return (
    <CenteredCard>
      <h3 style={styles.header}>Reservation Details</h3>
      <CardContent>
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
};