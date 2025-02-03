import React from 'react';
import { useGetAllSpots } from '../../queries/spots/useGetAllSpots';
import { Loader } from '../Loader/Loader';
import { ErrorCard } from '../Error/ErrorCard';
import { IconButton, List, ListItemText, Tooltip } from '@mui/material';
import { CommonListItem } from '../List/CommonListItem';
import { FONT_FAMILY } from '../../utils/consts';
import DeleteOutlineIcon from '@mui/icons-material/DeleteOutline';

export const SpotsList: React.FC = () => {

  const { data, loading, error } = useGetAllSpots();

  if (loading) {
    return <Loader />;
  }

  if (error) {
    return <ErrorCard text={error} />;
  }

  return (
    <List sx={styles.list}>
      {data?.map(({ id, location }) => (
        <CommonListItem key={id}>
          <ListItemText
            primaryTypographyProps={{
              fontFamily: FONT_FAMILY,
            }}
            primary={`Spot location: ${location}`}
          />
          <Tooltip title="Remove location">
            <IconButton>
              <DeleteOutlineIcon color='error' />
            </IconButton>
          </Tooltip>
        </CommonListItem>
      ))}
    </List>
  );
};

const styles = {
  list: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    marginTop: '2.5rem',
  },
};
