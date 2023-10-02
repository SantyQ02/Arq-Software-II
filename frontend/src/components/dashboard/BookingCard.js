

import { setActiveBooking } from "@/lib/api/booking"
import { getHotelById } from "@/lib/api/hotel"
import { useEffect, useState } from "react"

  
  function classNames(...classes) {
    return classes.filter(Boolean).join(' ')
  }
  
  export default function BookingCard({booking_data}) {
    const [active, setActive] = useState(booking_data.active)
    const [booking, setBooking] = useState(booking_data)

    const [hotel, setHotel] = useState(null)

    useEffect(() => {
      const get_hotel_by_id = async () => {
        const data = await getHotelById(booking.hotel_id)
        setHotel(data)
      }
      get_hotel_by_id()
    }, [])

    useEffect(() => {
      setBooking(booking_data)
      const get_hotel_by_id = async () => {
        const data = await getHotelById(booking.hotel_id)
        setHotel(data)
      }
      get_hotel_by_id()
    }, [booking, booking_data])

    const handleCancelBooking = async (e) => {
      e.preventDefault()
      await setActiveBooking(booking.booking_id, false)
      const newBooking = booking
      newBooking.active = false
      setBooking(newBooking)
      setActive(false)
  }

  const handleActiveBooking = async (e) => {
      e.preventDefault()
      await setActiveBooking(booking.booking_id, true)
      const newBooking = booking
      newBooking.active = true
      setBooking(newBooking)
      setActive(true)
  }

    return (
      <div className="bg-white rounded-md shadow-md p-4 mb-4">
        <div className="flex justify-between items-center mb-4">
          <div className="text-sm font-bold">Booking Status</div>
          <div className="text-sm">
          <div className=" sm:flex sm:flex-col sm:items-center ">
              <p className="text-sm leading-6 text-gray-900">Active</p>


              {booking.active || active ?

                <label className="relative inline-flex items-center cursor-pointer"
                  onClick={e => handleCancelBooking(e)}>
                  <input type="checkbox" checked={true} className="sr-only peer" onChange={()=>{}} />
                  <div className="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-indigo-300 dark:peer-focus:ring-indigo-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-indigo-600"></div>
                  <span className=" text-sm font-medium text-gray-900 dark:text-gray-300"></span>
                </label>
                :
                <label className="relative inline-flex items-center cursor-pointer"
                  onClick={e => handleActiveBooking(e)}>
                  <input type="checkbox" checked={false} className="sr-only peer" onChange={()=>{}}/>
                  <div className="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-indigo-300 dark:peer-focus:ring-indigo-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-indigo-600"></div>
                  <span className=" text-sm font-medium text-gray-900 dark:text-gray-300"></span>
                </label>
              }
            </div>
          </div>
        </div>
        <div className="flex justify-between mb-4">
          <div className="text-sm font-bold">Booking ID:</div>
          <div className="text-sm">{booking.booking_id}</div>
        </div>
        <div className="flex justify-between mb-4">
          <div className="text-sm font-bold">Hotel:</div>
          <div className="text-sm">{hotel != null && hotel.title}</div>
        </div>
        <div className="flex justify-between mb-4">
          <div className="text-sm font-bold">Hotel ID:</div>
          <div className="text-sm">{booking.hotel_id}</div>
        </div>
        <div className="text-sm mb-4">Rooms: {booking.rooms}</div>
        <div className="text-sm mb-4">Price: {booking.total}</div>
        <div className="flex justify-between border-t border-gray-200 pt-4">
          <div className="text-gray-500">Date in: <span className="font-medium text-gray-900">{booking.date_in.substring(0, 10)}</span></div>
          <div className="text-gray-500">Date out: <span className="font-medium text-gray-900">{booking.date_out.substring(0, 10)}</span></div>
        </div>
      </div>


    )
  }
  