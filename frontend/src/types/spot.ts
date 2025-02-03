export interface Spot {
  id: number;
  location: number;
}

export interface SpotsResponse {
  response: {
    spots: Array<Spot>;
  };
}

export interface SpotResponse {
  response: {
    message: string;
    spot: Spot;
  };
}
