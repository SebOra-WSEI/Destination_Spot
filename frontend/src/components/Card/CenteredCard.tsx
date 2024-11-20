import { Box, Card } from '@mui/material';
import React, { PropsWithChildren } from 'react';
import { MARGIN_TOP_CONTENT } from '../../utils/consts';

interface CenteredCardProps extends PropsWithChildren {
  isErrorCard?: boolean
}

export const CenteredCard: React.FC<CenteredCardProps> = ({ children, isErrorCard }) => (
  <Box sx={styles.box}>
    <Card style={{
      ...styles.card,
      ...(isErrorCard ? {
        width: '30rem'
      } : {})
    }}>
      {children}
    </Card >
  </Box>
)

const styles = {
  card: {
    bgcolor: 'background.paper',
    padding: '1.5rem',
    borderRadius: '1.5rem',
  },
  box: {
    display: 'flex',
    justifyContent: 'center',
    marginTop: MARGIN_TOP_CONTENT,
  },
};