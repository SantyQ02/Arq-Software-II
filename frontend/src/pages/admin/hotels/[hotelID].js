import HotelDetail from "@/components/admin/HotelDetail"
import DashboardAdmin from "@/layouts/DashboardAdmin"
import {getHotelById}from "@/lib/api/hotel"

import { UserContext } from "@/layouts/LayoutContext"
import { useRouter } from "next/router"
import { useContext, useEffect } from "react"
import { useState } from "react";

const DashboardHotelsDetail = () => {
  const [user, setUser] = useContext(UserContext);
  const router = useRouter()
  const hotel_id = router.query.hotelID
  const [hotel, setHotel] = useState(null)

  useEffect(() => {

    if (user === null || user.role !== "admin")
      router.push("/auth/login")
    
    const get_hotel_by_id = async () => {
      const data = await getHotelById(hotel_id)
      setHotel(data)
    }
    get_hotel_by_id()
    
  }, [])
  return (
    <DashboardAdmin title={"Hotels Detail"} current={"/admin/hotels"}>

      {/* Replace with your content */}
      {hotel != null && <HotelDetail hotel={hotel} setHotel={setHotel}/>}
      {/* /End replace */}

    </DashboardAdmin>
  )
}

export default DashboardHotelsDetail