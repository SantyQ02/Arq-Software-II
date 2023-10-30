/*
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
import { Fragment, useContext, useEffect, useState } from 'react'
import { Dialog, RadioGroup, Transition, Tab, Disclosure } from '@headlessui/react'
import { XIcon } from '@heroicons/react/outline'
import { StarIcon } from '@heroicons/react/solid'
import Image from 'next/image'
import BookForm from './BookFormCustom'
// import SuggestedHotels from './SuggestedHotels'

function classNames(...classes) {
    return classes.filter(Boolean).join(' ')
}

export default function HotelDetail({ hotel}) {

    const [dataCheck, setDataCheck] = useState(null)

    return (
        <>
            {hotel !== null && hotel !== undefined ?


                <div className="grid w-full grid-cols-1 items-start gap-x-6 gap-y-8 sm:grid-cols-1 lg:gap-x-8 lg:grid-cols-2">

                    <div className="sm:col-span-8 lg:col-span-1 order-1">
                        <h2 className="text-2xl font-bold text-gray-900 sm:pr-12">{hotel.title.toUpperCase()}</h2>

                        <section aria-labelledby="information-heading" className="mt-2">
                            <h3 id="information-heading" className="sr-only">
                                Hotel information
                            </h3>

                            <p className="text-2xl text-gray-900">$ {hotel.price_per_day}</p>


                        </section>

                        <section aria-labelledby="options-heading" className="mt-10">
                            <h3 id="options-heading" className="sr-only">
                                Product options
                            </h3>

                            <BookForm hotel={hotel} setDataCheck={setDataCheck} />

                            <div className="mt-10">
                                <div className="flex items-center justify-between">
                                    <h4 className="text-sm font-medium text-gray-900">Amenities</h4>

                                </div>

                                <RadioGroup className="mt-4">
                                    <RadioGroup.Label className="sr-only">Choose a size</RadioGroup.Label>
                                    <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
                                        {hotel.amenities && hotel.amenities.map((amenity) => (
                                            <RadioGroup.Option
                                                key={amenity.amenitie_id}
                                                value={amenity.title}
                                                disabled={!amenity}
                                                className={({ active }) =>
                                                    classNames(
                                                        amenity
                                                            ? ' bg-white text-gray-900 shadow-sm'
                                                            : 'cursor-not-allowed bg-gray-50 text-gray-200',
                                                        'group relative flex items-center justify-center rounded-md border py-3 px-4 text-sm font-medium uppercase hover:bg-gray-50  sm:flex-1'
                                                    )
                                                }
                                            >

                                                <RadioGroup.Label as="span">{amenity.title.toUpperCase()}</RadioGroup.Label>
                                                {amenity ? (
                                                    <span
                                                        className={classNames(
                                                        )}
                                                        aria-hidden="true"
                                                    />
                                                ) : (
                                                    <>
                                                    </>
                                                )}

                                            </RadioGroup.Option>
                                        ))}
                                    </div>
                                </RadioGroup>
                            </div>
                        </section>
                    </div>
                    <Tab.Group as="div" className="flex flex-col-reverse order-2 md:order-first">
                        {/* Image selector */}
                        <div className=" mt-6 w-full max-w-2xl mx-auto sm:block lg:max-w-none">
                            <Tab.List className="grid grid-cols-4 gap-6">
                                {hotel.photos != null && hotel.photos.map((photo) => (
                                    <Tab
                                        key={photo.photo_id}
                                        className="relative h-24 bg-white rounded-md flex items-center justify-center text-sm font-medium uppercase text-gray-900 cursor-pointer hover:bg-gray-50 focus:outline-none focus:ring focus:ring-offset-4 focus:ring-opacity-50"
                                    >
                                        {({ selected }) => (
                                            <>
                                                <span className="sr-only">{photo.url}</span>
                                                <span className="absolute inset-0 rounded-md overflow-hidden">
                                                    <Image 
                                                    src={`${photo.url}`} 
                                                    // src={`${process.env.NEXT_PUBLIC_URL_API}/api/public/${photo.url}`} 
                                                    alt="" 
                                                    className=" object-center object-cover" 
                                                    fill />
                                                </span>
                                                <span
                                                    className={classNames(
                                                        selected ? 'ring-indigo-500' : 'ring-transparent',
                                                        'absolute inset-0 rounded-md ring-2 ring-offset-2 pointer-events-none'
                                                    )}
                                                    aria-hidden="true"
                                                />
                                            </>
                                        )}
                                    </Tab>
                                ))}
                            </Tab.List>
                        </div>

                        <Tab.Panels className="w-full aspect-w-1 aspect-h-1">
                            {hotel.photos != null && hotel.photos.map((photo) => (
                                <Tab.Panel key={photo.photo_id}>
                                    <Image
                                        src={`${photo.url}`}
                                        // src={`${process.env.NEXT_PUBLIC_URL_API}/api/public/${photo.url}`}
                                        alt={photo.url}
                                        className="w-full object-center object-cover sm:rounded-lg h-96"
                                        width={1000}
                                        height={1000}
                                    // style={{ height: "30rem" }}
                                    />
                                </Tab.Panel>
                            ))}
                        </Tab.Panels>
                    </Tab.Group>
                    {/* aca va la lista de sugerencias */}
                    {/* <div className='sm:col-span-8 lg:col-span-2 order-3'>
                                                    {dataCheck !== null &&
                                                        <SuggestedHotels dataCheck={dataCheck} setHotel={setHotel} />
                                                    }
                                                </div> */}
                    {/* aca finaliza la lista de sugerencias */}

                </div>

                :
                <></>
            }
        </>

    )
}
