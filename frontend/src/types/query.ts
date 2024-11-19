export interface Query<T> {
  data?: T;
  loading?: boolean;
  error?: string;
  message?: string;
}

export interface QueryVariables<T> {
  variables: T;
  skip?: boolean;
}

export interface EmptyQueryVariables {
  skip?: boolean;
}
