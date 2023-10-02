
import ResetPasswordConfirm from '@/components/auth/ResetPasswordConfirm'
import MainLayout from '@/layouts/MainLayout'
export default function ResetPassVerify() {
  return (
    <>
      <MainLayout title={"Reset Password Verify"}>
        <ResetPasswordConfirm />
      </MainLayout>
    </>
  )
}
