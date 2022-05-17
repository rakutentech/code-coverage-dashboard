import styles from '../../styles/Basic.module.css'
import {Coverage} from '../../interfaces'
import {BranchesCoverageChart, AuthorCoverageChart, OvertimeCoverageChart, BranchesCoverageTable} from '..'
import Link from 'next/link'

interface Props {
    data: {string: Coverage[]};
    orgName: string;
}
export const RepositoryCoverageCard = (props:Props) => {
    const {orgName, data} = props
    let mainCoverage = false
    let mainCoverageBranch = ""
    let mainCoveragePercentage = 0.0
    return (
    <div className={styles.coverage_card + " card rounded m-5"}>
        <header className="card-header">
          <span className="card-header-title">
            <h1 className={styles.text_bright}>
                <b>{orgName}</b>
            </h1>
            <Link href={"/" + orgName}>
                <a className='ml-2'>
                    View full data
                </a>
            </Link>
          </span>
        </header>
        <h2 className={styles.coverage_total}>
        {Object.keys(data).map(orgName =>
            {if ((data[orgName].branch_name == 'develop' || data[orgName].branch_name == 'master') && !mainCoverage) {
                mainCoverage = true
                mainCoverageBranch = data[orgName].branch_name
                mainCoveragePercentage = parseFloat(data[orgName]['percentage'])
                return mainCoveragePercentage + "%"
            }
        })}
        </h2>
        {mainCoverageBranch &&
            <div className={styles.text_center + " mt-1"}>
                <small className={styles.text_muted}>Current coverage from {mainCoverageBranch}</small>
            </div>
        }
        <div className="columns p-2">
            <div className="column">
                <BranchesCoverageChart data={data} orgName={orgName} />
            </div>
            <div className="column">
                <AuthorCoverageChart data={data} orgName={orgName} />
            </div>
            <div className="column">
                <OvertimeCoverageChart data={data} orgName={orgName} />
            </div>
        </div>

        <div className="card-content">
          <div className="content">
            <BranchesCoverageTable data={data} orgName={orgName} mainPercentage={mainCoveragePercentage}/>
          </div>
        </div>
      </div>
    )
}