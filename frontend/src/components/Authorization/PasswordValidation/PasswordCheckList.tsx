import React from 'react';
import { Grid } from '@mui/material';
import { PasswordCheckItem } from './PasswordCheckItem';
import { PasswordRule } from '../../../utils/getPasswordValidationRules';

export const PasswordCheckList: React.FC<{ rules: Array<PasswordRule> }> = ({
  rules,
}) => (
  <Grid container direction='column' style={styles.grid}>
    {rules.map(({ rule, label }) => (
      <React.Fragment key={label}>
        <PasswordCheckItem isValid={rule} label={label} />
      </React.Fragment>
    ))}
  </Grid>
);

const styles = {
  grid: {
    marginTop: '0.8rem',
  },
};
