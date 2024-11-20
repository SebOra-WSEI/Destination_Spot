import React from 'react';
import { useHistory } from 'react-router';
import { CenteredCard } from '../Card/CenteredCard';
import { Button, CardActions, CardContent } from '@mui/material';
import { FONT_FAMILY } from '../../utils/consts';

interface ErrorProps {
  text: string
  link?: string
  isErrorCard?: boolean
}

export const ErrorCard: React.FC<ErrorProps> = ({ text, link, isErrorCard }) => {
  const history = useHistory();

  const handleClick = (): void => history.push(link ?? '')

  return (
    <CenteredCard isErrorCard={isErrorCard}>
      <CardContent>
        <h3 style={styles.header}>{text}</h3>
      </CardContent>
      {link && (
        <CardActions>
          <Button
            onClick={handleClick}
            style={styles.button}
            size='small'
          >
            {link.slice(1)}
          </Button>
        </CardActions>
      )}
    </CenteredCard>
  );
}

const styles = {
  button: {
    marginLeft: 'auto',
    fontFamily: FONT_FAMILY,
  },
  header: {
    fontSize: '1.5rem',
    marginBottom: '-0.5rem',
    fontFamily: FONT_FAMILY,
    color: '#757575',
    display: 'flex',
    justifyContent: 'center'
  },
}