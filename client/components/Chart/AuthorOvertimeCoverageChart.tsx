import styles from '../../styles/Basic.module.css'
import {Coverage} from '../../interfaces'
import moment from 'moment'

import { Chart as ChartJS, CategoryScale, LinearScale, BarController, BarElement, PointElement, LineElement } from "chart.js";
ChartJS.register(CategoryScale);
ChartJS.register(LinearScale);
ChartJS.register(BarController);
ChartJS.register(BarElement);
ChartJS.register(PointElement);
ChartJS.register(LineElement);
import { Line } from 'react-chartjs-2'


interface Props {
    data: {string: Coverage[]};
    orgName: string;
    commitAuthor: string;
}
export const AuthorOvertimeCoverageChart = (props:Props) => {
    const {orgName, commitAuthor, data} = props
    let dates: string[] = []
    let percentages: Number[] = []
    Object.keys(data).map(orgName => {
        if (data[orgName].commit_author !== commitAuthor) {
            return
        }
        let dateF = moment(data[orgName].updated_at, 'YYYY-MM-DD HH:SS').format('YYYY-MM-DD HH:SS')
        if (!dates.includes(dateF)) {
            dates.push(dateF)
            percentages.push(parseFloat(data[orgName].percentage))
        }
    })

    const cdata = {
        labels: dates,
        datasets: [
          {
            label: 'My First dataset',
            fill: false,
            lineTension: 0.1,
            backgroundColor: 'rgba(255,99,132,0.2)',
            borderColor: 'rgba(75,192,192,1)',
            // borderCapStyle: 'butt',
            borderDash: [],
            borderDashOffset: 0.0,
            // borderJoinStyle: 'miter',
            pointBorderColor: 'rgba(75,192,192,1)',
            pointBackgroundColor: '#fff',
            pointBorderWidth: 1,
            pointHoverRadius: 5,
            pointHoverBackgroundColor: 'rgba(75,192,192,1)',
            pointHoverBorderColor: 'rgba(220,220,220,1)',
            pointHoverBorderWidth: 2,
            pointRadius: 1,
            pointHitRadius: 10,
            data: percentages
          }
        ]
      };
    return (
    <div className={styles.chart_wrapper}>
        <Line
            options={{
                // make it horizontal
                indexAxis: 'x',
                scales: {
                    x: {
                        title: {
                            display: true,
                            text: "Commit Author",
                        },
                    },
                    y: {
                        title: {
                            display: true,
                            text: "Covergae Percentage",
                        },
                    },
                }
            }}
            data={cdata}
        />
    </div>
    )
}