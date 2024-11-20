import { Card } from '@mui/material';
import React, { PropsWithChildren } from 'react';
import { MARGIN_TOP_CONTENT } from '../../utils/consts';


export const CenteredCard: React.FC<PropsWithChildren> = ({ children }) => (
  <Card style={styles.card}>
    {children}
  </Card>
)

const styles = {
  card: {
    bgcolor: 'background.paper',
    padding: '1.5rem',
    borderRadius: '1.5rem',
    marginTop: MARGIN_TOP_CONTENT
  },
};