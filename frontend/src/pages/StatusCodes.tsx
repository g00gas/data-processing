import React, { useEffect, useState } from "react";
import axios from "axios";
import { Bar } from "react-chartjs-2";

interface AggregationResult {
  key: string;
  doc_count: number;
}

const StatusCodes: React.FC = () => {
  const [data, setData] = useState<AggregationResult[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      const response = await axios.get(
        "http://localhost:8080/api/aggregations/status_codes"
      );
      setData(response.data);
    };

    fetchData();
  }, []);

  const chartData = {
    labels: data.map((item) => item.key),
    datasets: [
      {
        label: "Kody HTTP",
        data: data.map((item) => item.doc_count),
        backgroundColor: "rgba(255, 159, 64, 0.2)",
        borderColor: "rgba(255, 159, 64, 1)",
        borderWidth: 1,
      },
    ],
  };

  return (
    <div>
      <h1 className="title">Kody HTTP</h1>
      <Bar data={chartData} height="90%" />
    </div>
  );
};

export default StatusCodes;
