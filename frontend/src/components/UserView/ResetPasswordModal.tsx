import { Box, Button, Card, DialogActions, DialogContent, Fade, Modal } from '@mui/material';
import React, { useState } from 'react';
import { PasswordInput } from '../Authorization/Form/PasswordInput';
import { ResetPasswordBody } from '../../types/authorization';
import { FONT_FAMILY, MARGIN_TOP_CONTENT } from '../../utils/consts';
import { PasswordCheckList } from '../Authorization/PasswordValidation/PasswordCheckList';
import { getPasswordValidationRules } from '../../utils/getPasswordValidationRules';
import { useResetPassword } from '../../queries/user/useResetPassword';

interface ResetPasswordModalProps {
  isModalOpen: boolean
  setIsModalOpen: (isModalOpen: boolean) => void
}

export const ResetPasswordModal: React.FC<ResetPasswordModalProps> = ({
  isModalOpen,
  setIsModalOpen
}) => {
  const [body, setBody] = useState<ResetPasswordBody>({
    currentPassword: '',
    newPassword: '',
    confirmNewPassword: '',
  });

  const onCloseModal = () => {
    setIsModalOpen(false);
    setBody({
      currentPassword: '',
      newPassword: '',
      confirmNewPassword: '',
    })
  }

  const handlePasswordChange = (
    evt: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
    field: keyof ResetPasswordBody,
  ): void => {
    setBody({
      ...body,
      [field]: evt.target.value,
    });
  };

  const { resetPassword } = useResetPassword(onCloseModal);

  const handleSubmit = (evt: React.FormEvent<HTMLFormElement>) => {
    evt.preventDefault();
    resetPassword(body);
  }

  const { currentPassword, newPassword, confirmNewPassword } = body;

  return (
    <Modal open={isModalOpen} closeAfterTransition>
      <Fade in={isModalOpen}>
        <Box component='form' onSubmit={handleSubmit} style={styles.box}>
          <Card style={styles.card}>
            <DialogContent>
              <h3 style={styles.header}>Reset Password</h3>
              <PasswordInput
                password={currentPassword}
                handlePasswordChange={(evt) => handlePasswordChange(evt, 'currentPassword')}
                label='Current password'
              />
              <PasswordInput
                password={newPassword}
                handlePasswordChange={(evt) => handlePasswordChange(evt, 'newPassword')}
                label='New password'
              />
              <PasswordInput
                password={confirmNewPassword}
                handlePasswordChange={(evt) => handlePasswordChange(evt, 'confirmNewPassword')}
                label='Confirm password'
              />
              {Boolean(newPassword) && (
                <PasswordCheckList
                  rules={getPasswordValidationRules(newPassword, confirmNewPassword)}
                />
              )}
            </DialogContent>
            <DialogActions>
              <Button
                variant='outlined'
                onClick={onCloseModal}
                style={styles.button}
              >
                Cancel
              </Button>
              <Button
                variant='contained'
                color='success'
                type='submit'
                disabled={newPassword === '' || confirmNewPassword === ''}
                style={styles.button}
              >
                Save
              </Button>
            </DialogActions>
          </Card>
        </Box>
      </Fade>
    </Modal>
  )
}

const styles = {
  header: {
    display: 'flex',
    justifyContent: 'center',
  },
  card: {
    width: '23rem',
    borderRadius: '0.5rem',
    boxShadow: '0.5rem 1rem 1rem rgba(0, 0, 0, 0.1)',
  },
  box: {
    display: 'flex',
    justifyContent: 'center',
    marginTop: MARGIN_TOP_CONTENT,
    borderRadius: '1rem',
    padding: '4rem',
  },
  button: {
    fontFamily: FONT_FAMILY,
    borderRadius: '0.5rem',
  },
};