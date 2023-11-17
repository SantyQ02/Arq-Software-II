// Next.js API route support: https://nextjs.org/docs/api-routes/introduction

export default async function handler(req, res) {
    try {
      const { hotelID } = req.query;
  
      const response = await fetch(`${process.env.NEXT_PUBLIC_URL_SERVICE_HOTELS}/api/hotel/amenitie/unloadamenities/${hotelID}`, {
        method: 'PUT',
        body: JSON.stringify(req.body), 
        headers: {
            ...req.headers.JSON,
            'Cookie': req.headers.cookie || ''
        },
        credentials: 'include',
      });
      let data
      try {
        data = await response.json()
      } catch (error) {
        data = ""
      }
      res.status(response.status).json(data);
    } catch (error) {
      res.status(500).json({ error: 'Internal Server Error' });
    }
  }