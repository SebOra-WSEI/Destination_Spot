interface ErrorData {
  error: string;
}

export interface ErrorResponse {
  response: {
    status: number;
    data: ErrorData;
  };
}
