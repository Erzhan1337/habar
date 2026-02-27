"use client";
import { usePathname } from "next/navigation";

const paths: Record<string, string> = {
  "/": "Дашборд",
  "/products": "Продукты",
  "/barcodes": "Ненайденные штрихкоды",
  "/ingredients": "Ингредиенты",
  "/taxonomy": "Таксономия",
};

export const Header = () => {
  const path = usePathname();
  return (
    <header className="h-16 p-5 font-bold text-xl bg-[#F8F9FA]">
      {paths[path]}
    </header>
  );
};
