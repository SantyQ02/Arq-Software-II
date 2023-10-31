import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import { Oval } from "react-loader-spinner";
import Datepicker from "react-tailwindcss-datepicker";

const SearchForm = ({className="", city_code="Cordoba", check_in_date=null, check_out_date=null}) => {

    const router = useRouter()

    const [loading, setLoading] = useState(false)

    const [city, setCity] = useState(city_code)

    const [dates, setDates] = useState({
        startDate: check_in_date,
        endDate: check_out_date
    });

    useEffect(()=>{
        setLoading(false)
    },[city_code,check_in_date, check_out_date])


    const handleChange = e => setCity(e.target.value);

    const handleDatesChange = (newDates) => {
        setDates(newDates);
    }

    const handleSubmit = e => {
        e.preventDefault()
        setLoading(true)
        const {startDate, endDate} = dates
        router.push(`/search?city_code=${city}&check_in_date=${startDate}&check_out_date=${endDate}`)
        // setLoading(false)
    }

    return (
        <form onSubmit={e => handleSubmit(e)}>
            <div className={`md:flex flex-row items-end justify-between gap-3 mb-10 w-full mt-5 ${className}`}>

                <div className="basis-3/12">
                    <label htmlFor="number" className="block text-sm font-medium leading-6 text-gray-900 pl-4">
                        City
                    </label>
                    <div className="mt-2">
                        <select
                            id="city"
                            name="city"
                            type="text"
                            value={city}
                            onChange={e => handleChange(e)}
                            autoComplete="city"
                            required
                            className="h-full w-full rounded-md border-0 bg-transparent py-2.5 pl-2  text-gray-500 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm relative transition-all duration-300 py-2.5 pl-4  border-r-8 w-full border-white rounded-lg tracking-wide font-light text-sm placeholder-gray-400 bg-white focus:ring disabled:opacity-40 disabled:cursor-not-allowed focus:border-blue-500 focus:ring-blue-500/20 light"
                        >
                            <option value={"Cordoba"}>Cordoba</option>
                            <option value={"Londres"}>Londres</option>
                            <option>3</option>
                            <option>4</option>
                            <option>5</option>
                            <option>6</option>
                            <option>7</option>
                            <option>8</option>
                            <option>9</option>
                            <option>10</option>
                        </select>
                    </div>
                </div>

                <div className="basis-7/12">
                    <label htmlFor="number" className="block text-sm font-medium leading-6 text-gray-900 pl-4">
                        Dates
                    </label>
                    <div className="mt-2 z-20 ">
                        <Datepicker
                            value={dates}
                            onChange={handleDatesChange}
                            inputClassName="relative transition-all duration-300 py-2.5 pl-4 pr-14 w-full border-gray-300 rounded-lg tracking-wide font-light text-sm placeholder-gray-400 bg-white focus:ring disabled:opacity-40 disabled:cursor-not-allowed focus:border-blue-500 focus:ring-blue-500/20 light"
                            className="h-full rounded-md border-0 bg-transparent py-0 pl-2 pr-7 text-gray-500 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm light"

                        />
                    </div>

                </div>

                <button
                    type="submit"
                    disabled={!!!dates.endDate || !!!dates.startDate}
                    className={`mt-6 flex w-full items-center justify-center rounded-md border border-transparent bg-indigo-600 disabled:bg-indigo-400 px-0! py-2 text-base font-medium text-white  hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 basis-2/12 `}
                >
                    {
                        loading ?
                        <Oval
                              type="Oval"
                              color="#fff"
                              width={20}
                              height={20}
                              />
                        :

                    <>
                    Search
                    </>
                    }
                
                </button>

            </div>



        </form>

    );
};

export default SearchForm;