import { Box, Card, Fade, Modal } from '@mui/material';
import React, { PropsWithChildren } from 'react';
import { MARGIN_TOP_CONTENT } from '../../utils/consts';

interface CenteredModalProps extends PropsWithChildren {
  isModalOpen: boolean;
  handleSubmit: (evt: React.FormEvent<HTMLFormElement>) => void;
}

export const CenteredModal: React.FC<CenteredModalProps> = ({
  children,
  isModalOpen,
  handleSubmit,
}) => (
  <Modal open={isModalOpen}>
    <Fade in={isModalOpen}>
      <Box component='form' onSubmit={handleSubmit} style={styles.box}>
        <Card style={styles.card}>{children}</Card>
      </Box>
    </Fade>
  </Modal>
);

const styles = {
  card: {
    width: '23rem',
    borderRadius: '0.5rem',
    boxShadow: '0.5rem 1rem 1rem rgba(0, 0, 0, 0.1)',
  },
  box: {
    display: 'flex',
    justifyContent: 'center',
    marginTop: MARGIN_TOP_CONTENT,
    borderRadius: '1rem',
    padding: '4rem',
  },
};
