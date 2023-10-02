import { alert } from "../utils/alert";

const { default: axios } = require("axios");

export async function getMe() {

    const config = {
        headers: {
            'Cache-Control': 'no-cache'
        }
    }

    try {
        const res = await axios.get('/api/user/me', { withCredentials: true })
        if (res.status === 200) {
            return res.data.user
        }
        else {
            return null
        }
    } catch (error) {
        return null
    }


}

export async function getUserById(user_id) {

    const config = {
        headers: {
            'Cache-Control': 'no-cache'
        }
    }

    try {
        const res = await axios.get(`/api/user/${user_id}`, { withCredentials: true })
        if (res.status === 200) {
            return res.data.user
        }
        else {
            return null
        }
    } catch (error) {
        return null
    }


}

export async function getUsers() {

    const config = {
        headers: {
            'Cache-Control': 'no-cache'
        }
    }

    try {
        const res = await axios.get('/api/user/', { withCredentials: true })
        if (res.status === 200) {
            return res.data.users
        }
        else {
            return null
        }
    } catch (error) {
        return null
    }


}

export async function updateUser(user_id, first_name, last_name, user_name, role, active) {
    const config = {
        headers: {
            'Content-Type': 'application/json'
        },
        withCredentials: true
    };

    const body = JSON.stringify({
        first_name,
        last_name,
        user_name,
        role,
        active
    });

    try {
        const res = await axios.put(`/api/user/update/${user_id}`, body, config)
        if (res.status === 200) {
            alert('success', 'User updated successfully')

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