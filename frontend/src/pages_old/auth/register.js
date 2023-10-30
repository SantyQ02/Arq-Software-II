
import Register from '@/components/auth/Register'
import MainLayout from '@/layouts/MainLayout'
export default function RegisterPage() {
  return (
    <>
      <MainLayout title={"Register"}>
        <Register />
      </MainLayout>
    </>
  )
}
