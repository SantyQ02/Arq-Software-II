import crypto from 'crypto'
import path from 'path';

export function generateFileName(file_name) {
    try {
        const fileExtension = path.extname(file_name);
    const name = crypto.randomUUID()
    const newFileName = `images/hotels/${name}${fileExtension}`;
    return newFileName
    } catch (error) {
        //console.log("Error generating file name",error)
    }
    
 
}