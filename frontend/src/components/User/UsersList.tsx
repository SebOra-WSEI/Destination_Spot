import React from 'react';
import { useGetAllUsers } from '../../queries/user/useGetAllUsers';
import { Loader } from '../Loader/Loader';
import { ErrorCard } from '../Error/ErrorCard';
import { IconButton, List, ListItemText, Tooltip } from '@mui/material';
import { CommonListItem } from '../List/CommonListItem';
import { FONT_FAMILY } from '../../utils/consts';
import DeleteOutlineIcon from '@mui/icons-material/DeleteOutline';
import ManageAccountsIcon from '@mui/icons-material/ManageAccounts';
import { useHistory } from 'react-router';
import { routeBuilder } from '../../utils/routes';
import { useRemoveUser } from '../../queries/user/useRemoveUser';

export const UsersList: React.FC = () => {
  const history = useHistory();

  const { data, loading, error } = useGetAllUsers();
  const { remove } = useRemoveUser()

  if (loading) {
    return <Loader />;
  }

  if (error) {
    return <ErrorCard text={error} />;
  }

  const handleRemove = async (id: string | undefined): Promise<void> =>
    await remove(id);

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
          <Tooltip title="View details">
            <IconButton onClick={() => history.push(routeBuilder.userDetails(String(id)))}>
              <ManageAccountsIcon />
            </IconButton>
          </Tooltip>
          <Tooltip title="Remove user">
            <IconButton onClick={() => handleRemove(String(id))}>
              <DeleteOutlineIcon color='error' />
            </IconButton>
          </Tooltip>
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
};
