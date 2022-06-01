import styles from '../../../styles/Basic.module.css'
import { AuthorOvertimeCoverageChart, BasicHeader, BasicFooter } from '../../../components'
import {Coverage} from '../../../interfaces'
import Link from 'next/link'

interface ServerSideProps {
    coverages: {
        data: {string: Coverage[]},
        has_next: boolean;
    };
    commitAuthor: string;
}

const UserRepoDetails: React.FC<ServerSideProps> = (props) => {
  const { coverages, commitAuthor } = props
  return (
    <div className="container">
        <BasicHeader />
        <section>
            {Object.keys(coverages.data).map(orgName =>
                <div key={orgName} className={styles.coverage_card + " card rounded m-5 p-2"}>
                    <header className="card-header">
                    <h1 className={styles.text_bright}>
                    <Link href={"/" + orgName}>
                        <a className='ml-2'>
                            {orgName}
                        </a>
                    </Link>
                    </h1>
                    </header>
                    <h2 className={styles.text_bright + " " + styles.text_center + " p-2"}>{commitAuthor}</h2>
                    <AuthorOvertimeCoverageChart data={coverages.data[orgName]} orgName={orgName} commitAuthor={commitAuthor}/>
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
    const commitAuthor = context.query.commit_author
    const queryURL = apiURL + `?org_name=${orgName}&commit_author=${commitAuthor}&full=true&per_page=10000`
    const res = await fetch(queryURL)
    const coverages = await res.json()
    return {
      props: {
        coverages,
        commitAuthor,
      },
    }
  }

export default UserRepoDetails
