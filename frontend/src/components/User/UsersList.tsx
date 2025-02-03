import React from 'react';
import { useGetAllUsers } from '../../queries/user/useGetAllUsers';
import { Loader } from '../Loader/Loader';
import { ErrorCard } from '../Error/ErrorCard';
import { List, ListItemText } from '@mui/material';
import { CommonListItem } from '../List/CommonListItem';
import { FONT_FAMILY } from '../../utils/consts';

export const UsersList: React.FC = () => {

  const { data, loading, error } = useGetAllUsers()

  if (loading) {
    return <Loader />;
  }

  if (error) {
    return <ErrorCard text={error} />;
  }

  return (
    <List sx={styles.list}>
      {data?.map(({ id, name, surname, email, role }) => (
        <CommonListItem key={id}>
          <ListItemText
            secondaryTypographyProps={{
              fontFamily: FONT_FAMILY,
            }}
            secondary={role}
          />
          <ListItemText
            primaryTypographyProps={{
              fontFamily: FONT_FAMILY,
            }}
            primary={`${name} ${surname}`}
          />
          <ListItemText
            secondaryTypographyProps={{
              fontFamily: FONT_FAMILY,
            }}
            secondary={email}
          />
        </CommonListItem>
      ))}
    </List>
  )
}

const styles = {
  list: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    marginTop: '2.5rem',
  },
  icon: {
    marginLeft: 'auto'
  },
};
