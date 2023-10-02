import Header from '@/components/layout/Header'
import Footer from '@/components/layout/Footer'
import CustomHead from '@/components/layout/CustomHead'

export default function MainLayout({title, children}){

    return (
        <div>

            <CustomHead title={title} />
            <Header/>
            <main className=''>
                {children}
            </main>
            <Footer/>
        </div>
    )   
}