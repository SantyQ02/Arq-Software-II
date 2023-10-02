import { getHotels, updateHotel } from "@/lib/api/hotel"
import Image from "next/image";
import Link from "next/link"
import { useEffect, useState } from "react";

export default function HotelsList() {

    const [hotels, setHotels] = useState(null)
    const get_hotels = async () => {
        const data = await getHotels()
        setHotels(data)
    }

    useEffect(() => {

        get_hotels()
    }, [])

    const handleCancelHotel = async (e, hotel) => {
        e.preventDefault()
        await updateHotel(hotel.hotel_id, hotel.title, hotel.description, hotel.price_per_day, hotel.rooms, false)
        get_hotels()
    }

    const handleRegisterHotel = async (e, hotel) => {
        e.preventDefault()
        await updateHotel(hotel.hotel_id, hotel.title, hotel.description, hotel.price_per_day, hotel.rooms, true)
        get_hotels()
    }

    return (
        <>
            <div className="bg-white">
                <div className="max-w-2xl mx-auto  pb-24 lg:max-w-7xl ">
                    
                    <form className="mt-3 lg:grid lg:grid-cols-12 lg:gap-x-12 lg:items-start xl:gap-x-16">
                        <section aria-labelledby="cart-heading" className="lg:col-span-12">
                            <h2 id="cart-heading" className="sr-only">
                                Items in your shopping cart
                            </h2>

                            <ul role="list" className="border-t border-b border-gray-200 divide-y divide-gray-200">
                                {hotels !== null && hotels !== undefined && hotels.map((hotel, hotelIdx) => (
                                    <>

                                        <li key={hotel.hotel_id} className="flex py-6 sm:py-10 justify-between">
                                            <div className=" flex ">
                                                
                                            <div className="ml-4 sm:ml-0 flex-shrink-0">
                                                <Image
                                                    src={`${process.env.NEXT_PUBLIC_URL_API}/api/public/${hotel.photos != null ? hotel.photos[0].url : ""}`}
                                                    alt={hotel.title}
                                                    className="w-24 h-24 rounded-md object-center object-cover sm:w-48 sm:h-48"
                                                    width={1000}
                                                    height={1000}
                                                />
                                            </div>

                                            <div className="sm:mt-0 ml-4 flex-1 flex flex-col justify-between sm:ml-6">
                                                <div className="relative pr-9 sm:grid sm:grid-cols-2 sm:gap-x-6 sm:pr-0">
                                                    <div>
                                                        <div className="flex justify-between">
                                                            <h3 className="text-lg font-bold text-gray-700 hover:text-gray-800">

                                                                Name: {hotel.title}
                                                                <p className="font-medium text-gray-700 hover:text-gray-800">
                                                                    Description: {hotel.description}
                                                                </p>
                                                            </h3>
                                                        </div>

                                                        <p className="mt-1 text-sm font-medium text-gray-900">{hotel.price_per_day} USD per day</p>
                                                    </div>



                                                </div>

                                            </div>
                                            </div>

                                            <div className="flex flex-col justify-center items-center">
                                                <Link
                                                    href={`/admin/hotels/${hotel.hotel_id}`}
                                                >
                                                    <button
                                                        className="inline-block rounded-md border border-transparent bg-indigo-600 px-2 py-1 md:py-2 text-center font-medium text-white hover:bg-indigo-700"
                                                    >
                                                        Update
                                                    </button>
                                                </Link>
                                                <div className="items-center justify-center mt-2">

                                                    <div className="flex flex-col items-center">


                                                        Active
                                                        {hotel.active ?

                                                            <label class="relative inline-flex items-center cursor-pointer"
                                                                onClick={e => handleCancelHotel(e, hotel)}>
                                                                <input type="checkbox" checked={true} class="sr-only peer" />
                                                                <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-indigo-300 dark:peer-focus:ring-indigo-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-indigo-600"></div>
                                                                <span class="text-sm font-medium text-gray-900 dark:text-gray-300"></span>
                                                            </label>
                                                            :
                                                            <label class="relative inline-flex items-center cursor-pointer"
                                                                onClick={e => handleRegisterHotel(e, hotel)}>
                                                                <input type="checkbox" checked={false} class="sr-only peer" />
                                                                <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-indigo-300 dark:peer-focus:ring-indigo-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-indigo-600"></div>
                                                                <span class="text-sm font-medium text-gray-900 dark:text-gray-300"></span>
                                                            </label>
                                                        }
                                                    </div>
                                                </div>

                                            </div>
                                        </li>

                                    </>
                                ))}
                            </ul>
                        </section>

                    </form>
                </div>
            </div>

        </>
    )
    //return (<div>chau</div>)
}