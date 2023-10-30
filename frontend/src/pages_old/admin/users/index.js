import UsersList from "@/components/admin/UsersList"
import DashboardAdmin from "@/layouts/DashboardAdmin"
import { UserContext } from "@/layouts/LayoutContext"
import { getUsers } from "@/lib/api/user"
import { useRouter } from "next/router"
import { useContext, useEffect, useState } from "react"


const Dashboard = () => {
  const [user, setUser] = useContext(UserContext);
  const [users, setUsers] = useState(null);
  const router = useRouter()

  useEffect(() => {

    if (user === null || user.role !== "admin")
      router.push("/auth/login")

    const get_users = async () => {
      const data = await getUsers()
      setUsers(data)
    }
    get_users()

  }, [])
  return (
    <DashboardAdmin title={"Users List"} current={"/admin/users"}>

      {/* Replace with your content */}
      <UsersList users = {users} setUsers={setUsers}/>
      {/* /End replace */}

    </DashboardAdmin>
  )
}

export default Dashboard