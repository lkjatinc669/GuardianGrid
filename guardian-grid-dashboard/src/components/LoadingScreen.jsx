import React, { useEffect, useRef } from 'react';
import { gsap } from 'gsap';

const LoadingScreen = ({ onComplete }) => {
    const scope = useRef(null);
    const progressRef = useRef(null);
    const textRef = useRef(null);

    useEffect(() => {
        const ctx = gsap.context(() => {
            const tl = gsap.timeline({
                onComplete: () => {
                    gsap.to(scope.current, {
                        opacity: 0,
                        duration: 0.8,
                        ease: "power2.inOut",
                        onComplete: onComplete
                    });
                }
            });

            // 1. Initial State
            gsap.set(".boot-line", { opacity: 0, x: -10 });
            gsap.set(".logo-symbol", { scale: 0, rotation: -180 });

            // 2. Boot Sequence
            tl.to(".logo-symbol", { scale: 1, rotation: 0, duration: 1, ease: "back.out(1.7)" })
              .to(".boot-line", { 
                  opacity: 1, 
                  x: 0, 
                  stagger: 0.1, 
                  duration: 0.4, 
                  ease: "power1.out" 
              }, "-=0.5")
              .to(progressRef.current, {
                  width: "100%",
                  duration: 2,
                  ease: "none",
                  onUpdate: function() {
                      const p = Math.floor(this.progress() * 100);
                      if (textRef.current) textRef.current.innerText = `INITIALIZING_CORE: ${p}%`;
                  }
              })
              .to(".status-dot", {
                  backgroundColor: "#00ff9d",
                  boxShadow: "0 0 10px #00ff9d",
                  duration: 0.3
              });

        }, scope);

        return () => ctx.revert();
    }, [onComplete]);

    return (
        <div ref={scope} className="fixed inset-0 z-[9999] bg-[#020617] flex flex-col items-center justify-center font-mono overflow-hidden">
            {/* Background Grid Decoration */}
            <div className="absolute inset-0 opacity-10 pointer-events-none" 
                 style={{ backgroundImage: 'linear-gradient(#1e293b 1px, transparent 1px), linear-gradient(90deg, #1e293b 1px, transparent 1px)', backgroundSize: '40px 40px' }} />

            <div className="relative z-10 w-full max-w-md px-8 text-center">
                {/* Animated Logo */}
                <div className="logo-symbol w-16 h-16 bg-blue-600 mx-auto mb-10 rounded-xl flex items-center justify-center shadow-[0_0_30px_rgba(37,99,235,0.3)]">
                    <span className="text-white text-3xl font-bold">G</span>
                </div>

                {/* Boot Lines */}
                <div className="space-y-2 mb-8 text-left">
                    <div className="boot-line text-[9px] text-slate-500 tracking-widest">{">"} KERNEL_ATTACHED: OK</div>
                    <div className="boot-line text-[9px] text-slate-500 tracking-widest">{">"} CRYPTO_HANDSHAKE: VERIFIED</div>
                    <div className="boot-line text-[9px] text-slate-500 tracking-widest">{">"} GRID_LINK: SYNCING</div>
                </div>

                {/* Progress Section */}
                <div className="relative h-1 bg-slate-900 overflow-hidden mb-4">
                    <div ref={progressRef} className="absolute top-0 left-0 h-full w-0 bg-blue-500 shadow-[0_0_15px_rgba(59,130,246,0.5)]" />
                </div>

                <div className="flex justify-between items-center">
                    <div ref={textRef} className="text-[#00ff9d] text-[10px] tracking-widest uppercase">Initializing_Core: 0%</div>
                    <div className="status-dot w-2 h-2 bg-slate-800 rounded-full" />
                </div>
            </div>

            {/* Scanning Scanline */}
            <div className="absolute top-0 left-0 w-full h-1 bg-blue-500/10 blur-sm animate-scanline" />

            <style jsx>{`
                @keyframes scanline {
                    0% { top: 0% }
                    100% { top: 100% }
                }
                .animate-scanline {
                    animation: scanline 4s linear infinite;
                }
            `}</style>
        </div>
    );
};

export default LoadingScreen;
