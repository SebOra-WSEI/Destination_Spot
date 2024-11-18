import React from 'react';
import CheckCircleIcon from '@mui/icons-material/CheckCircle';
import CancelIcon from '@mui/icons-material/Cancel';
import { Grid } from '@mui/material';
import { FONT_FAMILY } from '../../../utils/consts';
import { green, red } from '@mui/material/colors';

const FONT_SIZE = 12;

interface PasswordChecksProps {
  isValid: boolean;
  label: string;
}

export const PasswordCheckItem: React.FC<PasswordChecksProps> = ({
  isValid,
  label,
}) => (
  <Grid item container style={styles.gridItem}>
    <ValidIcons isValid={isValid} />
    <label style={styles.label}>{label}</label>
  </Grid>
);

const ValidIcons: React.FC<Pick<PasswordChecksProps, 'isValid'>> = ({
  isValid,
}) =>
  isValid ? (
    <CheckCircleIcon style={styles.checkIcon} />
  ) : (
    <CancelIcon style={styles.cancelIcon} />
  );

const styles = {
  grid: {
    marginTop: '0.8rem',
  },
  gridItem: {
    marginTop: '0.1rem',
  },
  label: {
    marginLeft: '0.2rem',
    fontSize: FONT_SIZE,
    fontFamily: FONT_FAMILY,
  },
  checkIcon: {
    color: green[500],
    fontSize: FONT_SIZE,
  },
  cancelIcon: {
    color: red[500],
    fontSize: FONT_SIZE,
  },
};