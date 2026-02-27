import { House, Grid2X2Icon, ScanBarcode, Network, Milk } from "lucide-react";
import Link from "next/link";
import { UserProfile } from "@/features/user/components/UserProfile";

const NAV_ITEMS = [
  { name: "Дашборд", href: "/", icon: House },
  { name: "Продукты", href: "/products", icon: Grid2X2Icon },
  { name: "Ненайденные штрихкоды", href: "/barcodes", icon: ScanBarcode },
  { name: "Ингредиенты", href: "/ingredients", icon: Milk },
  { name: "Таксономия", href: "/taxonomy", icon: Network },
];

export const Sidebar = () => {
  return (
    <aside className="flex flex-col justify-between w-64 min-h-screen border-r border-gray-200 py-3">
      <div>
        <div className="mb-5 px-5">
          <p className="font-semibold text-xl">Admin</p>
        </div>

        <nav>
          <ul className="space-y-1">
            {NAV_ITEMS.map((item) => {
              const Icon = item.icon;
              return (
                <li
                  key={item.name}
                  className="hover:bg-gray-100 px-5"
                >
                  <Link
                    href={item.href}
                    className="flex items-center gap-1 py-1"
                  >
                    <Icon size={20} />
                    <span className="truncate">{item.name}</span>
                  </Link>
                </li>
              );
            })}
          </ul>
        </nav>
      </div>
      <div className="px-5">
        <UserProfile />
      </div>
    </aside>
  );
};
