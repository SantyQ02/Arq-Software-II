export default async function handler(req, res) {
    try {
        const { hotelID } = req.query;
        const response = await fetch(`${process.env.NEXT_PUBLIC_URL_SERVICE_HOTELS}/api/hotel/${hotelID}`, {
            method: 'GET',
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