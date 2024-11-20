import React, { useState } from 'react';
import { useGetCurrentUser } from '../../queries/user/getCurrentUser';
import {
  Avatar,
  Box,
  Button,
  CardActions,
  CardContent,
  Typography
} from '@mui/material';
import { FONT_FAMILY } from '../../utils/consts';
import EmailIcon from '@mui/icons-material/Email';
import ManageAccountsIcon from '@mui/icons-material/ManageAccounts';
import { CenteredCard } from '../Card/CenteredCard';
import { signOut } from '../../utils/signOut';
import { red } from '@mui/material/colors';
import { Loader } from '../Loader/Loader';
import { UnknownError } from '../Error/UnknownError';
import { routeBuilder } from '../../utils/routes';

export const UserView: React.FC = () => {
  const [isModalOpen, setIsModalOpen] = useState<boolean>(false);

  const { data, loading, error } = useGetCurrentUser();

  const { name, surname, email, role } = data ?? {};

  const userInitials = (name?.[0] ?? '') + (surname?.[0] ?? '');
  const user = name + ' ' + surname

  const openModal = (): void => setIsModalOpen(true);

  if (loading) {
    return <Loader />
  }

  if (error) {
    return <UnknownError link={routeBuilder.profile} />
  }

  return (
    <CenteredCard>
      <CardContent>
        <Box sx={styles.initials}>
          <Avatar sx={styles.avatar}>{userInitials}</Avatar>
          <Typography sx={styles.name}>{user}</Typography>
        </Box>
        <Box sx={styles.email}>
          <EmailIcon />
          <Typography>{email}</Typography>
        </Box>
        <Box sx={styles.role}>
          <ManageAccountsIcon />
          <Typography>{role?.toUpperCase()}</Typography>
        </Box>
      </CardContent>
      <CardActions>
        <Button
          size='small'
          onClick={openModal}
          style={styles.resetButton}
        >
          Reset Password
        </Button>
        <Button
          size='small'
          onClick={signOut}
          style={styles.signOutButton}
        >
          Sign out
        </Button>
      </CardActions>
    </CenteredCard>
    // <ResetPasswordModal
    //   user={user}
    //   isModalOpen={isModalOpen}
    //   setIsModalOpen={setIsModalOpen}
    //   setSeverity={setSeverity}
    //   setSnackbarMessage={setSnackbarMessage}
    // />
  );
};

const styles = {
  avatar: {
    backgroundColor: '#1cf618',
    width: '3rem',
    height: '3rem',
    fontSize: '20px',
    boxShadow: 18,
    fontFamily: FONT_FAMILY,
  },
  name: {
    fontFamily: FONT_FAMILY,
    fontSize: '18px',
    fontWeight: 'bold',
    marginTop: '1rem',
  },
  initials: {
    display: 'flex',
    alignItems: 'center',
    flexDirection: 'column',
    margin: '2rem',
  },
  email: {
    fontFamily: FONT_FAMILY,
    marginTop: '1rem',
    color: '#757575',
    display: 'flex',
    gap: '0.2rem',
    padding: '0.5rem',
  },
  role: {
    fontFamily: FONT_FAMILY,
    color: '#757575',
    display: 'flex',
    gap: '0.2rem',
    padding: '0.5rem',
  },
  signOutButton: {
    fontFamily: FONT_FAMILY,
    marginLeft: 'auto',
    color: red[700]
  },
  resetButton: {
    fontFamily: FONT_FAMILY,
  },
}
