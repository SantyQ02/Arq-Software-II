// Next.js API route support: https://nextjs.org/docs/api-routes/introduction

export default async function handler(req, res) {
    try {
      const response = await fetch(`${process.env.NEXT_PUBLIC_URL_API}/api/auth/register`, {
        method: 'POST',
        body: JSON.stringify(req.body), 
        headers: req.headers.JSON
    });
  
      const data = await response.json();
      //console.log(data)
      res.status(response.status).json(data);
    } catch (error) {
      res.status(500).json({ error: 'Internal Server Error' });
    }
  }
  