import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Bar } from 'react-chartjs-2';
import "chart.js/auto";

interface AggregationResult {
    key: string;
    doc_count: number;
}

const DeviceNames: React.FC = () => {
    const [data, setData] = useState<AggregationResult[]>([]);

    useEffect(() => {
        const fetchData = async () => {
            const response = await axios.get('http://localhost:8080/api/aggregations/device_names');
            setData(response.data);
        };

        fetchData();
    }, []);

    const chartData = {
        labels: data.map(item => item.key),
        datasets: [
            {
                label: 'Nazwy urządzeń',
                data: data.map(item => item.doc_count),
                backgroundColor: 'rgba(75, 192, 192, 0.2)',
                borderColor: 'rgba(75, 192, 192, 1)',
                borderWidth: 1,
            },
        ],
    };

    return (
        <div>
            <h1 className="title">Nazwy urządzeń</h1>
            <Bar data={chartData} height="90%" />
        </div>
    );
};

export default DeviceNames;
