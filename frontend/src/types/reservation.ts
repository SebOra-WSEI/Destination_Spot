import { Spot } from './spot';
import { User } from './user';

export interface Reservation {
  details: Details;
  user: User;
  spot: Spot;
}

interface Details {
  id: number;
  spotId: number;
  userId: number;
  reservedFrom: string;
  reservedTo: string;
}
