import styles from '../../styles/Basic.module.css'
import Link from 'next/link'
import Image from 'next/image'

export const BasicHeader = () => {
    return (
        <div className={styles.text_center}>
            <h1>
            <Link href="/">
                <a>
                    {/* Insert Logo Image */}
                    <Image src="https://i.imgur.com/PzEBE2j.png" alt="Code Coverage Dashboard" width={200} height={150} />
                </a>
             </Link>
            </h1>
            <p className={styles.text_bright + " mb-4"}>
                Consolidated code coverage
            </p>
        </div>
    )
}