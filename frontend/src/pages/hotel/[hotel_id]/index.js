import HotelDetail from '@/components/hotel/HotelDetail'
import MainLayout from '@/layouts/MainLayout'

export async function getServerSideProps({query}){
  return {
        props: {
          hotel: mockHotel,
        },
      };
}

export default function Hotel({hotel}) {


  return (
    <>
      <MainLayout title={"Hotel Detail"}>
      <div className='mx-auto max-w-7xl px-4 sm:static sm:px-6 lg:px-8 my-20'>
        <HotelDetail hotel={hotel}/>
      </div>

      </MainLayout>
    </>
  )
}

const mockHotel = {
  "hotel_id": "2a410924-4805-40b8-b96b-5806d6fc3bea",
  "amadeus_id": "USH259",
  "city_code": "USH",
  "title": "Hotel 2",
  "description": "Este es un hotel de prueba en la ciudad de Nueva York.",
  "price_per_day": 152,
  "photos": [
    {"photo_id":"1", "url":"https://static.independent.co.uk/2023/03/24/09/Best%20New%20York%20boutique%20hotels.jpg"},
    {"photo_id":"2", "url":"https://static.independent.co.uk/2023/03/24/09/Best%20New%20York%20boutique%20hotels.jpg"}
  ],
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
