import CustomToastContainer from '@/components/alert/Alert'
import LayoutContext from '@/layouts/LayoutContext'
import '@/styles/globals.css'

export default function App({ Component, pageProps }) {
  return (
    <LayoutContext>
      <Component {...pageProps} />
      <CustomToastContainer />
    </LayoutContext>
  )
}
