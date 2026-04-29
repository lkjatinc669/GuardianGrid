import React, { useRef, useState, useEffect } from "react";
import { invoke } from "@tauri-apps/api/core";
import { gsap } from "gsap";

export default function Stage1({ onSuccess }) {
  const [values, setValues] = useState(["", "", "", ""]);
  const [port, setPort] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  // Refs for GSAP
  const scope = useRef(null);
  const cardRef = useRef(null);
  const lineRef = useRef(null);
  const borderRef = useRef(null);
  const inputs = useRef([]);

  useEffect(() => {
    const ctx = gsap.context(() => {
      const tl = gsap.timeline({ defaults: { ease: "expo.out" } });

      // Initial state
      gsap.set(".animate-in", { y: 20, opacity: 0 });
      gsap.set(borderRef.current, { scaleX: 0 });
      gsap.set(lineRef.current, { width: 0 });

      // Entrance Sequence
      tl.to(cardRef.current, {
        duration: 1.2,
        backgroundColor: "rgba(15, 23, 42, 0.8)",
        backdropFilter: "blur(12px)"
      })
        .to(borderRef.current, { scaleX: 1, duration: 1 }, "-=1")
        .to(lineRef.current, { width: "100%", duration: 0.8 }, "-=0.5")
        .to(".animate-in", {
          y: 0,
          opacity: 1,
          stagger: 0.05,
          duration: 0.6
        }, "-=0.4");

      // Ambient Background
      gsap.to(".glow-sphere", {
        x: "random(-40, 40)",
        y: "random(-40, 40)",
        duration: 6,
        repeat: -1,
        yoyo: true,
        ease: "sine.inOut"
      });
    }, scope);

    return () => ctx.revert();
  }, []);

  const handleChange = (value, index) => {
    if (!/^\d*$/.test(value)) return;
    if (value.length > 3) return;
    if (Number(value) > 255) return;

    const newValues = [...values];
    newValues[index] = value;
    setValues(newValues);

    if (value.length === 3 && index < 3) {
      inputs.current[index + 1]?.focus();
    }

    if (Number(`${value}0`) > 255) {
      inputs.current[index + 1]?.focus();
    }
  };

  const handleKeyDown = (e, index) => {
    if (e.key === "Backspace" && !values[index] && index > 0) {
      inputs.current[index - 1]?.focus();
    }

    if (e.key === ".") {
      inputs.current[index + 1]?.focus();
    }
  };

  const handleSubmit = async () => {
    if (values.some(v => v === "")) {
      setError("INCOMPLETE_ADDRESS_STRING");
      gsap.fromTo(".card-container", { x: -4 }, { x: 4, duration: 0.05, repeat: 5, yoyo: true });
      return;
    }

    setLoading(true);
    setError("");

    try {
      await invoke("scan_host", {
        target: values.join("."),
        port: port ? Number(port) : 80,
      });

      gsap.to(".animate-out", {
        opacity: 0,
        y: -20,
        stagger: 0.05,
        onComplete: onSuccess
      });
    } catch {
      setError("CONNECTION_TIMEOUT_FAILURE");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div ref={scope} className="min-h-screen bg-[#020617] text-slate-200 flex items-center justify-center font-mono overflow-hidden">

      {/* Background Depth */}
      <div className="glow-sphere absolute top-1/4 left-1/4 w-125 h-125 bg-blue-600/5 blur-[120px] rounded-full" />
      <div className="glow-sphere absolute bottom-1/4 right-1/3 w-100 h-100 bg-indigo-600/5 blur-[100px] rounded-full" />

      <div className="z-10 w-full max-w-130 px-6 card-container">

        {/* Header Section */}
        <div className="mb-10 space-y-1">
          <div ref={lineRef} className="h-px bg-blue-500/50 mb-4" />
          <h1 className="animate-in text-4xl font-bold tracking-tighter text-white">
            GUARDIANGRID <span className="text-blue-500 text-sm font-normal align-top">IP ENTRY v1.0</span>
          </h1>
          <p className="animate-in text-slate-500 text-[10px] uppercase tracking-[0.4em]">
            Establish Hardware-Level Handshake
          </p>
        </div>

        {/* Input Card */}
        <div
          ref={cardRef}
          className="animate-out relative border border-white/5 p-8 shadow-2xl overflow-hidden"
        >
          {/* Animated Top Border */}
          <div ref={borderRef} className="absolute top-0 left-0 w-full h-0.5 bg-linear-to-r from-transparent via-blue-400 to-transparent" />

          <div className="space-y-8">
            <div className="animate-in">
              <label className="block text-[9px] text-slate-400 uppercase tracking-widest mb-4 italic">Target_IP_Address_Vector</label>

              <div className="flex items-center gap-1">
                {values.map((val, i) => (
                  <React.Fragment key={i}>
                    <div className="flex-1">
                      <input
                        ref={(el) => (inputs.current[i] = el)}
                        value={val}
                        onChange={(e) => handleChange(e.target.value, i)}
                        onKeyDown={(e) => handleKeyDown(e, i)}
                        maxLength={3}
                        placeholder=""
                        className="w-16 bg-slate-950/50 border border-white/5 px-0.1 py-4 text-center text-xl font-bold focus:outline-none focus:border-blue-500/50 focus:bg-slate-900/80 transition-all text-blue-400"
                      />
                    </div>
                    {i < 3 && <span className="text-slate-400 font-bold">.</span>}
                  </React.Fragment>
                ))}

                <span className="text-slate-700 font-bold">:</span>

                <div className="w-20">
                  <input
                    value={port}
                    onChange={(e) => {
                      const value = e.target.value;
                      if (!/^\d*$/.test(value)) return;
                      if (value.length > 5) return;
                      const num = Number(value);
                      if (num > 65535) {
                        setPort("65535");
                        return;
                      }
                      setPort(value)
                    }}
                    placeholder="80"
                    className="w-full bg-slate-950/50 border border-white/5 px-2 py-4 text-center text-xl font-bold focus:outline-none focus:border-purple-500/50 focus:bg-slate-900/80 transition-all text-purple-400"
                  />
                </div>
              </div>
            </div>

            {/* Error Display */}
            {error && (
              <div className="animate-in text-red-400 text-[10px] bg-red-400/5 border border-red-400/20 p-3 italic">
                {"> "} ERROR: {error}
              </div>
            )}

            {/* Submit Action */}
            <button
              onClick={handleSubmit}
              disabled={loading}
              className="animate-in relative w-full group overflow-hidden"
            >
              <div className="absolute inset-0 bg-blue-600 translate-y-[101%] group-hover:translate-y-0 transition-transform duration-300 ease-out" />
              <div className="relative border border-blue-500/40 py-4 text-[11px] font-bold tracking-[0.6em] group-hover:text-white transition-colors uppercase">
                {loading ? "SCANNING_HOST..." : "EXECUTE_SCAN"}
              </div>
            </button>
          </div>
        </div>

        {/* Footer Meta */}
        <div className="animate-in mt-8 flex justify-between items-center text-[9px] text-slate-600 font-bold uppercase tracking-widest px-2">
          <span className="flex items-center gap-2">
            <span className="w-1.5 h-1.5 bg-blue-500 rounded-full animate-pulse shadow-[0_0_8px_#3b82f6]" />
            GRID_CONNECTION_IDLE
          </span>
          <span className="text-slate-800 tracking-tighter">PORT_MAP_V2</span>
        </div>

      </div>
    </div>
  );
}