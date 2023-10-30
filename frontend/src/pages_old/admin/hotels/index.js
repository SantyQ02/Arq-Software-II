import HotelsList from "@/components/admin/HotelsList"
import DashboardAdmin from "@/layouts/DashboardAdmin"

import { UserContext } from "@/layouts/LayoutContext"
import { useRouter } from "next/router"
import { useContext, useEffect } from "react"

const DashboardHotelsList = () => {
  const [user, setUser] = useContext(UserContext);  
  const router = useRouter()

  useEffect(()=>{

    if (user === null || user.role !== "admin")
    router.push("/auth/login")
  },[])
  return (
      <DashboardAdmin title={"Hotels List"} current={"/admin/hotels"}>

        {/* Replace with your content */}
          <HotelsList />
        {/* /End replace */}
        
      </DashboardAdmin>
  )
}

export default DashboardHotelsList