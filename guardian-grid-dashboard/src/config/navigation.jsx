import { 
  LayoutDashboard, 
  ShieldCheck, 
  Database, 
  Settings, 
  Activity, 
  Terminal, 
  FileText, 
  Layers, 
  HardDrive 
} from 'lucide-react';

const sidebarData = [
  { 
    title: "Dashboard", 
    path: "/", 
    icon: <LayoutDashboard size={20} /> 
  },
  {
    title: "Security",
    path: "/security",
    icon: <ShieldCheck size={20} />,
    subNav: [
      { title: "Active Monitoring", path: "/security/monitoring", icon: <Activity size={18} /> },
      { title: "IPS Config", path: "/security/ips", icon: <Terminal size={18} /> },
      { title: "Threat Logs", path: "/security/logs", icon: <FileText size={18} /> },
    ],
  },
  {
    title: "Databases",
    path: "/db",
    icon: <Database size={20} />,
    subNav: [
      { title: "PostgreSQL", path: "/db/postgres", icon: <Layers size={18} /> },
      { title: "SQLite", path: "/db/sqlite", icon: <HardDrive size={18} /> },
    ],
  },
  { 
    title: "Settings", 
    path: "/settings", 
    icon: <Settings size={20} /> 
  },
];

export default sidebarData;