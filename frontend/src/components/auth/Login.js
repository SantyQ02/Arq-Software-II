import { login } from "@/lib/api/auth";
import { getMe } from "@/lib/api/user";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import { Navigate } from 'react-router'
import { useContext } from 'react';
import { UserContext } from '../../layouts/LayoutContext';
import Banner from "./Banner";
import { createBooking } from "@/lib/api/booking";
import Link from "next/link";
import {Oval} from 'react-loader-spinner'



/*
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
export default function Login() {
  const [loading, setLoading] = useState(false)


  const [booking, setBooking] = useState(null)

  useEffect(() => {
    const booking_on_storage = JSON.parse(sessionStorage.getItem('booking'))
    setBooking(booking_on_storage)

  }, [])

  const router = useRouter()
  const [navigate, setNavigate] = useState(false)
  const [userPage, setUserPage] = useState(null)

  const [user, setUser] = useContext(UserContext);

  const [formData, setFormData] = useState({
    email: '',
    password: '',
  })

  const {
    email,
    password,
  } = formData;

  const onChange = e => setFormData({ ...formData, [e.target.name]: e.target.value });

  const onSubmit = async e => {
    e.preventDefault();
    setLoading(true)
    await login(email, password);
    const userMe = await getme()

    if (booking != null && userMe != null){
      const {rooms, total, start_date, end_date, hotel_id, hotel_title} = booking
      const create_booking = async () => {
        await createBooking(rooms, total, start_date, end_date, hotel_id, userMe.user_id, hotel_title)
        // //console.log("\nrooms: ",rooms,"\ntotal: ", total,"\ndate_in: ", start_date,"\ndate_out: ", end_date,"\nhotel_id: ", hotel_id,"\nuser_id: ", userMe.user_id)
        sessionStorage.removeItem('booking');

    }
  
    create_booking()
    }
    setLoading(false)
    
  }


  const getme = async () => {
    const userMe = await getMe()
    setUser(userMe)
    setUserPage(userMe)
    return userMe
  }

  if (userPage != null) {
    router.push("/");
  }


  return (
    <>
    {booking !== null && 
      <Banner booking={booking} setBooking={setBooking}/>
    }
      <div className="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
        <div className="sm:mx-auto sm:w-full sm:max-w-sm">
          <img
            className="mx-auto h-10 w-auto"
            src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600"
            alt="Your Company"
          />
          <div className="">
            
          <h2 className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">
            Sign in to your account
          </h2>
          <div className="text-sm flex justify-center"><p>or <span> </span>
                  <Link href="/auth/register" className="font-semibold text-indigo-600 hover:text-indigo-500">
                    register now.
                  </Link>
                  </p>
                </div>
          </div>
        </div>

        <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
          <form onSubmit={e => onSubmit(e)} className="space-y-6" action="#" method="POST">
            <div>
              <label htmlFor="email" className="block text-sm font-medium leading-6 text-gray-900">
                Email address
              </label>
              <div className="mt-2">
                <input
                  id="email"
                  name="email"
                  value={email}
                  onChange={e => onChange(e)}
                  type="email"
                  autoComplete="email"
                  required
                  className="block w-full rounded-md border-0 p-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                />
              </div>
            </div>

            <div>
              <div className="flex items-center justify-between">
                <label htmlFor="password" className="block text-sm font-medium leading-6 text-gray-900">
                  Password
                </label>
                <div className="text-sm">
                  <Link href="/auth/reset-pass" className="font-semibold text-indigo-600 hover:text-indigo-500">
                    Forgot password?
                  </Link>
                </div>
              </div>
              <div className="mt-2">
                <input
                  id="password"
                  name="password"
                  value={password}
                  onChange={e => onChange(e)}
                  type="password"
                  autoComplete="current-password"
                  required
                  className="block w-full rounded-md border-0 p-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                />
              </div>
            </div>

            <div>
            {loading ? 
                              <button
                              className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                            >
                              <Oval
                              type="Oval"
                              color="#fff"
                              width={20}
                              height={20}
                              />
                            </button>
                            :
                            <button
                                type="submit"
                                className="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                            >
                                Sign in
                            </button>
                            }
            </div>
          </form>

        </div>
      </div>
    </>
  )
}
