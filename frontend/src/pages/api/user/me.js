// Next.js API route support: https://nextjs.org/docs/api-routes/introduction

export default async function handler(req, res) {
  try {
    const response = await fetch(`${process.env.NEXT_PUBLIC_URL_API}/api/user/me`, {
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
