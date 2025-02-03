import { Typography } from "@mui/material";
import React from "react";

interface CardElementProps {
  prefix: string;
  text?: string | number
  gutterBottom?: boolean;
}

export const CardElement: React.FC<CardElementProps> = ({ prefix, text, gutterBottom }) => (
  <Typography gutterBottom={gutterBottom} variant="body2" sx={{ color: 'text.secondary' }}>
    {prefix}: <strong>{text}</strong>
  </Typography>
)