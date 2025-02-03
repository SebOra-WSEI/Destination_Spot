export interface Spot {
  id: number;
  location: number;
}

export interface SpotResponse {
  response: {
    spots: Array<Spot>;
  };
}

export interface SpotData {
  response: {
    message: string;
    spot: Spot;
  };
}
