import { verifyemail } from "@/lib/api/auth";
import Link from 'next/link'
import { useRouter } from 'next/router';


export default function VerifyEmail() {
        
    const router = useRouter();
    const { verificationCode } = router.query; // Obtiene el valor del parámetro dinámico "verification_code"

    const handleVerifyEmail = () => {
        verifyemail(verificationCode);
        router.push('/auth/login');
      };

    return (
        <>
        <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
      <button className="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600" 
       onClick={handleVerifyEmail}>Verify Email</button>
       </div>
    </>
    )
}