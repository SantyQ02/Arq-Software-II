import HotelDetail from "@/components/admin/HotelDetail"
import DashboardAdmin from "@/layouts/DashboardAdmin"
import {getHotelById}from "@/lib/api/hotel"

import { UserContext } from "@/layouts/LayoutContext"
import { useRouter } from "next/router"
import { useContext, useEffect } from "react"
import { useState } from "react";

export async function getServerSideProps({query}){
  const _hotel = await getHotelById(query.hotelID)
  return {
    props:{
      _hotel
    }
  }
}

export default function DashboardHotelsDetail({_hotel}){
  const [user, setUser] = useContext(UserContext);
  const router = useRouter()
  const [hotel, setHotel] = useState(_hotel)

  // useEffect(()=>{
  //   if (user === null || user.role !== "admin")
  //     router.push("/auth/login")
  // },[user])

  return (
    <DashboardAdmin title={"Hotels Detail"} current={"/admin/hotels"}>

      {/* Replace with your content */}
      {hotel != null && <HotelDetail hotel={hotel} setHotel={setHotel}/>}
      {/* /End replace */}

    </DashboardAdmin>
  )
}