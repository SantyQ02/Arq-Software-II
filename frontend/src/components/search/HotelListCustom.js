
import { useState } from 'react'
import Image from 'next/image'
import { useRouter } from 'next/router'
import Link from 'next/link'

export default function HotelsList({ hotels }) {

    const handleHotelDetail = (e, hotel) => {
        e.preventDefault()


    }

    const router = useRouter()

    return (
        <>
            <div className="bg-white" id="hotels-list">
                <div className="max-w-2xl mx-auto pb-24 px-4 sm:px-6 lg:max-w-7xl lg:px-8">
                    <h2 className="text-xl font-extrabold tracking-tight text-gray-900 sm:text-xl">Search Results</h2>
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
                                                    // src={`${hotel.thumbnail ?? "missing_hotel.png"}`}
                                                    src={`${process.env.NEXT_PUBLIC_URL_SERVICE_HOTELS}/api/public/${hotel.thumbnail ?? "missing_hotel.png"}`}
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

                                                <Link href={`/hotel/${hotel.hotel_id}?check_in_date=${router.query.check_in_date}&check_out_date=${router.query.check_out_date}`}>
                                                <button
                                                    className="inline-block rounded-md border border-transparent bg-indigo-600 px-3 md:px-8 py-1 md:py-3 text-center font-medium text-white hover:bg-indigo-700"
                                                >
                                                    Book
                                                </button>
                                                </Link>
                                            </div>
                                        </li>

                                    
                                ))}
                            </ul>
                        </section>

                    </form>
                </div>
            </div>
        </>
    )
}
