import UserDetail from "@/components/admin/UserDetail"
import DashboardAdmin from "@/layouts/DashboardAdmin"

import { UserContext } from "@/layouts/LayoutContext"
import { useRouter } from "next/router"
import { useContext, useEffect } from "react"

const DashboardUserDetail = () => {
  const [user, setUser] = useContext(UserContext);
  const router = useRouter()

  useEffect(() => {

    if (user === null || user.role !== "admin")
      router.push("/auth/login")
  }, [])
  return (
    <DashboardAdmin title={"User Detail"} current={"/admin/users"}>

      {/* Replace with your content */}
      <UserDetail />
      {/* /End replace */}

    </DashboardAdmin>
  )
}

export default DashboardUserDetail