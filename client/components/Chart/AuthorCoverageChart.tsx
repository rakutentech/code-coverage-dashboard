import styles from '../../styles/Basic.module.css'
import {Coverage} from '../../interfaces'

import { Chart as ChartJS, CategoryScale, LinearScale, BarController, BarElement } from "chart.js";
ChartJS.register(CategoryScale);
ChartJS.register(LinearScale);
ChartJS.register(BarController);
ChartJS.register(BarElement);
import { Bar } from 'react-chartjs-2'


interface Props {
    data: {string: Coverage[]};
    orgName: string;
}
export const AuthorCoverageChart = (props:Props) => {
    const {orgName, data} = props
    let authors: string[] = []
    let percentagesAvg: Number[] = []
    Object.keys(data).map(orgName => {
        if (!authors.includes(data[orgName].commit_author)) {
            authors.push(data[orgName].commit_author)
        }
    })
    authors.forEach(author => {
        let sum = 0
        let count = 0
        Object.keys(data).map(orgName => {
            if (data[orgName].commit_author === author) {
                sum += parseFloat(data[orgName].percentage)
                count++
            }
        })
        percentagesAvg.push(sum/count)
    })

    const cdata = {
        labels: authors,
        datasets: [
          {
            label: 'My First dataset',
            fill: false,
            lineTension: 0.1,
            backgroundColor: 'rgba(179,181,198,0.2)',
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
            data: percentagesAvg
          }
        ]
      };
    return (
    <div className={styles.chart_wrapper}>
        <Bar
            options={{
                indexAxis: 'x',
                scales: {
                    x: {
                        title: {
                            display: false,
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