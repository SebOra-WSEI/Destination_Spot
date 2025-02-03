import React, { useState } from 'react';
import { useGetAllSpots } from '../../queries/spots/useGetAllSpots';
import { Loader } from '../Loader/Loader';
import { ErrorCard } from '../Error/ErrorCard';
import { Button, IconButton, List, ListItemText, Tooltip } from '@mui/material';
import { CommonListItem } from '../List/CommonListItem';
import { FONT_FAMILY } from '../../utils/consts';
import DeleteOutlineIcon from '@mui/icons-material/DeleteOutline';
import { CreateSpotModal } from './CreateSpotModal';
import { useRemoveSpot } from '../../queries/spots/useRemoveSpot';

export const SpotsList: React.FC = () => {
  const [isModalOpen, setIsModalOpen] = useState<boolean>(false);

  const { data, loading, error } = useGetAllSpots();
  const { remove } = useRemoveSpot();

  if (loading) {
    return <Loader />;
  }

  if (error) {
    return <ErrorCard text={error} />;
  }

  const handleModalOpen = (): void => setIsModalOpen(true);
  const handleRemove = async (id: number): Promise<void> => await remove(id);

  return (
    <>
      <List sx={styles.list}>
        <Button style={styles.button} onClick={handleModalOpen}>
          Create new spot
        </Button>
        {data?.map(({ id, location }) => (
          <CommonListItem key={id}>
            <ListItemText
              primaryTypographyProps={{
                fontFamily: FONT_FAMILY,
              }}
              primary={`Spot location: ${location}`}
            />
            <Tooltip title='Remove location'>
              <IconButton onClick={() => handleRemove(id)}>
                <DeleteOutlineIcon color='error' />
              </IconButton>
            </Tooltip>
          </CommonListItem>
        ))}
      </List>
      <CreateSpotModal
        isModalOpen={isModalOpen}
        setModalOpen={setIsModalOpen}
      />
    </>
  );
};

const styles = {
  list: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    marginTop: '2.5rem',
  },
  button: {
    fontFamily: FONT_FAMILY,
    fontSize: 15,
    marginBottom: '3rem',
  },
};
