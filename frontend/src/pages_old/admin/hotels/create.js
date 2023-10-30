import CreateHotel from "@/components/admin/CreateHotel"
import HotelsList from "@/components/admin/HotelsList"
import DashboardAdmin from "@/layouts/DashboardAdmin"

import { UserContext } from "@/layouts/LayoutContext"
import { useRouter } from "next/router"
import { useContext, useEffect } from "react"

const DashboardHotelsCreate = () => {
  const [user, setUser] = useContext(UserContext);  
  const router = useRouter()

  useEffect(()=>{

    if (user === null || user.role !== "admin")
    router.push("/auth/login")
  },[])
  return (
      <DashboardAdmin title={"Create Hotel"} current={"/admin/hotels/create"}>

        {/* Replace with your content */}
          <CreateHotel />
        {/* /End replace */}
        
      </DashboardAdmin>
  )
}

export default DashboardHotelsCreate