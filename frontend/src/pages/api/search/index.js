// Next.js API route support: https://nextjs.org/docs/api-routes/introduction

export default async function handler(req, res) {
    try {
      const {city_code, check_in_date, check_out_date } = req.query;

      const response = await fetch(`${process.env.NEXT_PUBLIC_URL_SERVICE_SEARCH}/api/search?&city=${city_code}&check_in_date=${check_in_date}&check_out_date=${check_out_date}`, {
        method: 'GET',
      });
  
      const data = await response.json();
      res.status(response.status).json(data);
    } catch (error) {
      res.status(500).json({ error: 'Internal Server Error' });
    }
  }
  