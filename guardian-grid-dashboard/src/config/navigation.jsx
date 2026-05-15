import { 
  LayoutDashboard, 
  ShieldCheck, 
  Settings, 
  Activity, 
  Terminal, 
  FileText
} from 'lucide-react';

const sidebarData = [
  { 
    title: "Grid Overview", 
    path: "/", 
    icon: <LayoutDashboard size={20} /> 
  },
  {
    title: "Security Node",
    path: "/security",
    icon: <ShieldCheck size={20} />,
    subNav: [
      { title: "Active Monitoring", path: "/", icon: <Activity size={18} /> },
      { title: "Terminal Access", path: "/", icon: <Terminal size={18} /> },
      { title: "Event Logs", path: "/", icon: <FileText size={18} /> },
    ],
  },
  { 
    title: "Global Settings", 
    path: "/", 
    icon: <Settings size={20} /> 
  },
];

export default sidebarData;
