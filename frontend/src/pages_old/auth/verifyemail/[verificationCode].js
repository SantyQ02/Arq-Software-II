
import MainLayout from '@/layouts/MainLayout'
import VerifyEmail from '@/components/auth/VerifyEmail'

export default function VerifyEmailPage() {
  return (
    <>
      <MainLayout title={"Verify Email"}>
      <VerifyEmail />
      </MainLayout>
    </>
  )
}
