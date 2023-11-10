import MainLayout from "@/layouts/MainLayout"
import { useRouter } from "next/router"
import { useEffect } from "react"

const ThankYou = () => {

    const router = useRouter()
    const status = router.query.status
    return(
        <MainLayout title={status == 'success' ? 'Success' : 'Error'}>
            <div className="bg-white">
            <div className="max-w-7xl mx-auto py-16 px-4 sm:py-24 sm:px-6 lg:px-8">
                <div className="text-center">
                <p className="mt-1 text-4xl font-extrabold text-gray-900 sm:text-5xl sm:tracking-tight lg:text-6xl">
                    {status == 'success' ? 'Thank You' : 'Sorry'}
                </p>
                <p className="max-w-xl mt-5 mx-auto text-xl text-gray-500">
                    {status == 'success' ? 
                    'Hope you enjoyed shopping in our web.' 
                    : 
                    'An error has occurred, try again.'}
                </p>
                </div>
            </div>
            </div>
            </MainLayout>
    )
}

export default ThankYou