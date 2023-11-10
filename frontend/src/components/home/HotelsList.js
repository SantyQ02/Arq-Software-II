/*
  This example requires Tailwind CSS v2.0+ 
  
  This example requires some changes to your config:
  
  ```
  // tailwind.config.js
  module.exports = {
    // ...
    plugins: [
      // ...
      require('@tailwindcss/forms'),
    ],
  }
  ```
*/
import { CheckIcon, ClockIcon, QuestionMarkCircleIcon, XIcon } from '@heroicons/react/solid'
import HotelDetail from './HotelDetail'
import { useEffect, useState } from 'react'
import { getHotels } from '@/lib/api/hotel'
import Image from 'next/image'

const products = [
    {
        id: 1,
        name: 'Basic Tee',
        href: '#',
        price: '$32.00',
        color: 'Sienna',
        inStock: true,
        size: 'Large',
        imageSrc: 'https://tailwindui.com/img/ecommerce-images/shopping-cart-page-01-product-01.jpg',
        imageAlt: "Front of men's Basic Tee in sienna.",
    },
    {
        id: 2,
        name: 'Basic Tee',
        href: '#',
        price: '$32.00',
        color: 'Black',
        inStock: false,
        leadTime: '3–4 weeks',
        size: 'Large',
        imageSrc: 'https://tailwindui.com/img/ecommerce-images/shopping-cart-page-01-product-02.jpg',
        imageAlt: "Front of men's Basic Tee in black.",
    },
    {
        id: 3,
        name: 'Nomad Tumbler',
        href: '#',
        price: '$35.00',
        color: 'White',
        inStock: true,
        imageSrc: 'https://tailwindui.com/img/ecommerce-images/shopping-cart-page-01-product-03.jpg',
        imageAlt: 'Insulated bottle with white base and black snap lid.',
    },
]



export default function HotelsList({ hotels }) {

    const [open, setOpen] = useState(false)
    const [hotelDetail, setHotelDetail] = useState(null)

    const handleHotelDetail = (e, hotel) => {
        e.preventDefault()
        //console.log(e)
        setHotelDetail(hotel)
        setOpen(true)

    }

    return (
        <>
            <div className="bg-white" id="hotels-list">
                <div className="max-w-2xl mx-auto pt-16 pb-24 px-4 sm:px-6 lg:max-w-7xl lg:px-8">
                    <h1 className="text-3xl font-extrabold tracking-tight text-gray-900 sm:text-4xl">Hotels List</h1>
                    <form className="mt-12 lg:grid lg:grid-cols-12 lg:gap-x-12 lg:items-start xl:gap-x-16">
                        <section aria-labelledby="cart-heading" className="lg:col-span-12">
                            <h2 id="cart-heading" className="sr-only">
                                Items in your shopping cart
                            </h2>

                            <ul role="list" className="border-t border-b border-gray-200 divide-y divide-gray-200">
                                {hotels !== null && hotels !== undefined && hotels.map((hotel, hotelIdx) => (
                                    
                                        <li key={hotelIdx} className="flex py-6 sm:py-10">
                                            <div className="flex-shrink-0">
                                                <Image
                                                    src={`${process.env.NEXT_PUBLIC_URL_SERVICE_HOTELS}/api/public/${hotel.photos != null ? hotel.photos[0].url : "missing_hotel.png"}`}
                                                    alt={hotel.title}
                                                    width={1000}
                                                    height={1000}
                                                    className="w-24 h-24 rounded-md object-center object-cover sm:w-48 sm:h-48"
                                                />
                                            </div>

                                            <div className="ml-4 flex-1 flex flex-col justify-between sm:ml-6">
                                                <div className="relative pr-9 sm:grid sm:grid-cols-2 sm:gap-x-6 sm:pr-0">
                                                    <div>
                                                        <div className="flex justify-between">
                                                            <h3 className="text-xl font-bold text-gray-700 hover:text-gray-800">

                                                                {hotel.title.toUpperCase()}
                                                                <p className="text-base font-medium text-gray-700 hover:text-gray-800">
                                                                    {hotel.description}
                                                                </p>
                                                            </h3>
                                                        </div>

                                                        <p className="mt-1 text-sm font-medium text-gray-900">{hotel.price_per_day} USD per day</p>
                                                    </div>



                                                </div>

                                            </div>

                                            <div>

                                                <button
                                                    onClick={e => handleHotelDetail(e, hotel)}
                                                    className="inline-block rounded-md border border-transparent bg-indigo-600 px-3 md:px-8 py-1 md:py-3 text-center font-medium text-white hover:bg-indigo-700"
                                                >
                                                    Book
                                                </button>
                                            </div>
                                        </li>

                                    
                                ))}
                            </ul>
                        </section>

                    </form>
                </div>
            </div>

            <HotelDetail open={open} setOpen={setOpen} hotel={hotelDetail} setHotel={setHotelDetail} />

        </>
    )
}
