import { alert } from "../utils/alert";
import Cookies from "js-cookie";
const { default: axios } = require("axios");

export async function login(email, password) {

    const config = {
        headers: {
            'Content-Type': 'application/json'
        }
    };

    const body = JSON.stringify({
        email,
        password
    });

    try {
        const res = await axios.post('/api/auth/login', body, config)
        if (res.status === 200) {
            Cookies.set("token", res.data.token.toString());
            alert('success', 'Login successfully')

        }
        else{
            //console.log("res: " + res)
            alert('error', res.data.error.toString())
            Cookies.remove("token");
        }
    } catch (error) {
        const errorMessage = error.response?.data?.error ?? 'Unknown error occurred';
        alert('error', String(errorMessage));
        Cookies.remove("token");
    }

}

export async function logout() {

    try {
        const res = await axios.get('/api/auth/logout',{withCredentials: true})
        if (res.status === 200) {
            Cookies.remove("token");
            alert('success', String(res.data.success))

        }
        else{
            //console.log("res: " + res)
            alert('error', res.data.error.toString())
            Cookies.remove("token");
        }
    } catch (error) {
        const errorMessage = error.response?.data?.error ?? 'Unknown error occurred';
        alert('error', String(errorMessage));
        Cookies.remove("token");
    }

}

export async function refresh() {

    try {
        const res = await axios.get('/api/auth/refresh',{withCredentials: true})
        if (res.status === 200) {
            Cookies.set("token", res.data.token.toString());
        }
        else{
            //console.log("res: " + res)
            Cookies.remove("token");
        }
    } catch (error) {
        const errorMessage = error.response?.data?.error ?? 'Unknown error occurred';
        Cookies.remove("token");
    }

}

export async function register(first_name, last_name, email, user_name, password, password_confirm) {

    const config = {
        headers: {
            'Content-Type': 'application/json'
        }
    };

    const body = JSON.stringify({
        first_name,
        last_name,
        email,
        user_name,
        password,
        password_confirm,
    });

    try {
        const res = await axios.post('/api/auth/register', body, config)
        if (res.status === 201) {
            alert('success', 'We sent you an email with a verification code')

        }
        else{
            //console.log("res: " + res)
            alert('error', res.data.error.toString())
        }
    } catch (error) {
        const errorMessage = error.response?.data?.error ?? 'Unknown error occurred';
        alert('error', String(errorMessage));
    }

}

export async function verifyemail(verificationCode) {
    try {
        const res = await axios.get(`/api/auth/${verificationCode}`)
        if (res.status === 200) {
            alert('success', 'Verification successfully')

        }
        else{
            //console.log("res: " + res)
            alert('error', res.data.error.toString())
        }
    } catch (error) {
        const errorMessage = error.response?.data?.error ?? 'Unknown error occurred';
        alert('error', String(errorMessage));
    }

}

export async function resetPassword(email){
    const config = {
        headers: {
            'Content-Type': 'application/json'
        }
    };

    const body = JSON.stringify({
        email
    });

    try {
        const res = await axios.post('/api/auth/reset', body, config)
        if (res.status === 200) {
            alert('success', 'We send you an email to recovery your password')

        }
        else{
            //console.log("res: " + res)
            alert('error', res.data.error.toString())
        }
    } catch (error) {
        const errorMessage = error.response?.data?.error ?? 'Unknown error occurred';
        alert('error', String(errorMessage));
    }
}

export async function resetPasswordConfirm(verificationCode, password, password_confirm){
    const config = {
        headers: {
            'Content-Type': 'application/json'
        }
    };

    const body = JSON.stringify({
        password,
        password_confirm
    });

    try {
        const res = await axios.post(`/api/auth/reset/${verificationCode}`, body, config)
        if (res.status === 200) {
            alert('success', 'Successful password recovery')

        }
        else{
            //console.log("res: " + res)
            alert('error', res.data.error.toString())
        }
    } catch (error) {
        const errorMessage = error.response?.data?.error ?? 'Unknown error occurred';
        alert('error', String(errorMessage));
    }
}