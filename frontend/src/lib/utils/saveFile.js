import fs from 'fs';

export async function saveFile(file){
    const filePath = `/images/hotels/${file.name}`;

    // Guardar la foto en el sistema de archivos
    try {
      fs.writeFileSync(filePath, file.data);
      //console.log('File saved successfully.');
    } catch (error) {
      console.error('Error saving file:', error);
    }
}