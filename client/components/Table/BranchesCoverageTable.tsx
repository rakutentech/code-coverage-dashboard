import styles from '../../styles/Basic.module.css'
import {Coverage} from '../../interfaces'
import moment from 'moment'
import Link from 'next/link'

interface Props {
    data: {string: Coverage[]};
    mainPercentage: Number;
    orgName: string;
}
export const BranchesCoverageTable = (props:Props) => {
    const {orgName, mainPercentage, data} = props
    let githubURL = process.env.githubURL
    let apiURL = process.env.apiURL
    const truncateCommitHashToChars = 7; //pleasent to see 7 characters
    const truncateBranchNameToChars = 17; //pleasent to see 17 characters
    return (
    <table className={styles.coverage_table + " table is-narrow is-fullwidth is-bordered"}>
        <thead>
        <tr>
          <th scope="col" align="left">Branch</th>
          <th scope="col" align="left">Language</th>
          <th scope="col" align="left">Commit SHA</th>
          <th scope="col" align="left">PR</th>
          <th scope="col" align="left">Commit Author</th>
          <th scope="col" align="left">Coverage</th>
          <th scope="col" align="left">Report</th>
          <th scope="col" align="left">Branch</th>
          <th scope="col" align="left">Updated</th>
        </tr>
        </thead>
        <tbody>
              {Object.keys(data).map(key =>
                  <tr key={data[key].id}>
                      <td width="20%">
                        <a className={styles.external_link} target="_blank" rel="noreferrer" href={`${githubURL}/${orgName}/tree/${data[key].branch_name}`}>
                            {data[key].branch_name.substring(0, truncateBranchNameToChars)}
                            {data[key].branch_name.length > truncateBranchNameToChars && '...'}
                        </a>
                      </td>
                      <td width="5%">{data[key].language}</td>
                      <td width="10%">
                          <a className={styles.external_link} target="_blank" rel="noreferrer" href={`${githubURL}/${orgName}/commit/${data[key].commit_hash}`}>
                            {data[key].commit_hash.substring(0, truncateCommitHashToChars)}
                          </a>
                      </td>
                      <td width="10%">
                          <a className={styles.external_link} target="_blank" rel="noreferrer" href={`${githubURL}/${orgName}/pull/${data[key].pr_number}`}>
                            #{data[key].pr_number}
                          </a>
                      </td>
                      <td width="20%">
                            <Link href={"/" + data[key].org_name + "/user/" + data[key].commit_author}>
                                <a className='ml-2'>
                                    {data[key].commit_author}
                                </a>
                            </Link>
                      </td>
                      <td width="10%">
                          {mainPercentage === 0 &&
                            <span>{parseFloat(data[key].percentage).toFixed(2)}%</span>
                          }
                          {mainPercentage > 0 && mainPercentage ==  parseFloat(data[key].percentage) &&
                            <span>{parseFloat(data[key].percentage).toFixed(2)}%</span>
                          }
                          {mainPercentage > 0 && mainPercentage > parseFloat(data[key].percentage) &&
                            <b className={styles.text_danger}>{parseFloat(data[key].percentage).toFixed(2)}% ▼</b>
                          }
                          {mainPercentage > 0 && mainPercentage < parseFloat(data[key].percentage) &&
                            <b className={styles.text_success}>{parseFloat(data[key].percentage).toFixed(2)}% ▲</b>
                          }
                      </td>
                      <td width="10%">
                        <a target="_blank" rel="noreferrer" href={`${apiURL}/assets/${orgName}/${data[key].branch_name.replace (/\//g, "_fs_")}/`}>
                            artifacts
                        </a>
                        {data[key].link}
                      </td>
                      <td width="10%">
                          {data[key].deleted_at && <b className={styles.text_muted}>Deleted</b>}
                          {!data[key].deleted_at && <b className={styles.text_success}>Active</b>}
                      </td>
                      <td width="40%">
                          <small>{moment(data[key].updated_at).fromNow()}</small>
                      </td>
                  </tr>
              )}
        </tbody>
     </table>
    )
}
