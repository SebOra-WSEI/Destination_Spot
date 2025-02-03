import { Spot } from './spot';
import { User } from './user';

interface Details {
  id: number;
  spotId: number;
  userId: number;
  reservedFrom: string;
  reservedTo: string;
}

export interface Reservation {
  details: Details;
  user: User;
  spot: Spot;
}

export interface ReservationBody {
  userId: number;
  spotId: number;
  reservedFrom: string;
  reservedTo: string;
}

export interface ReservationsResponse {
  response: {
    reservations: Array<Reservation>;
  };
}

export interface ReservationResponse {
  response: {
    message: string;
    reservation: Reservation;
  };
}
