export interface Spot {
  id: number;
  location: number;
}

export interface SpotResponse {
  response: {
    spots: Array<Spot>;
  };
}
