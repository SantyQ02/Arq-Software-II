import { associateAmenities, createHotel, getAmenities, insertPhoto } from "@/lib/api/hotel";
import Image from "next/image";
import { useEffect, useState } from "react";
import HotelDetail from "./HotelDetail";

export default function CreateHotel() {

  // Create hotel
  const [hotel, setHotel] = useState(null)
  const [formData, setFormData] = useState({
    title: '',
    description: '',
    price: '',
    rooms: ''
  })

  const {
    title,
    description,
    price,
    rooms
  } = formData;

  const onChange = e => setFormData({ ...formData, [e.target.name]: e.target.value });
  const onSubmit = async e => {
    e.preventDefault();
    const newHotel = await createHotel(title, description, price, rooms)
    setHotel(newHotel)
  }
  useEffect(() => {
    // Este efecto secundario se ejecutará cuando el valor de 'hotel' cambie
    //console.log("hotel", hotel);
  }, [hotel]);

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
  };

  // Get Amenities
  const [amenities, setAmenities] = useState(null)
  const [selectedAmenities, setSelectedAmenities] = useState([]);

  useEffect(() => {
    const get_amenities = async () => {
      const data = await getAmenities()
      setAmenities(data)
    }
    get_amenities()
  }, [])

  const handleAmenityChange = (e, amenity_id) => {
    if (e.target.checked) {
      setSelectedAmenities([...selectedAmenities, amenity_id]);
    } else {
      setSelectedAmenities(selectedAmenities.filter((amenity) => amenity !== amenity_id));
    }
  };

  const handleSubmitAmenities = (e) => {
    e.preventDefault();
    const associate_amenities = async () => {
      await associateAmenities(hotel.hotel_id, selectedAmenities)
    }
    associate_amenities()
  };

  return (
    <div className="container">
      <div className="">
        <div className="space-y-12">

          {hotel === null ?
            <div className="border-b border-gray-900/10 pb-12">
              <h2 className="text-base font-semibold leading-7 text-gray-900">Hotel info</h2>
              <p className="mt-1 text-sm leading-6 text-gray-600">
                This information will be displayed publicly.
              </p>

              <form onSubmit={e => onSubmit(e)} className="">
                <div className="mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6">
                  <div className="sm:col-span-4">
                    <label htmlFor="title" className="block text-sm font-medium leading-6 text-gray-900">
                      Name
                    </label>
                    <div className="mt-2">
                      <div className="flex rounded-md shadow-sm ring-1 ring-inset ring-gray-300 focus-within:ring-2 focus-within:ring-inset focus-within:ring-indigo-600 sm:max-w-md">
                        <input
                          type="text"
                          name="title"
                          value={title}
                          onChange={e => onChange(e)}
                          id="title"
                          required
                          autoComplete="title"
                          className="block flex-1 border-0 bg-transparent py-1.5 pl-1 text-gray-900 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6"
                          placeholder="Hotel Title"
                        />
                      </div>
                    </div>
                  </div>

                  <div className="col-span-full">
                    <label htmlFor="description" className="block text-sm font-medium leading-6 text-gray-900">
                      Description
                    </label>
                    <div className="mt-2">
                      <textarea
                        id="description"
                        name="description"
                        value={description}
                        onChange={e => onChange(e)}
                        rows={3}
                        required
                        className="block w-full rounded-md border-0 p-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                        defaultValue={''}
                      />
                    </div>
                    <p className="mt-3 text-sm leading-6 text-gray-600">Write a few sentences about the hotel.</p>
                  </div>

                  <div className="sm:col-span-2 sm:col-start-1">
                    <label htmlFor="price" className="block text-sm font-medium leading-6 text-gray-900">
                      Price per day
                    </label>
                    <div className="mt-2">
                      <input
                        type="number"
                        name="price"
                        value={price}
                        onChange={e => onChange(e)}
                        id="price"
                        required
                        autoComplete="address-level2"
                        className="block w-full rounded-md border-0 p-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                      />
                    </div>
                  </div>

                  <div className="sm:col-span-2">
                    <label htmlFor="rooms" className="block text-sm font-medium leading-6 text-gray-900">
                      Rooms
                    </label>
                    <div className="mt-2">
                      <input
                        type="number"
                        name="rooms"
                        value={rooms}
                        onChange={e => onChange(e)}
                        id="rooms"
                        required
                        autoComplete="address-level1"
                        className="block w-full rounded-md border-0 p-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                      />
                    </div>
                  </div>
                </div>

                <div className="mt-6 flex items-center justify-end gap-x-6">
                  <button
                    type="submit"
                    className="rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                  >
                    Create Hotel
                  </button>
                </div>


              </form>
            </div>
            :
            <>
              <h2 className="text-base font-semibold leading-7 text-gray-900">Hotel Created</h2>
              <p className="mt-1 text-sm leading-6 text-gray-600">
                Now you can add photos and amenities.              </p>
            </>
          }


          {hotel != null ? 
          <HotelDetail hotel={hotel} setHotel={setHotel} />
        :
        <>
              <h2 className="text-base font-semibold leading-7 text-gray-900">Photos & Amenities</h2>
              <p className="mt-1 text-sm leading-6 text-gray-600">
                Available when creating a hotel.              </p>
            </>
            }
        </div>


      </div >
    </div >

  )
}
