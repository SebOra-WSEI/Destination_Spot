import React from 'react';
import { useGetAllUsers } from '../../queries/user/useGetAllUsers';
import { Loader } from '../Loader/Loader';
import { ErrorCard } from '../Error/ErrorCard';
import { IconButton, List, ListItemText, Tooltip } from '@mui/material';
import { CommonListItem } from '../List/CommonListItem';
import { FONT_FAMILY } from '../../utils/consts';
import PersonRemoveAlt1Icon from '@mui/icons-material/PersonRemoveAlt1';
import ManageAccountsIcon from '@mui/icons-material/ManageAccounts';
import { useHistory } from 'react-router';
import { routeBuilder } from '../../utils/routes';
import { useRemoveUser } from '../../queries/user/useRemoveUser';
import { Role } from '../../types/user';

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
          {role !== Role.Admin && (
            <Tooltip title="Remove user">
              <IconButton onClick={() => handleRemove(String(id))}>
                <PersonRemoveAlt1Icon color='error' />
              </IconButton>
            </Tooltip>
          )}
          <Tooltip title="View details">
            <IconButton onClick={() => history.push(routeBuilder.userDetails(String(id)))}>
              <ManageAccountsIcon />
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
