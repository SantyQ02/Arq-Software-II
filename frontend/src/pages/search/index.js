
import HotelsList from '@/components/search/HotelListCustom'
import SearchForm from '@/components/home/new/SearchForm'
import MainLayout from '@/layouts/MainLayout'
import { useRouter } from 'next/router'

export async function getServerSideProps({query}){
  return {
        props: {
          hotels: mockHotels,
        },
      };
}

export default function Search({hotels}) {

  const router = useRouter()

  return (
    <>
      <MainLayout title={"Search"}>
        <div className='mx-auto max-w-7xl px-4 sm:static sm:px-6 lg:px-8 mt-10'>

        <h1 className="text-2xl font-extrabold tracking-tight text-gray-900 sm:text-3xl text-center">You Searched For:</h1>

          <SearchForm
            className='mt-20'
            city_code={router.query.city_code}
            check_in_date={router.query.check_in_date}
            check_out_date={router.query.check_out_date}
          />

          <HotelsList hotels={hotels} />
        </div>

      </MainLayout>
    </>
  )
}

const mockHotels = [
  {
    "hotel_id": "2a410924-4805-40b8-b96b-5806d6fc3bea",
    "amadeus_id": "USH259",
    "city_code": "USH",
    "title": "Hotel 2",
    "description": "Este es un hotel de prueba en la ciudad de Nueva York.",
    "price_per_day": 152,
    "photos": null,
    "thumbnail": "https://static.independent.co.uk/2023/03/24/09/Best%20New%20York%20boutique%20hotels.jpg",
    "amenities": [
      {
        "amenity_id": "249a2834-fe5c-47fa-9083-eeb2f8843783",
        "title": "Piscina"
      },
      {
        "amenity_id": "534cbf4c-ec22-447a-b6e3-be607dce28ad",
        "title": "Gimnasio"
      }
    ],
    "active": true
  },
  {
    "hotel_id": "2a410924-4805-40b8-b96b-5806d6fc3bea",
    "amadeus_id": "USH259",
    "city_code": "USH",
    "title": "Hotel 2",
    "description": "Este es un hotel de prueba en la ciudad de Nueva York.",
    "price_per_day": 152,
    "photos": null,
    "thumbnail": "https://images.bubbleup.com/width1920/quality35/mville2017/1-brand/1-margaritaville.com/gallery-media/220803-compasshotel-medford-pool-73868-1677873697.jpg",
    "amenities": [
      {
        "amenity_id": "249a2834-fe5c-47fa-9083-eeb2f8843783",
        "title": "Piscina"
      },
      {
        "amenity_id": "534cbf4c-ec22-447a-b6e3-be607dce28ad",
        "title": "Gimnasio"
      }
    ],
    "active": true
  }
]