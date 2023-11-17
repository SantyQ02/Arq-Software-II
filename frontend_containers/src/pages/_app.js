import CustomToastContainer from '@/components/alert/Alert'
import '@/styles/globals.css'

export default function App({ Component, pageProps }) {
  return (
    <>
      <Component {...pageProps} />
      <CustomToastContainer />
    </>
  )
}
