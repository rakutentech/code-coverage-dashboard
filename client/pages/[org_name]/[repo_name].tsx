import { RepositoryCoverageCard, BasicHeader, BasicFooter } from '../../components'
import styles from '../styles/Basic.module.css'
import {Coverage} from '../../interfaces/'
import { useRouter } from 'next/router'

interface ServerSideProps {
    coverages: {
        data: {string: Coverage[]},
        has_next: boolean;
    };
}

const RepoDetails: React.FC<ServerSideProps> = (props) => {
  const { coverages } = props
  return (
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
  )
}

// This function gets called at build time on server-side.
export async function getServerSideProps(context: any) {
    let apiURL = process.env.apiURL
    const orgName = context.query.org_name
    const repoName = context.query.repo_name
    const queryURL = apiURL + `?org_name=${orgName}&repo_name=${repoName}&full=true&per_page=10000`
    console.log(queryURL)
    const res = await fetch(queryURL)
    const coverages = await res.json()
    return {
      props: {
        coverages,
      },
    }
  }

export default RepoDetails
