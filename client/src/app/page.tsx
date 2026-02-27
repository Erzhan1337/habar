import { InfoGrid } from "@/features/dashboard/components/InfoGrid";
import { SystemErrorsChart } from "@/features/dashboard/components/Chart";
import {
  SystemLogsTable,
} from "@/features/dashboard/components/LogsTable";

export default function Home() {
  return (
    <div className="p-5">
      <InfoGrid />
      <SystemErrorsChart />
      <SystemLogsTable />
    </div>
  );
}
