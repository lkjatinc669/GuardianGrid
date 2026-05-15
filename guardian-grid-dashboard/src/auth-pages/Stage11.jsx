import React, { useState, useEffect, useRef } from "react";
import { gsap } from "gsap";
import { useNavigate } from "react-router-dom";

export default function Stage11() {
  const navigate = useNavigate();
  const [form, setForm] = useState({
    username: "",
    code: "", // API expects 'code' for TOTP
  });

  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  const scope = useRef(null);
  const cardRef = useRef(null);
  const borderRef = useRef(null);
  const lineRef = useRef(null);

  useEffect(() => {
    const ctx = gsap.context(() => {
      const tl = gsap.timeline({ defaults: { ease: "expo.out" } });

      gsap.set(".animate-in", { y: 30, opacity: 0 });
      gsap.set(borderRef.current, { scaleX: 0 });
      gsap.set(lineRef.current, { width: 0 });

      tl.to(cardRef.current, { 
        duration: 1.5, 
        backgroundColor: "rgba(15, 23, 42, 0.8)", 
        backdropFilter: "blur(12px)" 
      })
      .to(borderRef.current, { scaleX: 1, duration: 1 }, "-=1")
      .to(lineRef.current, { width: "100%", duration: 0.8, ease: "power2.inOut" }, "-=0.5")
      .to(".animate-in", { 
        y: 0, 
        opacity: 1, 
        stagger: 0.08, 
        duration: 0.8 
      }, "-=0.4");
    }, scope);

    return () => ctx.revert();
  }, []);

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async () => {
    if (!form.username || !form.code) {
      setError("AUTHENTICATION_REQUIRED");
      return;
    }

    setLoading(true);
    setError("");

    try {
      const response = await fetch("http://localhost:8080/auth/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(form)
      });

      const result = await response.json();

      if (result.status === "success") {
        localStorage.setItem("token", result.data.token);
        localStorage.setItem("auth", "true");
        
        gsap.to(".animate-out", { 
          opacity: 0, 
          scale: 0.98, 
          stagger: 0.05, 
          duration: 0.4, 
          onComplete: () => navigate("/") 
        });
      } else {
        setError(result.message || "INVALID_CREDENTIALS");
      }
    } catch (err) {
      setError("SERVER_UNREACHABLE");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div ref={scope} className="min-h-screen flex items-center justify-center bg-[#020617] text-slate-200 overflow-hidden font-mono">
      <div className="glow-sphere absolute top/3 left-1/4 w-125 h-125 bg-blue-600/5 blur-[140px] rounded-full" />
      <div className="glow-sphere absolute bottom-1/3 right-1/4 w-100 h-100 bg-indigo-600/5 blur-[120px] rounded-full" />
      
      <div className="z-10 w-full max-w-130 px-6 card-container">
        <div className="mb-8 space-y-1">
          <div ref={lineRef} className="h-px bg-blue-500/50 mb-4" />
          <h1 className="animate-in text-3xl font-bold tracking-tighter text-white">
            GUARDIANGRID <span className="text-blue-500 text-sm font-normal align-top">LOGIN v1.0</span>
          </h1>
          <p className="animate-in text-slate-500 text-[10px] uppercase tracking-[0.3em]">
            Identity Verification Node
          </p>
        </div>

        <div ref={cardRef} className="animate-out relative border border-white/10 p-8 shadow-2xl overflow-hidden">
          <div ref={borderRef} className="absolute top-0 left-0 w-full h-0.5 bg-gradient-to-r from-transparent via-blue-400 to-transparent" />

          <div className="space-y-5">
            <div className="animate-in">
              <label className="block text-[9px] text-slate-400 uppercase tracking-widest mb-2">Username</label>
              <input
                type="text"
                name="username"
                value={form.username}
                onChange={handleChange}
                className="w-full bg-slate-950/50 border border-white/5 px-4 py-3 text-sm focus:outline-none focus:border-blue-500/50 focus:bg-slate-900/80 transition-all"
                placeholder="ENTRY_USER"
              />
            </div>

            <div className="animate-in">
              <label className="block text-[9px] text-slate-400 uppercase tracking-widest mb-2">TOTP Code</label>
              <input
                type="text"
                name="code"
                value={form.code}
                onChange={handleChange}
                className="w-full bg-slate-950/50 border border-white/5 px-4 py-3 text-sm focus:outline-none focus:border-blue-500/50 focus:bg-slate-900/80 transition-all"
                placeholder="000000"
              />
            </div>

            {error && (
              <div className="text-red-400 text-[10px] bg-red-400/5 border border-red-400/20 p-2 text-center italic">
                {"> "} {error}
              </div>
            )}

            <button onClick={handleSubmit} disabled={loading} className="animate-in relative w-full group overflow-hidden">
              <div className="absolute inset-0 bg-blue-600 translate-y-[101%] group-hover:translate-y-0 transition-transform duration-300 ease-out" />
              <div className="relative border border-blue-500/40 py-4 text-xs font-bold tracking-[0.4em] group-hover:text-white transition-colors">
                {loading ? "LINKING_TO_GRID..." : "EXECUTE_ACCESS"}
              </div>
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}
