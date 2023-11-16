export default async function handler(req, res) {
    try {
        const { container_id } = req.query;
        const response = await fetch(`${process.env.NEXT_PUBLIC_URL_CONTAINERS_SERVICE}/api/containers/start/${container_id}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Cache-Control': 'no-cache'
            },
        });

        const data = await response.json();
        res.status(response.status).json(data);
    } catch (error) {
        res.status(500).json({ error: 'Internal Server Error' });
    }
}