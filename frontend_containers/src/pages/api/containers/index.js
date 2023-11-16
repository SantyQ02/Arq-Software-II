export default async function handler(req, res) {
  if (req.method === 'GET') {
    try {
      const response = await fetch(`${process.env.NEXT_PUBLIC_URL_CONTAINERS_SERVICE}/api/containers`, {
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
  if (req.method === 'POST') {
    try {
      const response = await fetch(`${process.env.NEXT_PUBLIC_URL_CONTAINERS_SERVICE}/api/containers`, {
        method: 'POST',
        body: JSON.stringify(req.body),
        headers: {
          ...req.headers.JSON,
        },
      });

      const data = await response.json();
      res.status(response.status).json(data);
    } catch (error) {
      res.status(500).json({ error: 'Internal Server Error' });
    }
  }
}