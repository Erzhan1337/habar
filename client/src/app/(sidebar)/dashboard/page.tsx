"use client";

import { Card, Table, Radio } from "antd";
import {
  CodeSandboxOutlined,
  ExclamationCircleOutlined,
  SlidersOutlined,
  CoffeeOutlined,
  HistoryOutlined,
  LineChartOutlined,
} from "@ant-design/icons";
import {
  AreaChart,
  Area,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
} from "recharts";

const chartData = [
  { date: "Feb 11", products: 600, errors: 50 },
  { date: "Feb 12", products: 1258, errors: 300 },
  { date: "Feb 13", products: 800, errors: 150 },
  { date: "Feb 14", products: 600, errors: 100 },
  { date: "Feb 15", products: 750, errors: 80 },
  { date: "Feb 16", products: 780, errors: 120 },
];

const tableColumns = [
  {
    title: "Время",
    dataIndex: "time",
    key: "time",
    className: "text-gray-500",
  },
  {
    title: "Критичность",
    dataIndex: "severity",
    key: "severity",
    render: (severity: string) => {
      let bg = "bg-blue-50 text-blue-600 border-blue-200";
      if (severity === "Предупреждение") {
        bg = "bg-orange-50 text-orange-500 border-orange-200";
      } else if (severity === "Ошибка") {
        bg = "bg-red-50 text-red-500 border-red-200";
      }
      return (
        <span className={`px-3 py-1 rounded-full text-xs border ${bg}`}>
          {severity}
        </span>
      );
    },
  },
  {
    title: "Провайдер",
    dataIndex: "provider",
    key: "provider",
  },
  {
    title: "Сообщение",
    dataIndex: "message",
    key: "message",
    className: "text-gray-500",
  },
  {
    title: "Затронуто запросов",
    dataIndex: "affected",
    key: "affected",
  },
];

const tableData = [
  {
    key: "1",
    time: "2026-02-17 08:42:15",
    severity: "Предупреждение",
    provider: "OpenFood",
    message: "Rate limit exceeded (429)",
    affected: 142,
  },
  {
    key: "2",
    time: "2026-02-17 08:30:04",
    severity: "Ошибка",
    provider: "Kcell",
    message: "Timeout (504 Gateway Ti...",
    affected: 87,
  },
  {
    key: "3",
    time: "2026-02-17 07:55:21",
    severity: "Инфо",
    provider: "Kaspi.kz",
    message: "Cache обновлён",
    affected: 0,
  },
  {
    key: "4",
    time: "2026-02-17 07:40:18",
    severity: "Предупреждение",
    provider: "Kazakhtelecom",
    message: "Медленный ответ (>3s)",
    affected: 23,
  },
  {
    key: "5",
    time: "2026-02-17 07:15:02",
    severity: "Ошибка",
    provider: "Halyk Bank",
    message: "Unauthorized (401)",
    affected: 56,
  },
];

export default function DashboardPage() {
  return (
    <div className="flex flex-col gap-6 py-10">
      <h1 className="text-2xl font-bold m-0">Дашборд</h1>

      <div className="grid grid-cols-3 gap-6">
        <Card className="rounded-xl shadow-sm">
          <div className="flex justify-between items-start">
            <div>
              <p className="text-gray-400 text-sm mb-2">Всего продуктов</p>
              <h3 className="text-3xl font-bold mb-1">1,324</h3>
              <p className="text-green-600 text-xs m-0">+84 за неделю</p>
            </div>
            <CodeSandboxOutlined className="text-2xl text-gray-400" />
          </div>
        </Card>

        <Card className="rounded-xl shadow-sm">
          <div className="flex justify-between items-start">
            <div>
              <p className="text-gray-400 text-sm mb-2">
                Ненайденные штрихкоды
              </p>
              <h3 className="text-3xl font-bold mb-1">47</h3>
              <p className="text-red-500 text-xs m-0">+12 в ожидании</p>
            </div>
            <ExclamationCircleOutlined className="text-2xl text-gray-400" />
          </div>
        </Card>

        <Card className="rounded-xl shadow-sm">
          <div className="flex justify-between items-start">
            <div>
              <p className="text-gray-400 text-sm mb-2">Продукты с сахаром</p>
              <h3 className="text-3xl font-bold mb-1">73.2%</h3>
              <p className="text-gray-400 text-xs m-0">969 продуктов</p>
            </div>
            <SlidersOutlined className="text-2xl text-gray-400" />
          </div>
        </Card>

        <Card className="rounded-xl shadow-sm">
          <div className="flex justify-between items-start">
            <div>
              <p className="text-gray-400 text-sm mb-2">Продукты с кофеином</p>
              <h3 className="text-3xl font-bold mb-1">42.8%</h3>
              <p className="text-gray-400 text-xs m-0">567 продуктов</p>
            </div>
            <CoffeeOutlined className="text-2xl text-gray-400" />
          </div>
        </Card>

        <Card className="rounded-xl shadow-sm">
          <div className="flex justify-between items-start">
            <div>
              <p className="text-gray-400 text-sm mb-2">Устаревшие записи</p>
              <h3 className="text-3xl font-bold mb-1">28</h3>
              <p className="text-orange-400 text-xs m-0">Требует обновления</p>
            </div>
            <HistoryOutlined className="text-2xl text-gray-400" />
          </div>
        </Card>

        <Card className="rounded-xl shadow-sm">
          <div className="flex justify-between items-start">
            <div>
              <p className="text-gray-400 text-sm mb-2">
                Статус API провайдеров
              </p>
              <div className="mt-4 flex gap-2">
                <span className="px-2 py-0.5 bg-green-50 text-green-500 rounded-full text-xs border border-green-100">
                  3 онлайн
                </span>
                <span className="px-2 py-0.5 bg-red-50 text-red-500 rounded-full text-xs border border-red-100">
                  1 недоступен
                </span>
              </div>
            </div>
            <LineChartOutlined className="text-2xl text-gray-400" />
          </div>
        </Card>
      </div>

      <Card className="rounded-xl shadow-sm">
        <div className="flex justify-between items-center mb-6">
          <div>
            <h3 className="text-lg font-bold m-0">
              Последние системные ошибки
            </h3>
            <p className="text-gray-400 text-sm m-0">
              Недавние проблемы с API и системой
            </p>
          </div>
          <Radio.Group defaultValue="30" buttonStyle="solid">
            <Radio.Button value="1">Сегодня</Radio.Button>
            <Radio.Button value="7">7 дней</Radio.Button>
            <Radio.Button
              value="30"
              className="bg-gray-900 border-gray-900 text-white hover:text-white"
            >
              30 дней
            </Radio.Button>
          </Radio.Group>
        </div>
        <div className="h-62 w-full">
          <ResponsiveContainer width="100%" height="100%">
            <AreaChart
              data={chartData}
              margin={{ top: 10, right: 0, left: 0, bottom: 0 }}
            >
              <defs>
                <linearGradient id="colorProducts" x1="0" y1="0" x2="0" y2="1">
                  <stop offset="5%" stopColor="#2A9D90" stopOpacity={0.3} />
                  <stop offset="95%" stopColor="#2A9D90" stopOpacity={0} />
                </linearGradient>
                <linearGradient id="colorErrors" x1="0" y1="0" x2="0" y2="1">
                  <stop offset="5%" stopColor="#E76E50" stopOpacity={0.3} />
                  <stop offset="95%" stopColor="#E76E50" stopOpacity={0} />
                </linearGradient>
              </defs>
              <CartesianGrid
                strokeDasharray="3 3"
                vertical={false}
                stroke="#f0f0f0"
              />
              <XAxis
                dataKey="date"
                axisLine={false}
                tickLine={false}
                tick={{ fill: "#888", fontSize: 12 }}
                dy={10}
              />
              <YAxis axisLine={false} tickLine={false} tick={false} width={0} />
              <Tooltip
                contentStyle={{
                  borderRadius: "8px",
                  border: "none",
                  boxShadow: "0 4px 12px rgba(0,0,0,0.1)",
                }}
              />
              <Area
                type="monotone"
                dataKey="products"
                stroke="#20b2aa"
                strokeWidth={2}
                fill="url(#colorProducts)"
              />
              <Area
                type="monotone"
                dataKey="errors"
                stroke="#ff7a45"
                strokeWidth={2}
                fill="url(#colorErrors)"
              />
            </AreaChart>
          </ResponsiveContainer>
        </div>
      </Card>

      <Card className="rounded-xl shadow-sm">
        <div className="mb-6">
          <h3 className="text-lg font-bold m-0">Последние системные ошибки</h3>
          <p className="text-gray-400 text-sm m-0">
            Недавние проблемы с API и системой
          </p>
        </div>
        <Table
          columns={tableColumns}
          dataSource={tableData}
          pagination={false}
          className="custom-table"
        />
      </Card>
    </div>
  );
}
