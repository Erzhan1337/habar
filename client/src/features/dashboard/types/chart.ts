// types.ts
export interface ChartDataPoint {
  date: string;
  products: number;
  errors: number;
}

export type TimeFilter = "today" | "7days" | "30days";

// mockApi.ts
// Это временная заглушка. Позже вы просто замените ее на реальный fetch/axios запрос.
export const fetchChartData = async (
  filter: TimeFilter,
): Promise<ChartDataPoint[]> => {
  // Имитация задержки сети
  await new Promise((resolve) => setTimeout(resolve, 500));

  // Возвращаем данные в зависимости от фильтра
  if (filter === "7days") {
    return [
      { date: "Feb 11", products: 400, errors: 120 },
      { date: "Feb 12", products: 1258, errors: 450 },
      { date: "Feb 13", products: 800, errors: 200 },
      { date: "Feb 14", products: 500, errors: 380 },
      { date: "Feb 15", products: 900, errors: 250 },
      { date: "Feb 16", products: 850, errors: 300 },
    ];
  }
  return []; // Тут могут быть другие моки для 'today' и '30days'
};
