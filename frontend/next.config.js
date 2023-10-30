/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  images: {
    domains: [process.env.NEXT_PUBLIC_DOMAIN_API, "static.independent.co.uk", "images.bubbleup.com"],
  },

  // Configuración para servir archivos estáticos desde la carpeta public
  // Durante la fase de producción (npm run start)
  // Asegúrate de reemplazar 'your-domain.com' con tu dominio real si estás desplegando en un dominio personalizado
  async rewrites() {
    return [
      {
        source: '/images/:path*',
        destination: '/public/images/:path*',
      },
      // {
      //   source: '/public',
      //   destination: '/public',
      // },
    ];
  },
};

module.exports = nextConfig;
