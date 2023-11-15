import { alert } from "../utils/alert";
const { default: axios } = require("axios");

export async function createBooking(rooms, total, date_in, date_out, hotel_id, user_id, hotel_title) {

    rooms = parseInt(rooms)
    total = parseFloat(total)

    const config = {
        headers: {
            'Content-Type': 'application/json'
        },
        withCredentials: true
    };

    const body = JSON.stringify({
        rooms,
        total,
        date_in,
        date_out,
        hotel_id,
        user_id,
        hotel_title
    });

    try {
        const res = await axios.post('/api/booking', body, config)
        if (res.status === 201) {
            alert('success', 'Booking Created')
            return true
        
        }
        else {
            //console.log("res: " + res)
            alert('error', res.data.error.toString())
            return false
        }
    } catch (error) {
        const errorMessage = error.response?.data?.error ?? 'Unknown error occurred';
        //console.log(error)
        alert('error', String(errorMessage));
        return false
    }

}

export async function searchBookingsMe(hotel, date_in, date_out){
    const config = {
        headers:{
            'Cache-Control': 'no-cache'
        }
    }

    try {
        const res = await axios.get(`/api/booking/search/me?hotel=${hotel}&date_in=${date_in}&date_out=${date_out}`,{withCredentials: true})
        if (res.status === 200) {
            return res.data.bookings
        }
        else {
            return null
        }
    } catch (error) {
        return null
    }
}

export async function searchBookings(hotel,user, date_in, date_out){
    const config = {
        headers:{
            'Cache-Control': 'no-cache'
        }
    }

    try {
        const res = await axios.get(`/api/booking/search?hotel=${hotel}&user=${user}&date_in=${date_in}&date_out=${date_out}`,{withCredentials: true})
        if (res.status === 200) {
            return res.data.bookings
        }
        else {
            return null
        }
    } catch (error) {
        return null
    }
}

export async function setActiveBooking(booking_id, active){
    const config = {
        headers: {
            'Content-Type': 'application/json'
        },
        withCredentials: true
    };

    const body = JSON.stringify({
        active
    });

    try {
        const res = await axios.put(`/api/booking/update/${booking_id}`, body, config)
        if (res.status === 200) {
            alert('success', 'Booking updated successfully')

        }
        else {
            //console.log("res: " + res)
            alert('error', res.data.error.toString())
        }
    } catch (error) {
        const errorMessage = error.response?.data?.error ?? 'Unknown error occurred';
        alert('error', String(errorMessage));
    }
}