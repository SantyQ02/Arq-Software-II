export default async function handler(req, res) {
    try {
      const response = await fetch(`${process.env.NEXT_PUBLIC_URL_API}/api/auth/refresh`, {
        method: 'GET',
        headers: {
          ...req.headers.JSON,
          'Cookie': req.headers.cookie || ''
        },
        credentials: 'include',
      });
  
      const data = await response.json();
      res.status(response.status).json(data);
    } catch (error) {
      res.status(500).json({ error: 'Internal Server Error' });
    }
  }