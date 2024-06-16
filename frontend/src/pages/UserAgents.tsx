import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Bar } from 'react-chartjs-2';

interface AggregationResult {
    key: string;
    doc_count: number;
}

const UserAgents: React.FC = () => {
    const [data, setData] = useState<AggregationResult[]>([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchData = async () => {
            const response = await axios.get('http://localhost:8080/api/aggregations/user_agents');
            setData(response.data);
            setLoading(false);
        };

        fetchData();
    }, []);

    if (loading) {
        return <p>Loading...</p>;
    }

    const chartData = {
        labels: data.map(item => item.key),
        datasets: [
            {
                label: 'Agenty Użytkowników',
                data: data.map(item => item.doc_count),
                backgroundColor: 'rgba(255, 99, 132, 0.2)',
                borderColor: 'rgba(255, 99, 132, 1)',
                borderWidth: 1,
            },
        ],
    };

    return (
        <div>
            <h1 className="title">Agenty Użytkowników</h1>
            <Bar data={chartData} height="90%" />
        </div>
    );
};

export default UserAgents;
