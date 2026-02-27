import React from "react";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/shared/ui/table";
import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
} from "@/shared/ui/card";
import { Badge } from "@/shared/ui/badge";

// Мок-данные для таблицы
const logsData = [
  {
    id: "1",
    timestamp: "2024-02-17 08:42:15",
    severity: "warning",
    provider: "OpenFood",
    message: "Rate limit exceeded (429)",
    affectedRequests: 142,
  },
  {
    id: "2",
    timestamp: "2024-02-17 08:30:04",
    severity: "error",
    provider: "Kcell",
    message: "Timeout (504 Gateway Ti...",
    affectedRequests: 87,
  },
  {
    id: "3",
    timestamp: "2024-02-17 07:55:21",
    severity: "info",
    provider: "Kaspi.kz",
    message: "Cache обновлен",
    affectedRequests: 0,
  },
  {
    id: "4",
    timestamp: "2024-02-17 07:40:18",
    severity: "warning",
    provider: "Kazakhtelecom",
    message: "Медленный ответ (>3s)",
    affectedRequests: 23,
  },
  {
    id: "5",
    timestamp: "2024-02-17 07:15:02",
    severity: "error",
    provider: "Halyk Bank",
    message: "Unauthorized (401)",
    affectedRequests: 56,
  },
];

// Вспомогательная функция для стилизации бейджей
const getSeverityBadge = (severity: string) => {
  switch (severity) {
    case "warning":
      return (
        <Badge
          variant="outline"
          className="bg-orange-50 text-orange-600 border-orange-200 font-normal hover:bg-orange-50"
        >
          Предупреждение
        </Badge>
      );
    case "error":
      return (
        <Badge
          variant="outline"
          className="bg-red-50 text-red-600 border-red-200 font-normal hover:bg-red-50"
        >
          Ошибка
        </Badge>
      );
    case "info":
      return (
        <Badge
          variant="outline"
          className="bg-blue-50 text-blue-600 border-blue-200 font-normal hover:bg-blue-50"
        >
          Инфо
        </Badge>
      );
    default:
      return <Badge variant="outline">{severity}</Badge>;
  }
};

export function SystemLogsTable() {
  return (
    <Card className="shadow-sm rounded-2xl border-gray-100 mt-10">
      <CardHeader className="pb-4">
        <CardTitle className="text-lg font-bold text-gray-900">
          Последние системные ошибки
        </CardTitle>
        <CardDescription className="text-gray-500">
          Недавние проблемы с API и системой
        </CardDescription>
      </CardHeader>

      <CardContent className="p-0">
        <Table>
          <TableHeader className="bg-transparent">
            <TableRow className="hover:bg-transparent border-b border-gray-100">
              <TableHead className="text-gray-500 pl-6">Время</TableHead>
              <TableHead className="text-gray-500">Критичность</TableHead>
              <TableHead className="text-gray-500">Провайдер</TableHead>
              <TableHead className="text-gray-500">Сообщение</TableHead>
              <TableHead className="text-gray-500 pr-6 text-right">
                Затронуто запросов
              </TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {logsData.map((log) => (
              <TableRow
                key={log.id}
                className="border-b border-gray-50 last:border-0 hover:bg-gray-50/50"
              >
                <TableCell className="pl-6 py-4 text-gray-900">
                  {log.timestamp}
                </TableCell>
                <TableCell className="py-4">
                  {getSeverityBadge(log.severity)}
                </TableCell>
                <TableCell className="py-4 text-gray-900">
                  {log.provider}
                </TableCell>
                <TableCell className="py-4 text-gray-900">
                  {log.message}
                </TableCell>
                <TableCell className="pr-6 py-4 text-gray-900 text-right">
                  {log.affectedRequests}
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </CardContent>
    </Card>
  );
}
