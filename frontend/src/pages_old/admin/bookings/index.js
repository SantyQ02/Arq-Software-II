import DashboardBookings from "@/components/admin/DashboardBookings"
import UsersList from "@/components/admin/UsersList"
import DashboardAdmin from "@/layouts/DashboardAdmin"
import { UserContext } from "@/layouts/LayoutContext"
import { useRouter } from "next/router"
import { useContext, useEffect } from "react"

const DashboardAdminBookings = () => {
  const [user, setUser] = useContext(UserContext);  
  const router = useRouter()

  useEffect(()=>{

    if (user === null || user.role !== "admin")
    router.push("/auth/login")
  },[])
  return (
      <DashboardAdmin title={"Bookings"} current={"/admin/bookings"}>

        {/* Replace with your content */}
          <DashboardBookings />
        {/* /End replace */}
        
      </DashboardAdmin>
  )
}

export default DashboardAdminBookings