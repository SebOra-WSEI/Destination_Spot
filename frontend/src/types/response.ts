interface ErrorData {
  error: string;
}

export interface ErrorResponse {
  response: {
    status: number;
    data: ErrorData;
  };
}

export interface CommonResponse<T> {
  status: number;
  data: T;
}
