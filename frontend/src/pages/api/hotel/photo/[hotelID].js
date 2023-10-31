// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import { generateFileName } from '@/lib/utils/generateFileName';
import multer from 'multer';

const upload = multer({
    storage: multer.diskStorage({
        destination: "./public/", // destination folder
        // filename: (req, file, cb) => cb(null, file.originalname),
        filename: (req, file, cb) => cb(null, generateFileName(file.originalname)),
    }),
});


export default async function handler(req, res) {
    try {
        const { hotelID } = req.query;
        
        upload.single('file')(req, res, async (err) => {
            if (err) {
                console.error('Error uploading file:', err);
                res.status(400).json({ error: 'File upload failed' });
                return
            }

            const file = req.file;
            const fileName = file.filename;
            //console.log(fileName)


            const response = await fetch(`${process.env.NEXT_PUBLIC_URL_SERVICE_HOTELS}/api/hotel/photo/${hotelID}`, {
                method: 'POST',
                body: JSON.stringify({
                    url: fileName
                }),
                headers: {
                    'Content-Type': 'application/json',
                    'Cookie': req.headers.cookie || ''
                },
                credentials: 'include',
            });

            const data = await response.json();
            res.status(response.status).json(data);

        })


    } catch (error) {
        //console.log(error)
        res.status(500).json({ error: 'Internal Server Error' });
    }
}

export const config = {
    api: {
        bodyParser: false, // Disallow body parsing, consume as stream
    },
};