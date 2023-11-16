import Head from 'next/head'

const CustomHead = ({title}) => {
    return (
        <Head>
            <title>{`${title} - Project Arq. Software`}</title>
            <meta name="description" content="" />
            <meta name="viewport" content="width=device-width, initial-scale=1" />
            <link rel="icon" href="/favicon.ico" />
        </Head>
    )
}

export default CustomHead