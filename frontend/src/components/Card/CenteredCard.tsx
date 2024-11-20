import { Box, Card } from '@mui/material';
import React, { PropsWithChildren } from 'react';
import { MARGIN_TOP_CONTENT } from '../../utils/consts';

interface CenteredCardProps extends PropsWithChildren {
  isErrorCard?: boolean;
}

export const CenteredCard: React.FC<CenteredCardProps> = ({
  children,
  isErrorCard,
}) => (
  <Box style={styles.box}>
    <Card
      style={{
        ...styles.card(isErrorCard),
      }}
    >
      {children}
    </Card>
  </Box>
);

const styles = {
  card: (isErrorCard: boolean | undefined) => ({
    bgcolor: 'background.paper',
    padding: '1.5rem',
    borderRadius: '1.5rem',
    ...(isErrorCard
      ? {
          width: '30rem',
        }
      : {}),
  }),
  box: {
    display: 'flex',
    justifyContent: 'center',
    marginTop: MARGIN_TOP_CONTENT,
  },
};
