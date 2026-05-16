import React, { useEffect, useState, useRef } from 'react';
import { gsap } from 'gsap';
import { useNavigate } from 'react-router-dom';
import Sidebar from '../components/Sidebar';

const Dashboard = () => {
    const navigate = useNavigate();
    const [agents, setAgents] = useState([]);
    const [loading, setLoading] = useState(true);
    const ws = useRef(null);

    // Initial Fetch
    const fetchOverview = async () => {
        try {
            const token = localStorage.getItem("token");
            if (!token) {
                console.warn("No auth token found, redirecting...");
                navigate("/auth/login");
                return;
            }

            const response = await fetch("http://localhost:8080/dashboard/overview", {
                headers: { "Authorization": `Bearer ${token}` }
            });

            if (response.status === 401) {
                localStorage.removeItem("auth");
                navigate("/auth/login");
                return;
            }

            const result = await response.json();
            if (result.status === "success") {
                setAgents(result.data || []);
            }
        } catch (error) {
            console.error("Failed to fetch dashboard data", error);
            // Fallback for visual testing if needed: setAgents([])
        } finally {
            setLoading(false);
        }
    };

    // WebSocket Setup
    useEffect(() => {
        fetchOverview();

        const connectWS = () => {
            ws.current = new WebSocket("ws://localhost:8080/ws");

            ws.current.onopen = () => {
                console.log("📡 Connected to GuardianGrid Stream");
            };

            ws.current.onmessage = (event) => {
                const message = JSON.parse(event.data);
                
                if (message.type === "agent_registered") {
                    setAgents(prev => {
                        if (prev.find(a => a.agent_id === message.agent_id)) return prev;
                        return [...prev, { agent_id: message.agent_id, data: {} }];
                    });
                }

                if (message.type === "telemetry_update") {
                    setAgents(prevAgents => {
                        const exists = prevAgents.find(a => a.agent_id === message.agent_id);
                        if (!exists) {
                            return [...prevAgents, { agent_id: message.agent_id, data: message.data }];
                        }

                        return prevAgents.map(agent => {
                            if (agent.agent_id === message.agent_id) {
                                return {
                                    ...agent,
                                    data: {
                                        ...agent.data,
                                        ...message.data
                                    }
                                };
                            }
                            return agent;
                        });
                    });

                    const safeId = message.agent_id.replace(/[^a-zA-Z0-9]/g, '');
                    const card = document.querySelector(`.agent-card-${safeId}`);
                    if (card) {
                        // Pulse the whole card
                        gsap.fromTo(card, 
                            { borderColor: "#00ff9d", boxShadow: "0 0 20px rgba(0,255,157,0.2)" }, 
                            { borderColor: "rgba(255,255,255,0.1)", boxShadow: "0 0 0px rgba(0,0,0,0)", duration: 1 }
                        );

                        // Pulse metrics specifically
                        const metrics = card.querySelectorAll('.metric-value');
                        gsap.fromTo(metrics,
                            { color: "#00ff9d", scale: 1.05 },
                            { color: "", scale: 1, duration: 0.5, stagger: 0.05 }
                        );
                    }
                }
            };

            ws.current.onclose = () => {
                console.log("🔌 Stream disconnected. Retrying...");
                setTimeout(connectWS, 3000);
            };
        };

        connectWS();
        return () => ws.current?.close();
    }, []);

    useEffect(() => {
        if (!loading && agents.length > 0) {
            gsap.fromTo(".agent-card", 
                { opacity: 0, y: 20, scale: 0.95 }, 
                { opacity: 1, y: 0, scale: 1, stagger: 0.1, duration: 0.8, ease: "expo.out" }
            );
        }
    }, [loading, agents.length]);

    if (loading) return (
        <div className="flex-1 flex h-screen items-center justify-center bg-[#020617] ml-64">
            <div className="text-[#00ff9d] font-mono text-xl animate-pulse tracking-[0.5em]">INITIALIZING_GRID...</div>
        </div>
    );

    return (
        <div className="flex min-h-screen bg-[#020617] text-slate-300 font-mono w-full">
            <Sidebar />
            
            <main className="flex-1 ml-64 p-8 overflow-y-auto w-full">
                <header className="mb-12 border-b border-white/5 pb-8 flex justify-between items-end">
                    <div>
                        <h1 className="text-4xl font-bold text-white tracking-tighter mb-2">COMMAND_CENTER</h1>
                        <p className="text-slate-500 text-xs tracking-[0.3em] uppercase">Real-time Operations Hub v1.0</p>
                    </div>
                    <div className="text-right">
                        <div className="text-[#00ff9d] text-2xl font-bold">{agents.length}</div>
                        <div className="text-slate-500 text-[10px] uppercase tracking-widest">Active_Nodes</div>
                    </div>
                </header>

                <div className="grid grid-cols-1 xl:grid-cols-2 2xl:grid-cols-3 gap-6">
                    {agents.length > 0 ? agents.map((agent) => {
                        const safeId = agent.agent_id.replace(/[^a-zA-Z0-9]/g, '');
                        return (
                            <div key={agent.agent_id} className={`agent-card agent-card-${safeId} relative bg-slate-900/40 border border-white/10 p-6 backdrop-blur-md overflow-hidden group hover:border-blue-500/50 transition-all duration-500`}>
                                {/* ... existing card content ... */}
                                <div className="absolute top-0 left-0 w-full h-0.5 bg-gradient-to-r from-transparent via-blue-500 to-transparent scale-x-0 group-hover:scale-x-100 transition-transform duration-700" />
                                
                                <div className="flex justify-between items-start mb-6">
                                    <div>
                                        <div className="text-[10px] text-slate-500 uppercase tracking-widest mb-1">Host_Node</div>
                                        <div className="text-white font-bold text-sm truncate max-w-[200px]">{agent.agent_id}</div>
                                    </div>
                                    <div className="flex items-center gap-2">
                                        <span className="w-2 h-2 bg-[#00ff9d] rounded-full animate-pulse" />
                                        <span className="text-[9px] text-[#00ff9d] font-bold uppercase tracking-tighter">Live</span>
                                    </div>
                                </div>

                                <div className="space-y-6">
                                    <div className="metric-box bg-slate-950/50 p-4 border border-white/5">
                                        <h3 className="text-[9px] text-slate-500 uppercase tracking-widest mb-3 border-b border-white/5 pb-2">Network_Forensics</h3>
                                        <div className="flex justify-between items-center mb-2">
                                            <span className="text-xs text-slate-400">Active_Conns</span>
                                            <span className="metric-value text-blue-400 text-sm font-bold">{agent.data?.network?.network_activity?.count || 0}</span>
                                        </div>
                                        <div className="flex justify-between items-center">
                                            <span className="text-xs text-slate-400">Throughput</span>
                                            <span className="metric-value text-blue-400 text-xs">{((agent.data?.network?.network_activity?.bytes_sent || 0) / 1024).toFixed(2)} KB/s</span>
                                        </div>
                                    </div>

                                    <div className="metric-box bg-slate-950/50 p-4 border border-white/5">
                                        <h3 className="text-[9px] text-slate-500 uppercase tracking-widest mb-3 border-b border-white/5 pb-2">Persistence_Audit</h3>
                                        <div className="flex justify-between items-center">
                                            <span className="text-xs text-slate-400">Autorun_Entries</span>
                                            <span className={`metric-value ${(agent.data?.persistence?.count || 0) > 10 ? 'text-red-400' : 'text-orange-400'} text-sm font-bold`}>
                                                {agent.data?.persistence?.count || 0}
                                            </span>
                                        </div>
                                    </div>

                                    <div className="metric-box bg-slate-950/50 p-4 border border-white/5">
                                        <h3 className="text-[9px] text-slate-500 uppercase tracking-widest mb-3 border-b border-white/5 pb-2">System_Integrity</h3>
                                        <div className="flex justify-between items-center">
                                            <span className="text-xs text-slate-400">Running_Procs</span>
                                            <span className="metric-value text-[#00ff9d] text-sm font-bold">{agent.data?.processes?.count || 0}</span>
                                        </div>
                                    </div>
                                </div>

                                <button 
                                    onClick={() => navigate(`/agent/${agent.agent_id}`)}
                                    className="mt-6 w-full py-3 bg-blue-600/10 border border-blue-500/20 text-blue-400 text-[10px] font-bold tracking-[0.3em] uppercase hover:bg-blue-600 hover:text-white transition-all duration-300"
                                >
                                    ACCESS_FORENSICS_NODE
                                </button>

                                <div className="absolute bottom-0 right-0 w-8 h-8 opacity-10 pointer-events-none">
                                    <div className="absolute bottom-0 right-0 w-full h-px bg-blue-500" />
                                    <div className="absolute bottom-0 right-0 h-full w-px bg-blue-500" />
                                </div>
                            </div>
                        );
                    }) : (
                        <div className="col-span-full py-20 text-center border border-dashed border-white/10 rounded-lg">
                            <div className="text-slate-600 uppercase tracking-[0.4em] text-sm">No_Active_Nodes_Identified</div>
                            <p className="text-slate-700 text-[10px] mt-4 uppercase">Waiting for agent registration handshake...</p>
                        </div>
                    )}
                </div>
            </main>
        </div>
    );
};

export default Dashboard;
