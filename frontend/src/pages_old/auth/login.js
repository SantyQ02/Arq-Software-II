
import MainLayout from '@/layouts/MainLayout'
import Login from '@/components/auth/Login'
export default function LoginPage() {
  return (
    <>
      <MainLayout title={"Login"}>
        <Login />
      </MainLayout>
    </>
  )
}
