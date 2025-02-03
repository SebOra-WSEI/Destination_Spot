import { Button, DialogActions, DialogContent } from '@mui/material';
import React, { useState } from 'react';
import { PasswordInput } from '../Authorization/Form/PasswordInput';
import { ResetPasswordBody } from '../../types/authorization';
import { FONT_FAMILY } from '../../utils/consts';
import { PasswordCheckList } from '../Authorization/PasswordValidation/PasswordCheckList';
import { getPasswordValidationRules } from '../../utils/getPasswordValidationRules';
import { useResetPassword } from '../../queries/user/useResetPassword';
import { CenteredModal } from '../Modal/CenteredModal';
import { useParams } from 'react-router';

interface ResetPasswordModalProps {
  isModalOpen: boolean;
  setIsModalOpen: (isModalOpen: boolean) => void;
}

export const ResetPasswordModal: React.FC<ResetPasswordModalProps> = ({
  isModalOpen,
  setIsModalOpen,
}) => {
  const { id: idParams } = useParams<{ id: string }>();
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
    });
  };

  const handlePasswordChange = (
    evt: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
    field: keyof ResetPasswordBody
  ): void => {
    setBody({
      ...body,
      [field]: evt.target.value,
    });
  };

  const { resetPassword } = useResetPassword(onCloseModal);

  const handleSubmit = async (
    evt: React.FormEvent<HTMLFormElement>
  ): Promise<void> => {
    evt.preventDefault();
    await resetPassword({
      ...(body.currentPassword
        ? {
            currentPassword: body.currentPassword,
          }
        : {}),
      newPassword: body.newPassword,
      confirmNewPassword: body.confirmNewPassword,
    });
  };

  const { currentPassword, newPassword, confirmNewPassword } = body;

  return (
    <CenteredModal isModalOpen={isModalOpen} handleSubmit={handleSubmit}>
      <>
        <DialogContent>
          <h3 style={styles.header}>Reset Password</h3>
          {!idParams && (
            <PasswordInput
              password={currentPassword}
              handlePasswordChange={(evt) =>
                handlePasswordChange(evt, 'currentPassword')
              }
              label='Current password'
            />
          )}
          <PasswordInput
            password={newPassword}
            handlePasswordChange={(evt) =>
              handlePasswordChange(evt, 'newPassword')
            }
            label='New password'
          />
          <PasswordInput
            password={confirmNewPassword}
            handlePasswordChange={(evt) =>
              handlePasswordChange(evt, 'confirmNewPassword')
            }
            label='Confirm password'
          />
          {Boolean(newPassword) && (
            <PasswordCheckList
              rules={getPasswordValidationRules(
                newPassword,
                confirmNewPassword
              )}
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
      </>
    </CenteredModal>
  );
};

const styles = {
  header: {
    display: 'flex',
    justifyContent: 'center',
  },
  button: {
    fontFamily: FONT_FAMILY,
    borderRadius: '0.5rem',
  },
};
