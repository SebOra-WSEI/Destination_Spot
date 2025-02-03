import React from "react";
import { useGetReservationById } from "../../../queries/reservation/useGetReservationById";
import { useParams } from "react-router";
import { UserNotLogged } from "../../Error/UserNotLogged";
import { Loader } from "../../Loader/Loader";
import { UnknownError } from "../../Error/UnknownError";
import { CenteredCard } from "../../Card/CenteredCard";
import { CardContent, Divider, Typography } from "@mui/material";

interface CardElementProps {
  prefix: string;
  text?: string | number
  gutterBottom?: boolean;
}

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
    return <UnknownError />;
  }

  const { details, user, spot } = data ?? {};

  return (
    <CenteredCard>
      <h3 style={styles.header}>Reservation Details</h3>
      <CardContent>
        <Typography gutterBottom variant="h5" component="div" textAlign='center'>
          {new Date(Number(details?.reservedFrom)).toDateString()}
        </Typography>
        <Divider variant='middle' />
        <Typography gutterBottom />
        <CardElement prefix='Spot' text={spot?.location} />
        <CardElement
          prefix='Reserved from'
          text={new Date(Number(details?.reservedFrom)).toLocaleTimeString()}
        />
        <CardElement
          prefix='Reserved to'
          text={new Date(Number(details?.reservedTo)).toLocaleTimeString()}
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

const CardElement: React.FC<CardElementProps> = ({ prefix, text, gutterBottom }) => (
  <Typography gutterBottom={gutterBottom} variant="body2" sx={{ color: 'text.secondary' }}>
    {prefix}: <strong>{text}</strong>
  </Typography>
)

const styles = {
  header: {
    display: 'flex',
    justifyContent: 'center',
  },
};