export default async function handler(req, res) {
    try {
        const { amenitieID } = req.query;
        //   const response = await fetch(`${process.env.NEXT_PUBLIC_URL_API}/api/hotel`, {
        const response = await fetch(`${process.env.NEXT_PUBLIC_URL_API}/api/hotel/amenitie/${amenitieID}`, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json',
                'Cookie': req.headers.cookie || '',
                'Cache-Control': 'no-cache'
            },
            credentials: 'include',
        });

        const data = await response.json();
        res.status(response.status).json(data);
    } catch (error) {
        res.status(500).json({ error: 'Internal Server Error' });
    }
}