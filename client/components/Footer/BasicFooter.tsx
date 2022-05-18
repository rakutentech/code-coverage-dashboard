import styles from '../../styles/Basic.module.css'
import Link from 'next/link'

export const BasicFooter = () => {
    return (
        <footer className={styles.text_center}>
            <p>
                <small>Developed By</small>
            </p>
            <small className={styles.text_muted}>
                <Link href="https://github.com/rakutentech/code-coverage-dashboard">Rakutentech</Link>
                <span className='ml-1'>- Code Coverage Dashboard</span>
            </small>
        </footer>
    )
}