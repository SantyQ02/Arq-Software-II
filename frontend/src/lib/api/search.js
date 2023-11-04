import axios from "axios";

export async function search_hotels(city_code, check_in_date, check_out_date) {

    try {
        const res = await axios.get(`http://localhost:${process.env.PORT}/api/search?city_code=${city_code}&check_in_date=${check_in_date}&check_out_date=${check_out_date}`)
        if (res.status != 200) {
            return []
        }
        const hotels = res.data.hotels
        return hotels
    } catch (error) {
        const errorMessage = error.response?.data?.error ?? 'Unknown error occurred';
        console.log(error)
        return []
    }
}