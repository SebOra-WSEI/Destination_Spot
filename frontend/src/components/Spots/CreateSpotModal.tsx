import React, { useState } from 'react';
import { CenteredModal } from '../Modal/CenteredModal';
import { Button, DialogActions, DialogContent, TextField } from '@mui/material';
import { BUTTON_RADIUS, FONT_FAMILY } from '../../utils/consts';
import { useCreateSpot } from '../../queries/spots/useCreateSpot';

interface CreateLocationModalProps {
  isModalOpen: boolean;
  setModalOpen: (isModalOpen: boolean) => void;
}

export const CreateSpotModal: React.FC<CreateLocationModalProps> = ({
  isModalOpen,
  setModalOpen,
}) => {
  const [location, setLocation] = useState<number | undefined>(undefined);

  const { create } = useCreateSpot();

  const onCloseModal = (): void => {
    setLocation(undefined);
    setModalOpen(false);
  };

  const handleSubmit = async (
    evt: React.FormEvent<HTMLFormElement>
  ): Promise<void> => {
    evt.preventDefault();
    await create(location as number);
  };

  return (
    <CenteredModal isModalOpen={isModalOpen} handleSubmit={handleSubmit}>
      <DialogContent>
        <h3 style={styles.header}>Create New Spot</h3>
        <TextField
          label='Spot location'
          variant='standard'
          type='number'
          autoComplete='Spot location'
          autoFocus
          required
          defaultValue={location}
          fullWidth
          onChange={(evt) => setLocation(Number(evt.target.value))}
        />
      </DialogContent>
      <DialogActions>
        <Button variant='outlined' onClick={onCloseModal} style={styles.button}>
          Close
        </Button>
        <Button
          variant='contained'
          color='success'
          type='submit'
          style={styles.button}
          disabled={!location}
        >
          Create
        </Button>
      </DialogActions>
    </CenteredModal>
  );
};

const styles = {
  header: {
    display: 'flex',
    justifyContent: 'center',
  },
  button: {
    fontFamily: FONT_FAMILY,
    borderRadius: BUTTON_RADIUS,
  },
};
