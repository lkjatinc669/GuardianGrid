import React, { useState, useEffect, useRef } from "react";
import { gsap } from "gsap";

export default function Stage10({ onSignup }) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
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

      // 2. Entrance Sequence
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
        stagger: 0.1, 
        duration: 0.8 
      }, "-=0.4");

      // Ambient Background Pulse
      gsap.to(".glow-sphere", {
        duration: 8,
        x: "random(-50, 50)",
        y: "random(-50, 50)",
        repeat: -1,
        yoyo: true,
        ease: "sine.inOut"
      });
    }, scope);

    return () => ctx.revert();
  }, []);

  const handleSignup = async () => {
    if (!username || !password) {
      setError("AUTHENTICATION_REQUIRED");
      // High-frequency shake for error
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
      // Outro animation
      gsap.to(".animate-out", { 
        opacity: 0, 
        scale: 0.95, 
        stagger: 0.05, 
        duration: 0.4, 
        onComplete: onSignup 
      });
    } catch {
      setError("SYSTEM_FAILURE_RETRY");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div ref={scope} className="min-h-screen flex items-center justify-center bg-[#020617] text-slate-200 overflow-hidden font-mono">
      
      {/* Dynamic Background */}
      <div className="glow-sphere absolute top-1/4 left-1/4 w-100 h-100 bg-blue-600/10 blur-[120px] rounded-full" />
      <div className="glow-sphere absolute bottom-1/4 right-1/4 w-100 h-100 bg-indigo-600/10 blur-[120px] rounded-full" />
      
      <div className="z-10 w-full max-w-130 px-6 card-container">
        
        {/* Header Section */}
        <div className="mb-8 space-y-1">
          <div ref={lineRef} className="h-px bg-blue-500/50 mb-4" />
          <h1 className="animate-in text-3xl font-bold tracking-tighter text-white">
            GUARDIANGRID <span className="text-blue-500 text-sm font-normal align-top">SIGNUP v1.0</span>
          </h1>
          <p className="animate-in text-slate-500 text-[10px] uppercase tracking-[0.2em]">
            Establish Secure Handshake
          </p>
        </div>

        {/* Main Card */}
        <div 
          ref={cardRef} 
          className="animate-out relative border border-white/5 p-8 rounded-sm shadow-2xl overflow-hidden"
        >
          {/* Animated Top Border */}
          <div ref={borderRef} className="absolute top-0 left-0 w-full h-0.5 bg-linear-to-r from-transparent via-blue-400 to-transparent" />

          {/* Form Fields */}
          <div className="space-y-6">
            <div className="animate-in">
              <label className="block text-[9px] text-slate-400 uppercase tracking-widest mb-2">Primary_Identity</label>
              <input
                type="text"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                className="w-full bg-slate-950/50 border border-white/10 px-4 py-3 text-sm focus:outline-none focus:border-blue-500/50 focus:bg-slate-900/80 transition-all rounded-sm"
                placeholder="USER_ID"
              />
            </div>

            <div className="animate-in">
              <label className="block text-[9px] text-slate-400 uppercase tracking-widest mb-2">Access_Cipher</label>
              <input
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                className="w-full bg-slate-950/50 border border-white/10 px-4 py-3 text-sm focus:outline-none focus:border-blue-500/50 focus:bg-slate-900/80 transition-all rounded-sm"
                placeholder="********"
              />
            </div>

            {error && (
              <div className="text-red-400 text-[10px] bg-red-400/5 border border-red-400/20 p-2 animate-pulse">
                &gt; ERROR: {error}
              </div>
            )}

            <button
              onClick={handleSignup}
              disabled={loading}
              className="animate-in relative w-full group overflow-hidden"
            >
              <div className="absolute inset-0 bg-blue-600 translate-y-[101%] group-hover:translate-y-0 transition-transform duration-300 ease-out" />
              <div className="relative border border-blue-500/50 py-4 text-xs font-bold tracking-[0.3em] group-hover:text-white transition-colors">
                {loading ? "INITIALIZING..." : "EXECUTE_SIGNUP"}
              </div>
            </button>
          </div>
        </div>

        {/* Footer Meta */}
        <div className="animate-in mt-6 flex justify-between items-center text-[8px] text-slate-600 uppercase tracking-widest">
          <span>Shield_Active</span>
          <span className="flex gap-2">
            <span className="w-1 h-1 bg-green-500 rounded-full animate-ping" />
            Node: 77-B
          </span>
        </div>

      </div>
    </div>
  );
}