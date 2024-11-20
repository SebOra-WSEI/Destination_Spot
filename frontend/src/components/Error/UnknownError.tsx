import React from 'react';
import { CenteredCard } from '../Card/CenteredCard';
import { Button, CardActions, CardContent } from '@mui/material';
import { useHistory } from 'react-router';
import { FONT_FAMILY } from '../../utils/consts';


export const UnknownError: React.FC<{ link?: string }> = ({ link }) => {
  const history = useHistory();

  const handleClick = (): void => history.push(link ?? '')

  return (
    <CenteredCard>
      <CardContent>
        <h3 style={styles.header}>Unknown Error</h3>
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
};

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
  },
}