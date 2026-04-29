import React, { useState, useEffect, useRef } from "react";
import { gsap } from "gsap";

export default function Stage11({ onSuccess }) {
  const [form, setForm] = useState({
    username: "",
    password: "",
    server: ""
  });

  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  // Refs for precise targeting
  const scope = useRef(null);
  const cardRef = useRef(null);
  const borderRef = useRef(null);
  const lineRef = useRef(null);

  useEffect(() => {
    const ctx = gsap.context(() => {
      const tl = gsap.timeline({ defaults: { ease: "expo.out" } });

      // 1. Initial State
      gsap.set(".animate-in", { y: 30, opacity: 0 });
      gsap.set(borderRef.current, { scaleX: 0 });
      gsap.set(lineRef.current, { width: 0 });

      // 2. Entrance Sequence (Matching Previous Logic)
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

      // Ambient Background Pulse
      gsap.to(".glow-sphere", {
        duration: 8,
        x: "random(-60, 60)",
        y: "random(-60, 60)",
        repeat: -1,
        yoyo: true,
        ease: "sine.inOut"
      });
    }, scope);

    return () => ctx.revert();
  }, []);

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async () => {
    if (!form.username || !form.password) {
      setError("AUTHENTICATION_REQUIRED");
      gsap.fromTo(".card-container", 
        { x: -4 }, 
        { x: 4, duration: 0.05, repeat: 5, yoyo: true, onComplete: () => gsap.set(".card-container", { x: 0 }) }
      );
      return;
    }

    setLoading(true);
    setError("");

    try {
      await new Promise((res) => setTimeout(res, 2000));
      localStorage.setItem("auth", "true");
      
      // Outro animation before transition
      gsap.to(".animate-out", { 
        opacity: 0, 
        scale: 0.98, 
        stagger: 0.05, 
        duration: 0.4, 
        onComplete: onSuccess 
      });
    } catch {
      setError("SYSTEM_FAILURE_RETRY");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div ref={scope} className="min-h-screen flex items-center justify-center bg-[#020617] text-slate-200 overflow-hidden font-mono">
      
      {/* Background Orbs */}
      <div className="glow-sphere absolute top-1/3 left-1/4 w-125 h-125 bg-blue-600/5 blur-[140px] rounded-full" />
      <div className="glow-sphere absolute bottom-1/3 right-1/4 w-100 h-100 bg-indigo-600/5 blur-[120px] rounded-full" />
      
      <div className="z-10 w-full max-w-130 px-6 card-container">
        
        {/* Header Section */}
        <div className="mb-8 space-y-1">
          <div ref={lineRef} className="h-px bg-blue-500/50 mb-4" />
          <h1 className="animate-in text-3xl font-bold tracking-tighter text-white">
            GUARDIANGRID <span className="text-blue-500 text-sm font-normal align-top">LOGIN v1.0</span>
          </h1>
          <p className="animate-in text-slate-500 text-[10px] uppercase tracking-[0.3em]">
            Identity Verification Node
          </p>
        </div>

        {/* Form Card */}
        <div 
          ref={cardRef} 
          className="animate-out relative border border-white/10 p-8 shadow-2xl overflow-hidden"
        >
          {/* Top Animated Border */}
          <div ref={borderRef} className="absolute top-0 left-0 w-full h-0.5 bg-linear-to-r from-transparent via-blue-400 to-transparent" />

          <div className="space-y-5">
            {/* Username */}
            <div className="animate-in">
              <label className="block text-[9px] text-slate-400 uppercase tracking-widest mb-2">Subject_ID</label>
              <input
                type="text"
                name="username"
                value={form.username}
                onChange={handleChange}
                className="w-full bg-slate-950/50 border border-white/5 px-4 py-3 text-sm focus:outline-none focus:border-blue-500/50 focus:bg-slate-900/80 transition-all"
                placeholder="ENTRY_USER"
              />
            </div>

            {/* Password */}
            <div className="animate-in">
              <label className="block text-[9px] text-slate-400 uppercase tracking-widest mb-2">Auth_Cipher</label>
              <input
                type="password"
                name="password"
                value={form.password}
                onChange={handleChange}
                className="w-full bg-slate-950/50 border border-white/5 px-4 py-3 text-sm focus:outline-none focus:border-blue-500/50 focus:bg-slate-900/80 transition-all"
                placeholder="********"
              />
            </div>

            {/* Error Message */}
            {error && (
              <div className="text-red-400 text-[10px] bg-red-400/5 border border-red-400/20 p-2 text-center italic">
                {"> "} {error}
              </div>
            )}

            {/* Action Button */}
            <button
              onClick={handleSubmit}
              disabled={loading}
              className="animate-in relative w-full group overflow-hidden"
            >
              <div className="absolute inset-0 bg-blue-600 translate-y-[101%] group-hover:translate-y-0 transition-transform duration-300 ease-out" />
              <div className="relative border border-blue-500/40 py-4 text-xs font-bold tracking-[0.4em] group-hover:text-white transition-colors">
                {loading ? "LINKING_TO_GRID..." : "EXECUTE_ACCESS"}
              </div>
            </button>
          </div>
        </div>

        {/* Bottom Metadata */}
        <div className="animate-in mt-6 flex justify-between items-center text-[8px] text-slate-600 uppercase tracking-widest">
          <span>Shield_Active</span>
          <span className="flex items-center gap-2">
            <span className="w-1.5 h-1.5 bg-blue-500/40 rounded-full animate-pulse" />
            SECURE_PORT_443
          </span>
        </div>

      </div>
    </div>
  );
}