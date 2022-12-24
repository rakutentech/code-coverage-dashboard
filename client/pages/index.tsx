import styles from '../styles/Basic.module.css'
import { useState, useEffect, useCallback } from 'react'
import { GetServerSideProps } from 'next'
import { useRouter } from 'next/router'
import { RepositoryCoverageCard, BasicHeader, BasicFooter } from '../components'
import {Coverage} from '../interfaces/'
import Head from 'next/head'
import { SearchForm } from '../components/Form/SearchForm'
import { queryToString } from '../libs/strings'
import { debounce } from '../libs/debounce'

interface ServerSideProps {
    coverages: {
        data: {string: Coverage[]},
        has_next: boolean;
    };
}

const debounceIntervalMsec = 500

const Home: React.FC<ServerSideProps> = ({ coverages }) => {
  const router = useRouter()
  const { query, pathname } = router
  const { keyword: defaultKeyword } = query

  const [keyword, setKeyword] = useState<string>(queryToString(defaultKeyword))

  const debounceCallback = async (keyword: string) => {
    // update keyword query param
    const queryParam = keyword ? { keyword } : null
    router.push({
      query: queryParam, pathname,
    }, undefined, { shallow: false })
  }

  // make a function to debounce processing every 500 msec
  // NOTE: Needs to be wrapped in useCallback as it is overwritten at each rendering
  // eslint-disable-next-line react-hooks/exhaustive-deps
  const debounceOnChangeHandler = useCallback(debounce<string>(debounceCallback, debounceIntervalMsec), [])

  useEffect(() => {
    debounceOnChangeHandler(keyword)
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [ keyword ])

  return (
      <>
        <Head>
            <title>Code Coverage Dashboard</title>
        </Head>
        <div className="container">
            <BasicHeader />
            <SearchForm keyword={keyword} onChangeKeyword={setKeyword} />
            <section>
            {!coverages?.data || Object.keys(coverages.data).length < 1 && (
                <article className={styles.dark_mode + " m-5 message"}>
                  <div className={styles.text_bright + " message-body text-center"}>
                    <h3>Please search with different keywords</h3>
                  </div>
                </article>
            )}
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
export const getServerSideProps: GetServerSideProps<ServerSideProps> = async (ctx) => {
  const apiURL = process.env.apiURL
  const { keyword } = ctx.query
  const queryURL = `${apiURL}?per_page=10000${keyword ? `&keyword=${keyword}` : ''}`  
  
  try {
    const res = await fetch(queryURL)
    const coverages = await res.json()
    return {
      props: {
        coverages
      },
    }    
  } catch (err) {
    console.log(err)
    return {
      props: {
        coverages: {
          data: {},
          has_next: false,
        }
      },
    }
  }
}

export default Home
