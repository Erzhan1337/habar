import { ReactNode } from "react";
import { cn } from "@/shared/lib/utils";

export const CardInfo = ({ children }: { children: ReactNode }) => {
  return (
    <div className="rounded-lg h-36 shadow-lg border-2 border-gray-200 p-4">
      {children}
    </div>
  );
};

CardInfo.Header = function CardHeader({
  title,
  icon,
}: {
  title: string;
  icon: ReactNode;
}) {
  return (
    <div className="flex items-center justify-between">
      <h3 className="text-[#71717A] text-sm font-medium">{title}</h3>
      <div className="text-gray-500">{icon}</div>
    </div>
  );
};

CardInfo.Value = function CardValue({ children }: { children: ReactNode }) {
  return <div className="text-2xl font-semibold mt-4">{children}</div>;
};

CardInfo.Footer = function CardFooter({
  children,
  className,
}: {
  children: ReactNode;
  className?: string;
}) {
  return (
    <div className={cn("text-xs font-medium mt-1", className)}>{children}</div>
  );
};
