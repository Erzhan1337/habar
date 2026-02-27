"use client";
import React, { useState, useEffect } from "react";
import {
  AreaChart,
  Area,
  XAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
} from "recharts";
import { ChartDataPoint, TimeFilter, fetchChartData } from "../types/chart";

// --- Кастомный Tooltip (всплывающая подсказка) ---
const CustomTooltip = ({ active, payload, label }: any) => {
  if (active && payload && payload.length) {
    return (
      <div className="bg-white p-3 border border-gray-100 shadow-lg rounded-lg text-sm">
        <p className="text-gray-500 font-medium mb-1">{label}</p>
        {payload.map((entry: any, index: number) => (
          <div key={`item-${index}`} className="flex items-center gap-2 mb-1">
            {/* Цветная точка */}
            <span
              className="w-2 h-2 rounded-full"
              style={{ backgroundColor: entry.color }}
            />
            <span className="text-gray-700">
              {entry.name === "products" ? "Продукты" : "Количество ошибок"}:
            </span>
            <span className="font-bold">{entry.value}</span>
          </div>
        ))}
      </div>
    );
  }
  return null;
};

// --- Основной компонент ---
export const SystemErrorsChart = () => {
  const [filter, setFilter] = useState<TimeFilter>("7days");
  const [data, setData] = useState<ChartDataPoint[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  // Имитация загрузки данных с бэкенда
  useEffect(() => {
    let isMounted = true;
    setIsLoading(true);

    fetchChartData(filter).then((fetchedData) => {
      if (isMounted) {
        setData(fetchedData);
        setIsLoading(false);
      }
    });

    return () => {
      isMounted = false;
    };
  }, [filter]);

  return (
    <div className="bg-white rounded-2xl border border-gray-100 p-6 shadow-sm w-full mt-10">
      <div className="flex justify-between items-start mb-8">
        <div>
          <h2 className="text-lg font-bold text-gray-900">
            Последние системные ошибки
          </h2>
          <p className="text-sm text-gray-500 mt-1">
            Недавние проблемы с API и системой
          </p>
        </div>

        <div className="flex gap-2">
          {(["today", "7days", "30days"] as TimeFilter[]).map((f) => (
            <button
              key={f}
              onClick={() => setFilter(f)}
              className={`px-4 py-2 text-sm font-medium rounded-lg transition-colors ${
                filter === f
                  ? "bg-gray-900 text-white"
                  : "bg-white text-gray-700 border border-gray-200 hover:bg-gray-50"
              }`}
            >
              {f === "today" ? "Сегодня" : f === "7days" ? "7 дней" : "30 дней"}
            </button>
          ))}
        </div>
      </div>

      <div className="h-[300px] w-full">
        {isLoading ? (
          <div className="w-full h-full flex items-center justify-center text-gray-400">
            Загрузка данных...
          </div>
        ) : (
          <ResponsiveContainer width="100%" height="100%">
            <AreaChart
              data={data}
              margin={{ top: 10, right: 0, left: 0, bottom: 0 }}
            >
              <defs>
                <linearGradient id="colorProducts" x1="0" y1="0" x2="0" y2="1">
                  <stop offset="5%" stopColor="#6ee7b7" stopOpacity={0.4} />
                  <stop offset="95%" stopColor="#6ee7b7" stopOpacity={0} />
                </linearGradient>
                <linearGradient id="colorErrors" x1="0" y1="0" x2="0" y2="1">
                  <stop offset="5%" stopColor="#fdba74" stopOpacity={0.4} />
                  <stop offset="95%" stopColor="#fdba74" stopOpacity={0} />
                </linearGradient>
              </defs>

              <CartesianGrid
                strokeDasharray="3 3"
                vertical={false}
                stroke="#f3f4f6"
              />

              <XAxis
                dataKey="date"
                axisLine={false}
                tickLine={false}
                tick={{ fill: "#9ca3af", fontSize: 12 }}
                dy={10}
              />

              <Tooltip content={<CustomTooltip />} />

              <Area
                type="monotone"
                dataKey="products"
                stroke="#10b981"
                strokeWidth={2}
                fillOpacity={1}
                fill="url(#colorProducts)"
              />
              <Area
                type="monotone"
                dataKey="errors"
                stroke="#f97316"
                strokeWidth={2}
                fillOpacity={1}
                fill="url(#colorErrors)"
              />
            </AreaChart>
          </ResponsiveContainer>
        )}
      </div>
    </div>
  );
};
