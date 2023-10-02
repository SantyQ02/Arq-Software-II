
import ResetPassword from '@/components/auth/ResetPassword'
import MainLayout from '@/layouts/MainLayout'
export default function ResetPass() {
  return (
    <>
      <MainLayout title={"Reset Password"}>
        <ResetPassword />
      </MainLayout>
    </>
  )
}
