"use client";

import React, { useState } from "react";
import { Layout, Menu, Avatar, Dropdown, Button } from "antd";
import {
  DashboardOutlined,
  AppstoreOutlined,
  BarcodeOutlined,
  TagsOutlined,
  ApartmentOutlined,
  UserOutlined,
  LogoutOutlined,
  MenuFoldOutlined,
  MenuUnfoldOutlined,
} from "@ant-design/icons";
import { useRouter, usePathname } from "next/navigation";
import { cn } from "@/shared/lib/utils";

const { Sider, Content } = Layout;

interface DashboardLayoutProps {
  children: React.ReactNode;
}

export const Sidebar = ({ children }: DashboardLayoutProps) => {
  const [collapsed, setCollapsed] = useState(false);
  const router = useRouter();
  const pathname = usePathname();

  const menuItems = [
    { key: "/dashboard", icon: <DashboardOutlined />, label: "Дашборд" },
    { key: "/products", icon: <AppstoreOutlined />, label: "Каталог напитков" },
    {
      key: "/barcodes",
      icon: <BarcodeOutlined />,
      label: "Ненайденные штрихкоды",
    },
    { key: "/ingredients", icon: <TagsOutlined />, label: "Ингредиенты" },
    { key: "/taxonomy", icon: <ApartmentOutlined />, label: "Таксономия" },
  ];

  return (
    <Layout className="h-screen overflow-hidden">
      <Sider collapsible collapsed={collapsed} trigger={null} theme="light">
        <div className="flex flex-col h-full">
          <div
            className={cn(
              "h-16 flex items-center border-b border-[#f0f0f0] transition-all",
              collapsed ? "justify-center px-0" : "justify-between px-4",
            )}
          >
            {!collapsed && <h1 className="text-xl font-bold m-0">Admin</h1>}
            <Button
              type="text"
              icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
              onClick={() => setCollapsed(!collapsed)}
              className="flex items-center justify-center text-lg w-8 h-8"
            />
          </div>

          <div className="flex-1 overflow-y-auto overflow-x-hidden">
            <Menu
              theme="light"
              mode="inline"
              selectedKeys={[pathname]}
              items={menuItems}
              onClick={({ key }) => router.push(key)}
              className="border-none"
            />
          </div>

          <div
            className={cn(
              "flex p-4 border-t border-[#f0f0f0]",
              collapsed ? "justify-center" : "justify-start",
            )}
          >
            <Dropdown
              menu={{
                items: [
                  {
                    key: "logout",
                    label: "Выйти",
                    icon: <LogoutOutlined />,
                    danger: true,
                  },
                ],
              }}
              placement="topRight"
            >
              <div className="flex items-center gap-3 cursor-pointer w-full">
                <Avatar icon={<UserOutlined />} style={{ flexShrink: 0 }} />
                {!collapsed && (
                  <div className="overflow-hidden">
                    <div className="font-medium whitespace-nowrap text-ellipsis">
                      User
                    </div>
                    <div className="text-sm text-[#888]">Администратор</div>
                  </div>
                )}
              </div>
            </Dropdown>
          </div>
        </div>
      </Sider>

      <Layout className="overflow-y-auto overflow-x-hidden [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none]">
        <Content
          style={{
            margin: "0 40px",
          }}
        >
          {children}
        </Content>
      </Layout>
    </Layout>
  );
};
