import { Sidebar } from "@/shared/ui/sidebar";
import { ReactNode } from "react";

export default function ProtectedLayout({ children }: { children: ReactNode }) {
  return <Sidebar>{children}</Sidebar>;
}
