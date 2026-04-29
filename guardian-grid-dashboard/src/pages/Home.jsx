import React from "react";

const StatCard = ({ title, value, color, icon }) => {
  return (
    <div className="bg-slate-900 border border-slate-800 rounded-xl p-4 flex items-center justify-between hover:border-slate-700 transition">
      <div>
        <p className="text-sm text-gray-400">{title}</p>
        <h3 className={`text-2xl font-semibold ${color}`}>{value}</h3>
      </div>
      <div className="text-gray-400">{icon}</div>
    </div>
  );
};

const LogItem = ({ type, message, time }) => {
  const colors = {
    critical: "text-red-500",
    warning: "text-yellow-400",
    info: "text-blue-400",
  };

  return (
    <div className="flex justify-between text-sm py-2 border-b border-slate-800">
      <span className={`${colors[type]} font-medium`}>
        [{type.toUpperCase()}]
      </span>
      <span className="flex-1 ml-3 text-gray-300">{message}</span>
      <span className="text-gray-500">{time}</span>
    </div>
  );
};

export default function Dashboard() {
  return (
    <div className="p-6 text-white space-y-6">

      {/* Header */}
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold">Security Dashboard</h1>
        <button className="bg-blue-600 hover:bg-blue-700 px-4 py-2 rounded-lg">
          Run Scan
        </button>
      </div>

      {/* Stats */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
        <StatCard
          title="Threats Detected"
          value="23"
          color="text-red-500"
          icon="⚠️"
        />
        <StatCard
          title="Active Monitoring"
          value="ON"
          color="text-green-400"
          icon="🟢"
        />
        <StatCard
          title="Systems Protected"
          value="12"
          color="text-blue-400"
          icon="🖥️"
        />
        <StatCard
          title="Logs Today"
          value="1,248"
          color="text-purple-400"
          icon="📄"
        />
      </div>

      {/* Main Grid */}
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">

        {/* Threat Activity */}
        <div className="lg:col-span-2 bg-slate-900 border border-slate-800 rounded-xl p-4">
          <h2 className="text-lg font-semibold mb-4">Threat Activity</h2>

          <div className="h-48 flex items-center justify-center text-gray-500">
            {/* Replace later with chart */}
            Chart Placeholder
          </div>
        </div>

        {/* System Status */}
        <div className="bg-slate-900 border border-slate-800 rounded-xl p-4">
          <h2 className="text-lg font-semibold mb-4">System Status</h2>

          <ul className="space-y-3 text-sm">
            <li className="flex justify-between">
              Firewall <span className="text-green-400">Active</span>
            </li>
            <li className="flex justify-between">
              IDS <span className="text-green-400">Running</span>
            </li>
            <li className="flex justify-between">
              Database <span className="text-yellow-400">Degraded</span>
            </li>
            <li className="flex justify-between">
              API <span className="text-green-400">Online</span>
            </li>
          </ul>
        </div>
      </div>

      {/* Logs */}
      <div className="bg-slate-900 border border-slate-800 rounded-xl p-4">
        <h2 className="text-lg font-semibold mb-4">Recent Threat Logs</h2>

        <div className="max-h-60 overflow-y-auto">
          <LogItem
            type="critical"
            message="SQL Injection attempt blocked"
            time="2 min ago"
          />
          <LogItem
            type="warning"
            message="Multiple failed login attempts"
            time="10 min ago"
          />
          <LogItem
            type="info"
            message="System scan completed"
            time="30 min ago"
          />
        </div>
      </div>
    </div>
  );
}
