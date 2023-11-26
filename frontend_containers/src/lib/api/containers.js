import { alert } from "../utils/alert";

const { default: axios } = require("axios");

export async function getContainersStats() {

    const config = {
        headers: {
            'Cache-Control': 'no-cache'
        }
    }

    try {
        const res = await axios.get('/api/containers', config)
        if (res.status === 200) {
            return res.data.containers_stats
        }
        else {
            return null
        }
    } catch (error) {
        console.log(error)
        return null
    }


}

export async function restartContainer(container_id) {

    const config = {
        headers: {
            'Cache-Control': 'no-cache'
        }
    }

    try {
        const res = await axios.get(`/api/containers/restart/${container_id}`, config)
        if (res.status === 200) {
            alert('success', 'Container restarted successfully')
            return res.data.success
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

export async function startContainer(container_id) {

    const config = {
        headers: {
            'Cache-Control': 'no-cache'
        }
    }

    try {
        const res = await axios.get(`/api/containers/start/${container_id}`, config)
        if (res.status === 200) {
            alert('success', 'Container started successfully')
            return res.data.success
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

export async function stopContainer(container_id) {

    const config = {
        headers: {
            'Cache-Control': 'no-cache'
        }
    }

    try {
        const res = await axios.get(`/api/containers/stop/${container_id}`, config)
        if (res.status === 200) {
            alert('success', 'Container stopped successfully')
            return res.data.success
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

export async function deleteContainer(container_id) {

    const config = {
        headers: {
            'Cache-Control': 'no-cache'
        }
    }

    try {
        const res = await axios.delete(`/api/containers/delete/${container_id}`, config)
        if (res.status === 200) {
            alert('success', 'Container deleted successfully')
            return res.data.success
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

export async function createContainer(service, quantity = 0) {
    const config = {
        headers: {
            'Content-Type': 'application/json'
        },
    };

    const body = JSON.stringify({
        service,
        quantity
    });

    try {
        const res = await axios.post(`/api/containers`, body, config)
        if (res.status === 201) {
            alert('success', 'Container created successfully')

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