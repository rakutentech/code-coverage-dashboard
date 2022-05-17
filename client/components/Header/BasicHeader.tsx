import styles from '../../styles/Basic.module.css'
import Link from 'next/link'

export const BasicHeader = () => {
    return (
        <div className={styles.text_center}>
            <h1>
            <Link href="/">
                <a>Code Coverage Dashboard</a>
             </Link>
            </h1>
            <p className={styles.text_bright + " mb-4"}>
                Consolidated code coverage
            </p>
        </div>
    )
}