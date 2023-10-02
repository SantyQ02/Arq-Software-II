import { searchBookings, searchBookingsMe } from '@/lib/api/booking';
import { SortAscendingIcon, HomeIcon, UsersIcon } from '@heroicons/react/solid'
import { useEffect, useState } from 'react';
import Datepicker from "react-tailwindcss-datepicker";


export default function SearchBar({ setBookings, admin=false }) {

  const [hotel, setHotel] = useState('')
  const OnChange = e => setHotel(e.target.value)

  const [user, setUser] = useState('')
  const OnUserChange = e => setUser(e.target.value)

  const [dates, setDates] = useState({
    startDate: null,
    endDate: null
  });
  const handleDatesChange = (newDates) => setDates(newDates)


  const search_bookings_me = async () => {
    const start_date = dates.startDate !== null ? new Date(dates.startDate).toISOString() : ""
    const end_date = dates.startDate !== null ? new Date(dates.endDate).toISOString() : ""
    const bookings = await searchBookingsMe(hotel, start_date, end_date)
    setBookings(bookings)
  }

  const search_bookings = async () => {
    const start_date = dates.startDate !== null ? new Date(dates.startDate).toISOString() : ""
    const end_date = dates.startDate !== null ? new Date(dates.endDate).toISOString() : ""
    const bookings = await searchBookings(hotel, user, start_date, end_date)
    setBookings(bookings)
  }

  useEffect(() => {
    if (admin == true) {
      search_bookings()
    } else {
      search_bookings_me()
    }

  }, [])

  useEffect(() => {
    if (admin == true) {
      search_bookings()
    } else {
      search_bookings_me()
    }

  }, [hotel, user, dates])




  return (
    <div>
      <label htmlFor="email" className="block text-sm font-medium text-gray-700">
        Search bookings
      </label>
      <div className="mt-1 sm:flex rounded-md shadow-sm">
   
      
    
        <div className="mt-1 relative flex items-stretch flex-grow focus-within:z-10">
          <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
            <HomeIcon className="h-5 w-5 text-gray-400" aria-hidden="true" />
          </div>
          <input
            type="text"
            name="text"
            id="text"
            value={hotel}
            onChange={e => OnChange(e)}
            className="focus:ring-indigo-500 focus:border-indigo-500 block w-full rounded-none rounded-l-md pl-10 sm:text-sm border-gray-300"
            placeholder="Hotel ID, Hotel Name"
          />
        </div>
        {
          admin &&
          <>
          <div className="mt-1 relative flex items-stretch flex-grow focus-within:z-10">
          <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
            <UsersIcon className="h-5 w-5 text-gray-400" aria-hidden="true" />
          </div>
          <input
            type="text"
            name="text"
            id="text"
            value={user}
            onChange={e => OnUserChange(e)}
            className="focus:ring-indigo-500 focus:border-indigo-500 block w-full rounded-none rounded-l-md pl-10 sm:text-sm border-gray-300"
            placeholder="User ID, User Name"
          />
        </div>
        
          </>
        }
        
        <div className="m-2">
          <Datepicker
            value={dates}
            onChange={handleDatesChange}
            className="h-full rounded-md border-0 bg-transparent py-0 pl-2 pr-7 text-gray-500 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm"

          />
        </div>

      </div>
    </div>
  )
}