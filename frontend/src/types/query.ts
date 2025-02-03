export interface Query<T> {
  data?: T;
  loading?: boolean;
  error?: string;
  message?: string;
}

export interface QueryVariables {
  variables: {
    id: string;
  };
  skip?: boolean;
}
