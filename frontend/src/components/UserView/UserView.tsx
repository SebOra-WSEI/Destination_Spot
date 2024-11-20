import React from 'react';
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
import { UserNotLogged } from '../Error/UserNotLogged';
import { CookieName, getCookieValueByName } from '../../utils/cookies';

export const UserView: React.FC = () => {
  const userId = getCookieValueByName(CookieName.UserId)

  const { data, loading, error } = useGetCurrentUser({
    skip: !userId
  });

  const { name, surname, email, role, id } = data ?? {};

  const userInitials = (name?.[0] ?? '') + (surname?.[0] ?? '');

  if (!id) {
    return <UserNotLogged />;
  }

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
          <Typography style={styles.name}>{name + ' ' + surname}</Typography>
        </Box>
        <Box style={styles.email}>
          <EmailIcon />
          <Typography>{email}</Typography>
        </Box>
        <Box style={styles.role}>
          <ManageAccountsIcon />
          <Typography>{role?.toUpperCase()}</Typography>
        </Box>
      </CardContent>
      <CardActions>
        <Button
          size='small'
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
    fontFamily: FONT_FAMILY,
    display: 'flex',
    alignItems: 'center',
    flexDirection: 'column',
    marginBottom: '1rem',
  },
  email: {
    fontFamily: FONT_FAMILY,
    color: '#757575',
    display: 'flex',
    gap: '0.4rem',
  },
  role: {
    fontFamily: FONT_FAMILY,
    color: '#757575',
    display: 'flex',
    gap: '0.4rem',
    marginTop: '0.3rem'
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
