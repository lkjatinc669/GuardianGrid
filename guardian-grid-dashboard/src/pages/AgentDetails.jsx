import React, { useEffect, useState, useRef } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { gsap } from 'gsap';
import Sidebar from '../components/Sidebar';
import { 
  ArrowLeft, 
  Activity, 
  Network, 
  ShieldAlert, 
  Package, 
  Cpu, 
  Globe 
} from 'lucide-react';

const AgentDetails = () => {
    const { id } = useParams();
    const navigate = useNavigate();
    const [agent, setAgent] = useState(null);
    const [loading, setLoading] = useState(true);
    const [activeTab, setActiveTab] = useState('network');
    const containerRef = useRef(null);

    const fetchAgentDetails = async () => {
        try {
            const token = localStorage.getItem("token");
            const response = await fetch(`http://localhost:8080/dashboard/agent/${id}`, {
                headers: { "Authorization": `Bearer ${token}` }
            });
            const result = await response.json();
            if (result.status === "success") {
                setAgent(result.data);
            }
        } catch (error) {
            console.error("Failed to fetch agent details", error);
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        fetchAgentDetails();
        const interval = setInterval(fetchAgentDetails, 10000);
        return () => clearInterval(interval);
    }, [id]);

    useEffect(() => {
        if (!loading && agent) {
            gsap.fromTo(".fade-in", 
                { opacity: 0, y: 15 }, 
                { opacity: 1, y: 0, stagger: 0.1, duration: 0.6, ease: "power2.out" }
            );
        }
    }, [loading, activeTab]);

    if (loading) return (
        <div className="flex h-screen items-center justify-center bg-[#020617]">
            <div className="text-[#00ff9d] font-mono text-xl animate-pulse tracking-[0.3em]">RECONSTRUCTING_NODE_DATA...</div>
        </div>
    );

    const renderTabContent = () => {
        const data = agent?.data || {};
        
        switch (activeTab) {
            case 'network':
                const net = data.network?.network_activity || {};
                const conns = net.connections || [];
                return (
                    <div className="space-y-6 fade-in">
                        <div className="grid grid-cols-4 gap-4 mb-6">
                            <div className="bg-slate-950/50 p-4 border border-white/5 rounded">
                                <div className="text-[10px] text-slate-500 uppercase mb-1">Bytes_Sent</div>
                                <div className="text-blue-400 font-bold">{net.bytes_sent || 0}</div>
                            </div>
                            <div className="bg-slate-950/50 p-4 border border-white/5 rounded">
                                <div className="text-[10px] text-slate-500 uppercase mb-1">Bytes_Recv</div>
                                <div className="text-blue-400 font-bold">{net.bytes_recv || 0}</div>
                            </div>
                            <div className="bg-slate-950/50 p-4 border border-white/5 rounded">
                                <div className="text-[10px] text-slate-500 uppercase mb-1">Packets_Sent</div>
                                <div className="text-blue-400 font-bold">{net.packets_sent || 0}</div>
                            </div>
                            <div className="bg-slate-950/50 p-4 border border-white/5 rounded">
                                <div className="text-[10px] text-slate-500 uppercase mb-1">Packets_Recv</div>
                                <div className="text-blue-400 font-bold">{net.packets_recv || 0}</div>
                            </div>
                        </div>
                        
                        <div className="bg-slate-900/40 border border-white/10 rounded-lg overflow-hidden">
                            <table className="w-full text-left text-xs">
                                <thead className="bg-white/5 text-slate-400 uppercase tracking-tighter">
                                    <tr>
                                        <th className="p-4">Local_Address</th>
                                        <th className="p-4">Remote_Address</th>
                                        <th className="p-4">Status</th>
                                        <th className="p-4">PID</th>
                                    </tr>
                                </thead>
                                <tbody className="divide-y divide-white/5">
                                    {conns.length > 0 ? conns.map((c, i) => (
                                        <tr key={i} className="hover:bg-white/5 transition-colors">
                                            <td className="p-4 text-slate-300 font-mono">{c.local_addr}</td>
                                            <td className="p-4 text-blue-400 font-mono">{c.remote_addr || "0.0.0.0:0"}</td>
                                            <td className="p-4">
                                                <span className={`px-2 py-0.5 rounded-full text-[9px] font-bold uppercase ${c.status === 'ESTABLISHED' ? 'bg-green-500/20 text-green-400' : 'bg-slate-700/50 text-slate-400'}`}>
                                                    {c.status}
                                                </span>
                                            </td>
                                            <td className="p-4 text-slate-500">{c.pid}</td>
                                        </tr>
                                    )) : (
                                        <tr><td colSpan="4" className="p-8 text-center text-slate-600">No active connections identified.</td></tr>
                                    )}
                                </tbody>
                            </table>
                        </div>
                    </div>
                );

            case 'persistence':
                const persistence = data.persistence?.persistence || [];
                return (
                    <div className="grid grid-cols-1 gap-4 fade-in">
                        {persistence.length > 0 ? persistence.map((p, i) => (
                            <div key={i} className="bg-slate-900/40 border border-white/10 p-4 flex justify-between items-center group hover:border-orange-500/30 transition-all">
                                <div>
                                    <div className="text-[10px] text-orange-400 font-bold uppercase tracking-widest mb-1">{p.type}</div>
                                    <div className="text-white text-sm font-bold mb-1">{p.name}</div>
                                    <div className="text-[10px] text-slate-500 font-mono truncate max-w-2xl">{p.location}</div>
                                </div>
                                <div className="text-right">
                                    <div className="text-[9px] text-slate-600 uppercase mb-1">Executable</div>
                                    <div className="text-xs text-slate-400 font-mono">{p.path || "N/A"}</div>
                                </div>
                            </div>
                        )) : (
                            <div className="p-12 text-center text-slate-600 border border-dashed border-white/10 rounded-lg">No persistence mechanisms detected.</div>
                        )}
                    </div>
                );

            case 'processes':
                const processes = data.processes?.processes || [];
                return (
                    <div className="grid grid-cols-2 gap-4 fade-in">
                        {processes.map((proc, i) => (
                            <div key={i} className="bg-slate-950/30 border border-white/5 p-3 rounded flex gap-4 items-center">
                                <div className="bg-blue-600/10 p-2 rounded text-blue-400">
                                    <Activity size={16} />
                                </div>
                                <div className="flex-1 overflow-hidden">
                                    <div className="text-xs text-white font-bold truncate">{proc.name}</div>
                                    <div className="text-[9px] text-slate-600 font-mono">PID: {proc.pid} | CPU: {proc.cpu_percent?.toFixed(1)}%</div>
                                </div>
                                <div className="text-right">
                                    <div className="text-[10px] text-[#00ff9d] font-bold">{(proc.memory_rss / (1024*1024)).toFixed(1)} MB</div>
                                </div>
                            </div>
                        ))}
                    </div>
                );

            case 'dns':
                const dns = data.dnscache?.dns_entries || [];
                return (
                    <div className="bg-slate-900/40 border border-white/10 rounded-lg overflow-hidden fade-in">
                        <table className="w-full text-left text-xs">
                            <thead className="bg-white/5 text-slate-400 uppercase tracking-tighter">
                                <tr>
                                    <th className="p-4">Hostname</th>
                                    <th className="p-4">IP_Address</th>
                                    <th className="p-4">Type</th>
                                </tr>
                            </thead>
                            <tbody className="divide-y divide-white/5">
                                {dns.length > 0 ? dns.map((d, i) => (
                                    <tr key={i} className="hover:bg-white/5 transition-colors">
                                        <td className="p-4 text-slate-300 font-mono">{d.hostname}</td>
                                        <td className="p-4 text-blue-400 font-mono">{d.ip}</td>
                                        <td className="p-4 text-slate-500 uppercase">{d.type || "A"}</td>
                                    </tr>
                                )) : (
                                    <tr><td colSpan="3" className="p-8 text-center text-slate-600">No DNS entries identified.</td></tr>
                                )}
                            </tbody>
                        </table>
                    </div>
                );

            case 'programs':
                const programs = data.programs?.programs || [];
                return (
                    <div className="grid grid-cols-1 gap-2 fade-in">
                        {programs.map((prog, i) => (
                            <div key={i} className="bg-slate-950/30 border border-white/5 p-3 flex justify-between items-center">
                                <div>
                                    <div className="text-xs text-white font-bold">{prog.name}</div>
                                    <div className="text-[9px] text-slate-600 uppercase tracking-widest">{prog.publisher || "Unknown Publisher"}</div>
                                </div>
                                <div className="text-[10px] text-blue-400 bg-blue-400/5 px-2 py-1 border border-blue-400/10">v{prog.version || "0.0.0"}</div>
                            </div>
                        ))}
                    </div>
                );

            case 'system':
                const pc = data.pcdata?.pc_data || {};
                return (
                    <div className="grid grid-cols-2 gap-6 fade-in">
                        <div className="bg-slate-900/40 border border-white/10 p-6">
                            <h4 className="text-[10px] text-slate-500 uppercase tracking-widest mb-4 border-b border-white/5 pb-2">Hardware_Specifications</h4>
                            <div className="space-y-4">
                                <div>
                                    <div className="text-[9px] text-slate-600 uppercase">CPU_Model</div>
                                    <div className="text-sm text-white">{pc.cpu_model} ({pc.cpu_cores} Cores)</div>
                                </div>
                                <div>
                                    <div className="text-[9px] text-slate-600 uppercase">Memory_Capacity</div>
                                    <div className="text-sm text-white">{pc.ram_total_mb} MB RAM</div>
                                </div>
                                <div>
                                    <div className="text-[9px] text-slate-600 uppercase">Architecture</div>
                                    <div className="text-sm text-white">{pc.architecture}</div>
                                </div>
                            </div>
                        </div>
                        <div className="bg-slate-900/40 border border-white/10 p-6">
                            <h4 className="text-[10px] text-slate-500 uppercase tracking-widest mb-4 border-b border-white/5 pb-2">OS_Environment</h4>
                            <div className="space-y-4">
                                <div>
                                    <div className="text-[9px] text-slate-600 uppercase">Platform</div>
                                    <div className="text-sm text-white">{pc.os} ({pc.platform_ver})</div>
                                </div>
                                <div>
                                    <div className="text-[9px] text-slate-600 uppercase">Kernel_Version</div>
                                    <div className="text-sm text-white">{pc.kernel_version}</div>
                                </div>
                                <div>
                                    <div className="text-[9px] text-slate-600 uppercase">Hostname</div>
                                    <div className="text-sm text-white">{pc.hostname}</div>
                                </div>
                            </div>
                        </div>
                    </div>
                );

            default:
                return <div className="text-slate-500 italic p-8 text-center">Module data reconstruction pending...</div>;
        }
    };

    return (
        <div ref={containerRef} className="flex min-h-screen bg-[#020617] text-slate-300 font-mono">
            <Sidebar />
            
            <main className="flex-1 ml-64 p-8 overflow-y-auto">
                <header className="mb-8 flex justify-between items-center">
                    <button onClick={() => navigate("/")} className="flex items-center gap-2 text-slate-500 hover:text-white transition-colors text-xs uppercase tracking-widest">
                        <ArrowLeft size={16} /> Back_To_Grid
                    </button>
                    <div className="text-right">
                        <div className="text-[10px] text-slate-500 uppercase tracking-widest">Target_Host</div>
                        <div className="text-white font-bold text-lg">{id}</div>
                    </div>
                </header>

                <div className="flex gap-2 mb-8 border-b border-white/5 pb-4 overflow-x-auto">
                    {[
                        { id: 'network', label: 'Network', icon: <Network size={14} /> },
                        { id: 'persistence', label: 'Persistence', icon: <ShieldAlert size={14} /> },
                        { id: 'processes', label: 'Processes', icon: <Activity size={14} /> },
                        { id: 'programs', label: 'Programs', icon: <Package size={14} /> },
                        { id: 'dns', label: 'DNS_Cache', icon: <Globe size={14} /> },
                        { id: 'system', label: 'Hardware', icon: <Cpu size={14} /> },
                    ].map(tab => (
                        <button 
                            key={tab.id}
                            onClick={() => setActiveTab(tab.id)}
                            className={`flex items-center gap-2 px-4 py-2 text-[10px] font-bold uppercase tracking-widest transition-all rounded ${activeTab === tab.id ? 'bg-blue-600 text-white shadow-[0_0_15px_rgba(37,99,235,0.3)]' : 'text-slate-500 hover:bg-white/5 hover:text-slate-300'}`}
                        >
                            {tab.icon} {tab.label}
                        </button>
                    ))}
                </div>

                <div className="tab-viewport">
                    {renderTabContent()}
                </div>
            </main>
        </div>
    );
};

export default AgentDetails;
