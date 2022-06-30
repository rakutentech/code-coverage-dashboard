import styles from '../../styles/Basic.module.css'
interface Props {
  keyword: string;
  onChangeKeyword: (key: string) => void;
}

export const SearchForm: React.FC<Props> = ({keyword, onChangeKeyword}) => {
  const onChangeHandler = (e: React.ChangeEvent<HTMLInputElement>) => onChangeKeyword(e.target.value)
  return (
    <div className="field mx-5 mt-6 mb-0">
      <p className="control">
        <input className={"input " + styles.dark_mode} value={keyword} type="email" placeholder="Filter by org / repo name" onChange={onChangeHandler} />
      </p>
    </div>
  )
}