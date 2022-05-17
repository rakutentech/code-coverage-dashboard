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
export const BranchesCoverageChart = (props:Props) => {
    const {orgName, data} = props
    let branches: string[] = []
    let percentages: Number[] = []
    Object.keys(data).map(orgName => {
        branches.push(data[orgName].branch_name)
        percentages.push(parseFloat(data[orgName].percentage))
    })

    const cdata = {
        labels: branches,
        datasets: [
          {
            label: 'My First dataset',
            fill: false,
            lineTension: 0.1,
            backgroundColor: '#0f91e8',
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
        <Bar
            options={{
                // make it horizontal
                indexAxis: 'y',
            }}
            data={cdata}
        />
    </div>
    )
}