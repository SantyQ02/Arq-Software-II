// Next.js API route support: https://nextjs.org/docs/api-routes/introduction

export default async function handler(req, res) {
    try {
      const { hotel, date_in, date_out } = req.query;

      const response = await fetch(`${process.env.NEXT_PUBLIC_URL_SERVICE_BUSINESS}/api/booking/search/me?hotel=${hotel}&date_in=${date_in}&date_out=${date_out}`, {
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
  