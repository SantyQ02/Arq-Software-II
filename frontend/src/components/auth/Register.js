import { register } from "@/lib/api/auth";
import { useEffect, useState } from "react";
import { useContext } from 'react';
import Banner from "./Banner";
import {Oval} from 'react-loader-spinner'


export default function Register() {
        const [formData, setFormData] = useState({
          first_name: '',
          last_name: '',
          email: '',
          user_name: '',
          password: '',
          password_confirm: ''
        })
      
        const {
            first_name,
            last_name,
            email,
            user_name,
            password,
            password_confirm
        } = formData;
      
        const onChange = e => setFormData({ ...formData, [e.target.name]: e.target.value });

        const [loading, setLoading] = useState(false)
        const onSubmit = async e => {
          e.preventDefault();
          setLoading(true)
          await register(first_name, last_name, email, user_name, password, password_confirm)
          setLoading(false)
          //setHotel(newHotel)
        }


    return (
        <>
            <div className="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
                <div className="sm:mx-auto sm:w-full sm:max-w-sm">
                    <img
                        className="mx-auto h-10 w-auto"
                        src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600"
                        alt="Your Company"
                    />
                    <h2 className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">
                        Register Now
                    </h2>
                </div>

                <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
                    <form onSubmit={e => onSubmit(e)} className="space-y-6" action="#" method="POST">

                        <div>
                            <label htmlFor="text" className="block text-sm font-medium leading-6 text-gray-900">
                                First Name
                            </label>
                            <div className="mt-2">
                                <input
                                    id="first_name"
                                    name="first_name"
                                    value={first_name}
                                    onChange={e => onChange(e)}
                                    type="text"
                                    autoComplete="first_name"
                                    required
                                    className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6 px-2"
                                />
                            </div>
                        </div>

                        <div>
                            <label htmlFor="text" className="block text-sm font-medium leading-6 text-gray-900">
                                Last Name
                            </label>
                            <div className="mt-2">
                                <input
                                    id="last_name"
                                    name="last_name"
                                    value={last_name}
                                    onChange={e => onChange(e)}
                                    type="text"
                                    autoComplete="last_name"
                                    required
                                    className="px-2 block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                />
                            </div>
                        </div>

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
                                    className="px-2 block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                />
                            </div>
                        </div>

                        <div>
                            <label htmlFor="username" className="block text-sm font-medium leading-6 text-gray-900">
                                User Name
                            </label>
                            <div className="mt-2">
                                <input
                                    id="user_name"
                                    name="user_name"
                                    value={user_name}
                                    onChange={e => onChange(e)}
                                    type="username"
                                    autoComplete="user_name"
                                    required
                                    className="px-2 block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                />
                            </div>
                        </div>

                        <div>
                            <div className="flex items-center justify-between">
                                <label htmlFor="password" className="block text-sm font-medium leading-6 text-gray-900">
                                    Password
                                </label>
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
                                    className="px-2 block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                />
                            </div>
                        </div>

                        <div>
                            <div className="flex items-center justify-between">
                                <label htmlFor="password" className="block text-sm font-medium leading-6 text-gray-900">
                                    Confirm Password
                                </label>
                            </div>
                            <div className="mt-2">
                                <input
                                    id="password_confirm"
                                    name="password_confirm"
                                    value={password_confirm}
                                    onChange={e => onChange(e)}
                                    type="password"
                                    autoComplete="current-password"
                                    required
                                    className="px-2 block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
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
                                Register
                            </button>
                            }
                            
                        </div>
                    </form>

                </div>
            </div>
        </>
    )
}
