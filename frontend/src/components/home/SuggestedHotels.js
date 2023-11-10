import { getAvailableHotels } from "@/lib/api/hotel"
import Image from "next/image"
import { useEffect, useState } from "react"

/*
  This example requires Tailwind CSS v2.0+ 
  
  This example requires some changes to your config:
  
  ```
  // tailwind.config.js
  module.exports = {
    // ...
    plugins: [
      // ...
      require('@tailwindcss/aspect-ratio'),
    ],
  }
  ```
*/
const products = [
  {
    id: 1,
    name: 'Black Basic Tee',
    price: '$32',
    href: '#',
    imageSrc: 'https://tailwindui.com/img/ecommerce-images/home-page-03-favorite-01.jpg',
    imageAlt: "Model wearing women's black cotton crewneck tee.",
  },
  {
    id: 1,
    name: 'Black Basic Tee',
    price: '$32',
    href: '#',
    imageSrc: 'https://tailwindui.com/img/ecommerce-images/home-page-03-favorite-01.jpg',
    imageAlt: "Model wearing women's black cotton crewneck tee.",
  },
  {
    id: 1,
    name: 'Black Basic Tee',
    price: '$32',
    href: '#',
    imageSrc: 'https://tailwindui.com/img/ecommerce-images/home-page-03-favorite-01.jpg',
    imageAlt: "Model wearing women's black cotton crewneck tee.",
  },
  // More products...
]

export default function SuggestedHotels({ dataCheck, setHotel }) {

  const [hotels, setHotels] = useState(null)
  const { rooms, date_in, date_out, currentHotelId } = dataCheck

  const get_available_hotels = async () => {
    let availableHotels = await getAvailableHotels(rooms, date_in, date_out)
    if (availableHotels !== null)
      availableHotels = availableHotels.filter(hotel => hotel.hotel_id !== currentHotelId)
    //console.log(availableHotels)
    setHotels(availableHotels)
  }
  useEffect(() => {
    get_available_hotels()

  }, [dataCheck])

  const OnClick = async (hotel) => setHotel(hotel)

  return (
    <div className="bg-white">
      <div className="max-w-7xl mx-auto py-16 px-4 sm:py-24 sm:px-6 lg:px-8">
        <div className="sm:flex sm:items-baseline sm:justify-between">
          <h2 className="text-2xl font-extrabold tracking-tight text-gray-900">Suggested Hotels</h2>

        </div>

        <div className="mt-6 grid grid-cols-1 gap-y-10 sm:grid-cols-3 sm:gap-y-0 sm:gap-x-6 lg:gap-x-8">
          {hotels !== null ?
          <>
           {hotels.slice(0,3).map((hotel) => (
            <div key={hotel.hotel_id} className="group relative">
              <div className="w-full h-96 rounded-lg overflow-hidden group-hover:opacity-75 sm:h-auto sm:aspect-w-2 sm:aspect-h-3">
                <Image
                  src={`${process.env.NEXT_PUBLIC_URL_SERVICE_HOTELS}/api/public/${hotel.photos !== null ? hotel.photos[0].url : "missing_hotel.png"}`}
                  alt={hotel.title}
                  className="w-full object-center object-cover h-32"
                  width={1000} height={1000}
                />
              </div>
              <h3 className="mt-4 text-base font-semibold text-gray-900">
                <button onClick={() => OnClick(hotel)}>
                  <span className="absolute inset-0" />
                  {hotel.title}
                </button>
              </h3>
              <p className="mt-1 text-sm text-gray-500">{hotel.price_per_day} USD per day</p>
            </div>
          ))
          }</>
        :
        <p  className="hidden text-sm font-semibold text-indigo-600 hover:text-indigo-500 sm:block">
            Select a date range to view suggestions<span aria-hidden="true"></span>
          </p>
        }
        </div>

      </div>
    </div>
  )
}
