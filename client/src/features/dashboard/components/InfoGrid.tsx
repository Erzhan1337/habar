import { Card } from "@/shared/ui/Card";
import {
  Candy,
  CircleAlert,
  ClockFading,
  Coffee,
  PackageOpen,
  TrendingUp,
} from "lucide-react";

export const InfoGrid = () => {
  return (
    <div className="grid grid-cols-3 gap-10">
      <Card>
        <Card.Header title="Всего продуктов" icon={<PackageOpen size={30} />} />
        <Card.Value>1,324</Card.Value>
        <Card.Footer className="text-green-500">+84 за неделю</Card.Footer>
      </Card>

      <Card>
        <Card.Header
          title="Ненайденные штрихкоды"
          icon={<CircleAlert size={30} />}
        />
        <Card.Value>47</Card.Value>
        <Card.Footer className="text-red-500">+12 в ожидании</Card.Footer>
      </Card>

      <Card>
        <Card.Header title="Продукты с сахаром" icon={<Candy size={30} />} />
        <Card.Value>73.2%</Card.Value>
        <Card.Footer className="text-gray-400">969 продуктов</Card.Footer>
      </Card>

      <Card>
        <Card.Header title="Продукты с кофеином" icon={<Coffee size={30} />} />
        <Card.Value>42.8%</Card.Value>
        <Card.Footer className="text-gray-400">969 продуктов</Card.Footer>
      </Card>

      <Card>
        <Card.Header
          title="Устаревшие запросы"
          icon={<ClockFading size={30} />}
        />
        <Card.Value>28</Card.Value>
        <Card.Footer className="text-orange-400">
          Требует обновление
        </Card.Footer>
      </Card>

      <Card>
        <Card.Header
          title="Статус API провайдеров"
          icon={<TrendingUp size={30} />}
        />
        <Card.Footer>
          <div className="flex items-center gap-2 mt-2">
            <span className="px-2 py-1 bg-green-200 font-semibold rounded-2xl text-green-500">
              3 online
            </span>
            <span className="px-2 py-1 bg-red-200 font-semibold rounded-2xl text-red-500">
              1 offline
            </span>
          </div>
        </Card.Footer>
      </Card>
    </div>
  );
};
