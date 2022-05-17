import { RepositoryCoverageCard, BasicHeader, BasicFooter } from '../components'
import styles from '../styles/Basic.module.css'
import {Coverage} from '../interfaces/'
import Head from 'next/head'

interface ServerSideProps {
    coverages: {
        data: {string: Coverage[]},
        has_next: boolean;
    };
}

const Home: React.FC<ServerSideProps> = (props) => {
  const { coverages } = props
  return (
      <>
        <Head>
            <title>Code Coverage Dashboard</title>
        </Head>
        <div className="container">
            <BasicHeader />
            <section>
            {Object.keys(coverages.data).map(orgName =>
                <div key={orgName} className="pt-5">
                    <RepositoryCoverageCard orgName={orgName} data={coverages.data[orgName]} />
                </div>
            )}
            </section>
            <BasicFooter />
        </div>
    </>
  )
}

// This function gets called at build time on server-side.
export async function getServerSideProps(context: any) {
    let apiURL = process.env.apiURL
    const queryURL = apiURL + '?per_page=10000'
    const res = await fetch(queryURL)
    const coverages = await res.json()
    return {
      props: {
        coverages
      },
    }
  }

export default Home
