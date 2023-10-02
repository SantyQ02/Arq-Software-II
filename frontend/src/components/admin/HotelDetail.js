/*
  This example requires Tailwind CSS v2.0+ 
  
  This example requires some changes to your config:
  
  ```
  // tailwind.config.js
  module.exports = {
    // ...
    plugins: [
      // ...
      require('@tailwindcss/typography'),
      require('@tailwindcss/aspect-ratio'),
    ],
  }
  ```
*/
import { useState } from 'react'
import { associateAmenities, deletePhoto, getAmenities, getHotelById, insertPhoto, updateHotel } from "@/lib/api/hotel";
import { Disclosure, RadioGroup, Tab } from '@headlessui/react'
import { StarIcon } from '@heroicons/react/solid'
import { HeartIcon, MinusSmIcon, PlusSmIcon, TrashIcon } from '@heroicons/react/outline'
import { useContext, useEffect } from "react"
import Image from 'next/image';
import { dissociateAmenities } from '@/lib/api/hotel';


function classNames(...classes) {
  return classes.filter(Boolean).join(' ')
}

export default function HotelDetail({ hotel, setHotel }) {

  const get_hotel_by_id = async () => {

    const updatedHotel = await getHotelById(hotel.hotel_id)
    const data = await getAmenities()
    const filtered_data = data?.filter(amenity => !updatedHotel.amenities?.some(hotelAmenity => hotelAmenity.amenitie_id === amenity.amenitie_id));
    setAmenities(filtered_data)
    setHotel(updatedHotel)
  }

  // Get Amenities
  const [amenities, setAmenities] = useState(null)
  const [selectedAmenities, setSelectedAmenities] = useState([]);

  const get_amenities = async () => {
    const data = await getAmenities()
    const filtered_data = data?.filter(amenity => !hotel.amenities?.some(hotelAmenity => hotelAmenity.amenitie_id === amenity.amenitie_id));
    setAmenities(filtered_data)
  }

  // const [change, setChange] = useState(false)
  useEffect(() => {
    get_amenities()
  }, [])

  //Add amenity
  const handleAddAmenity = async amenity_id => {
    const amenities_array = [amenity_id]

    await associateAmenities(hotel.hotel_id, amenities_array)
    get_hotel_by_id()
    const filtered_amenities = amenities?.filter(amenity => amenity.amenitie_id !== amenity_id);

    setAmenities(filtered_amenities)

  }

  //Remove amenity
  const handleRemoveAmenity = async amenity_id => {
    const amenities_array = [amenity_id]

    await dissociateAmenities(hotel.hotel_id, amenities_array)
    get_hotel_by_id()
    // const filtered_amenities = amenities?.filter(amenity => amenity.amenitie_id !== amenity_id);

    // setAmenities(filtered_amenities)


  }

  // Add File
  const [file, setFile] = useState(null);

  const handleChangeFile = (e) => {
    const addFile = e.target.files[0];
    setFile(addFile);
    if (addFile) {
      const reader = new FileReader();

      reader.onload = function (e) {
        const imagePreview = document.getElementById('image-preview');
        imagePreview.src = e.target.result;
      };

      reader.readAsDataURL(addFile);
    }
  };



  const handleSubmitFile = async (e) => {
    e.preventDefault();

    if (file) {
      await insertPhoto(hotel.hotel_id, file)
    }
    setFile(null)

    get_hotel_by_id()
  };
  // end add file

  //Delete photo
  const handleDeletePhoto = async (e, photo_id) => {
    e.preventDefault()
    await deletePhoto(photo_id)
    get_hotel_by_id()

  }
  //end delete photo

  const [selectedSize, setSelectedSize] = useState(null);
  const [editableFields, setEditableFields] = useState({
    title: hotel.title,
    price_per_day: hotel.price_per_day,
    rooms: hotel.rooms,
    description: hotel.description,
  });

  const handleFieldChange = (e) => {
    const { name, value } = e.target;
    setEditableFields((prevFields) => ({
      ...prevFields,
      [name]: value,
    }));
  };

  const handleEditClick = async () => {
    const { title, price_per_day, rooms, description } = editableFields;
    const updatedHotel = await updateHotel(hotel.hotel_id, title, description, price_per_day, rooms, hotel.active);

    if (updatedHotel) {
      // Realizar acciones adicionales después de la actualización exitosa
      //console.log('Hotel actualizado:', updatedHotel);
    } else {
      // Manejar el error de actualización
      console.error('Error al actualizar el hotel');
    }

    //console.log(editableFields);
  };


  return (
    <div className="bg-white">
      <div className="max-w-2xl mx-auto py-3 px-4 sm:py-3 sm:px-6 lg:max-w-7xl lg:px-8">
        <div className="lg:grid lg:grid-cols-2 lg:gap-x-8 lg:items-start">
          {/* photo gallery */}
          <Tab.Group as="div" className="flex flex-col-reverse">
            <div className=" mt-6 w-full max-w-2xl mx-auto sm:block lg:max-w-none">
              <Tab.List className="grid grid-cols-4 gap-6">
                {hotel != null && hotel.photos != null && hotel.photos.map((photo) => (
                  <Tab
                    key={photo.photo_id}
                    className="relative h-24 bg-white rounded-md flex items-center justify-center text-sm font-medium uppercase text-gray-900 cursor-pointer hover:bg-gray-50 focus:outline-none focus:ring focus:ring-offset-4 focus:ring-opacity-50"
                  >
                    {({ selected }) => (
                      <>
                        <span className="absolute inset-0 rounded-md overflow-hidden">
                          <Image src={`${process.env.NEXT_PUBLIC_URL_API}/api/public/${photo.url}`} alt="" className="w-full h-full object-center object-cover" width={1000} height={1000}/>
                        </span>
                        <button className='z-50 absolute end-0 top-0'
                          onClick={e => handleDeletePhoto(e, photo.photo_id)}>

                          <TrashIcon className='h-7 w-7 text-red-400' />
                        </button>
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
              {hotel != null && hotel.photos != null && hotel.photos.map((photo) => (
                <Tab.Panel key={photo.photo_id}>
                  <Image
                    src={`${process.env.NEXT_PUBLIC_URL_API}/api/public/${photo.url}`}
                    className="w-full object-center object-cover sm:rounded-lg h-96"
                    width={1000} height={1000}
                  />
                </Tab.Panel>
              ))}
            </Tab.Panels>
          </Tab.Group>

          {/* hotel info */}
          <div className="mt-10  sm:px-0 sm:mt-16 lg:mt-0">
            <h1 className="text-3xl font-extrabold tracking-tight text-gray-900">
              <input
                name="title"
                value={editableFields.title}
                onChange={handleFieldChange}
                className="w-full border-gray-300 rounded-md  py-2"
              />
            </h1>

            <div className="mt-3">
              <h2 className="text-base text-gray-700  space-y-6">
                Price per day:{' '}
                <input
                type="number"
                  name="price_per_day"
                  value={editableFields.price_per_day}
                  onChange={handleFieldChange}
                  className="w-full border-gray-300 rounded-md  py-2"
                />
              </h2>
            </div>

            <div className="mt-3">
              <h2 className="text-base text-gray-700 space-y-6">
                Rooms:{' '}
                <input
                  name="rooms"
                  type="number"
                  value={editableFields.rooms}
                  onChange={handleFieldChange}
                  className="w-full border-gray-300 rounded-md py-2"
                />
              </h2>
            </div>

            <div className="mt-3">
              <h3 className="text-base text-gray-700 space-y-6">Description:</h3>

              <textarea
                name="description"
                value={editableFields.description}
                onChange={handleFieldChange}
                className="w-full border-gray-300 rounded-md py-2 text-gray-700"
              />
            </div>

            <button
              className=" inline-block rounded-md border border-transparent bg-indigo-600 px-3 md:px-8 py-3 text-center font-medium text-white hover:bg-indigo-700 ml-auto"
              onClick={handleEditClick}
            >
              Update Changes
            </button>

            <section aria-labelledby="details-heading" className="mt-12">
              <h2 id="details-heading" className="text-3xl text-gray-900">
                Amenities:
              </h2>
              <RadioGroup value={selectedSize} onChange={setSelectedSize} className="mt-4">
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
                      <button className='flex items-center' onClick={e => handleRemoveAmenity(amenity.amenitie_id)}>
                          
                        <TrashIcon className='h-5 w-5 text-red-400' />
                        <RadioGroup.Label as="span">{amenity.title.toUpperCase()}</RadioGroup.Label>
                      </button>
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


            </section>
          </div>

          {/* add photos */}
          <>
            <div className="mt-5 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6">

              <form className="col-span-full" onSubmit={e => handleSubmitFile(e)}>
                <label htmlFor="cover-photo" className="block text-sm font-medium leading-6 text-gray-900">
                  Hotel photos
                </label>
                <div className="mt-2 flex justify-center rounded-lg border border-dashed border-gray-900/25 px-6 py-10">
                  <div className="text-center">
                    {/* <PhotoIcon className="mx-auto h-12 w-12 text-gray-300" aria-hidden="true" /> */}
                    {file === null || file === undefined ?
                      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-6 h-6 mx-auto h-12 w-12 text-gray-300" aria-hidden="true">
                        <path fillRule="evenodd" d="M1.5 6a2.25 2.25 0 012.25-2.25h16.5A2.25 2.25 0 0122.5 6v12a2.25 2.25 0 01-2.25 2.25H3.75A2.25 2.25 0 011.5 18V6zM3 16.06V18c0 .414.336.75.75.75h16.5A.75.75 0 0021 18v-1.94l-2.69-2.689a1.5 1.5 0 00-2.12 0l-.88.879.97.97a.75.75 0 11-1.06 1.06l-5.16-5.159a1.5 1.5 0 00-2.12 0L3 16.061zm10.125-7.81a1.125 1.125 0 112.25 0 1.125 1.125 0 01-2.25 0z" clipRule="evenodd" />
                      </svg>
                      :

                      <div className="flex justify-center">

                        <Image src="" alt="Preview" id="image-preview" height={100} width={100} />
                      </div>
                    }

                    <div className="mt-4 flex text-sm leading-6 text-gray-600">
                      <label
                        htmlFor="file"
                        className="relative cursor-pointer rounded-md bg-white font-semibold text-indigo-600  hover:text-indigo-500"
                      >
                        <span>Upload a image</span>
                        <input
                          id="file"
                          name="file"
                          onChange={e => handleChangeFile(e)}
                          type="file"
                          className="sr-only"
                          accept="image/*"
                          multiple
                        />
                      </label>
                      <p className="pl-1">or drag and drop</p>
                    </div>
                    <p className="text-xs leading-5 text-gray-600">PNG, JPG, JPEG, WEBP, SVG up to 10MB</p>
                  </div>
                </div>
                <div className="mt-6 flex items-center justify-end gap-x-6">
                  <button
                    type="submit"
                    className="rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                  >
                    Add Photo
                  </button>
                </div>
              </form>

            </div>
          </>

          {/* add amenitie */}
          <>
            <section aria-labelledby="details-heading" className="mt-5">
              <h2 id="details-heading" className="block text-sm font-medium leading-6 text-gray-900">
                Add amenities:
              </h2>
              <RadioGroup value={selectedSize} onChange={setSelectedSize} className="mt-4">
                <RadioGroup.Label className="sr-only">Choose a size</RadioGroup.Label>
                <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
                  {amenities && amenities.map((amenity) => (
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
                      <button onClick={e => handleAddAmenity(amenity.amenitie_id)} className='flex items-center'>
                        <PlusSmIcon className='h-5 w-5 text-gray-400' />
                        <RadioGroup.Label as="span">{amenity.title.toUpperCase()}</RadioGroup.Label>
                      </button>
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


            </section>
          </>
        </div>
      </div>
    </div>
  );
  // return (<div>hola</div>)
}
