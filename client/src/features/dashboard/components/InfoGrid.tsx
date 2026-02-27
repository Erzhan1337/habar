import { CardInfo } from "@/shared/ui/card-info";
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
      <CardInfo>
        <CardInfo.Header title="Всего продуктов" icon={<PackageOpen size={30} />} />
        <CardInfo.Value>1,324</CardInfo.Value>
        <CardInfo.Footer className="text-green-500">+84 за неделю</CardInfo.Footer>
      </CardInfo>

      <CardInfo>
        <CardInfo.Header
          title="Ненайденные штрихкоды"
          icon={<CircleAlert size={30} />}
        />
        <CardInfo.Value>47</CardInfo.Value>
        <CardInfo.Footer className="text-red-500">+12 в ожидании</CardInfo.Footer>
      </CardInfo>

      <CardInfo>
        <CardInfo.Header title="Продукты с сахаром" icon={<Candy size={30} />} />
        <CardInfo.Value>73.2%</CardInfo.Value>
        <CardInfo.Footer className="text-gray-400">969 продуктов</CardInfo.Footer>
      </CardInfo>

      <CardInfo>
        <CardInfo.Header title="Продукты с кофеином" icon={<Coffee size={30} />} />
        <CardInfo.Value>42.8%</CardInfo.Value>
        <CardInfo.Footer className="text-gray-400">969 продуктов</CardInfo.Footer>
      </CardInfo>

      <CardInfo>
        <CardInfo.Header
          title="Устаревшие запросы"
          icon={<ClockFading size={30} />}
        />
        <CardInfo.Value>28</CardInfo.Value>
        <CardInfo.Footer className="text-orange-400">
          Требует обновление
        </CardInfo.Footer>
      </CardInfo>

      <CardInfo>
        <CardInfo.Header
          title="Статус API провайдеров"
          icon={<TrendingUp size={30} />}
        />
        <CardInfo.Footer>
          <div className="flex items-center gap-2 mt-2">
            <span className="px-2 py-1 bg-green-200 font-semibold rounded-2xl text-green-500">
              3 online
            </span>
            <span className="px-2 py-1 bg-red-200 font-semibold rounded-2xl text-red-500">
              1 offline
            </span>
          </div>
        </CardInfo.Footer>
      </CardInfo>
    </div>
  );
};
