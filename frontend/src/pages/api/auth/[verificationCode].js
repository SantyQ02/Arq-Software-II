export default async function handler(req, res) {
    try {
        const { verificationCode } = req.query;
      const response = await fetch(`${process.env.NEXT_PUBLIC_URL_API}/api/auth/verifyemail/${verificationCode}`, {
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