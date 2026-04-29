import React, { useState, useRef, useEffect } from "react";
import { NavLink, useLocation } from "react-router-dom";
import { gsap } from "gsap";
import { ChevronDown } from "lucide-react";
import sidebarData from "../config/navigation.jsx";
import Logo from "../assets/rocket.svg"

const SubMenu = ({ item }) => {
  const [open, setOpen] = useState(false);
  const submenuRef = useRef(null);
  const arrowRef = useRef(null);
  const { pathname } = useLocation();

  const hasSub = !!item.subNav;
  
  // Check if a child route is currently active to keep the menu open on refresh
  const isChildActive = hasSub && item.subNav.some(sub => sub.path === pathname);

  useEffect(() => {
    if (isChildActive) setOpen(true);
  }, [isChildActive]);

  const toggleMenu = () => {
    if (!hasSub) return;

    if (!open) {
      setOpen(true);
      // Animate Open
      gsap.to(submenuRef.current, {
        height: "auto",
        opacity: 1,
        duration: 0.4,
        ease: "power2.out"
      });
      // Rotate Arrow
      gsap.to(arrowRef.current, { rotate: 180, duration: 0.3 });
      // Stagger children
      gsap.fromTo(
        submenuRef.current.children,
        { y: -10, opacity: 0 },
        { y: 0, opacity: 1, stagger: 0.05, duration: 0.3, delay: 0.1 }
      );
    } else {
      // Animate Closed
      gsap.to(submenuRef.current, {
        height: 0,
        opacity: 0,
        duration: 0.3,
        ease: "power2.in",
        onComplete: () => setOpen(false)
      });
      gsap.to(arrowRef.current, { rotate: 0, duration: 0.3 });
    }
  };

  return (
    <div className="mb-1">
      {/* Parent Link / Toggle */}
      <div
        onClick={toggleMenu}
        className={`flex items-center justify-between px-4 py-2.5 rounded-lg transition-colors cursor-pointer group ${
          open ? "bg-slate-800/50 text-white" : "text-gray-400 hover:bg-slate-800 hover:text-gray-200"
        }`}
      >
        <div className="flex items-center gap-3">
          <span className={`${open ? "text-blue-400" : "group-hover:text-blue-400"} transition-colors`}>
            {item.icon}
          </span>
          {/* If it has no subNav, make it a link; otherwise just text */}
          {!hasSub ? (
            <NavLink to={item.path} className="text-sm font-medium">
              {item.title}
            </NavLink>
          ) : (
            <span className="text-sm font-medium">{item.title}</span>
          )}
        </div>

        {hasSub && (
          <div ref={arrowRef} className="text-gray-500">
            <ChevronDown size={16} />
          </div>
        )}
      </div>

      {/* Submenu Container */}
      <div
        ref={submenuRef}
        className="overflow-hidden h-0 opacity-0 ml-9 space-y-1 border-l border-slate-700/50 mt-1"
      >
        {item.subNav?.map((subItem, index) => (
          <NavLink
            key={index}
            to={subItem.path}
            className={({ isActive }) => 
              `flex items-center gap-3 text-xs px-4 py-2 rounded-md transition-all ${
                isActive 
                ? "text-blue-400 bg-blue-400/10 font-semibold" 
                : "text-gray-500 hover:text-gray-200 hover:bg-slate-800"
              }`
            }
          >
            {subItem.icon}
            {subItem.title}
          </NavLink>
        ))}
      </div>
    </div>
  );
};

const Sidebar = () => {
  return (
    <aside className="fixed left-0 top-0 h-screen w-64 bg-[#0B0F1A] border-r border-slate-800 p-4 flex flex-col">
      {/* Logo Section */}
      <div className="flex items-center gap-3 px-2 mb-10">
        <img src={Logo} className="w-8 h-8"/>
        {/* <div className="w-8 h-8 bg-blue-600 rounded-lg flex items-center justify-center shadow-[0_0_15px_rgba(37,99,235,0.4)]">
          <span className="text-white font-bold text-xl">G</span>
        </div> */}
        <h2 className="text-lg font-bold tracking-tight text-white uppercase">
          Guardian<span className="text-blue-500">Grid</span>
        </h2>
      </div>

      {/* Navigation Scroll Area */}
      <div className="flex-1 overflow-y-auto space-y-1 custom-scrollbar">
        {sidebarData.map((item, index) => (
          <SubMenu item={item} key={index} />
        ))}
      </div>

      {/* Footer / User Info */}
      <div className="mt-auto pt-4 border-t border-slate-800">
        <div className="flex items-center gap-3 px-2">
          <div className="w-8 h-8 rounded-full bg-slate-700" />
          <div className="text-xs">
            <p className="text-white font-medium">System Admin</p>
            <p className="text-gray-500">v1.0.4-stable</p>
          </div>
        </div>
      </div>
    </aside>
  );
};

export default Sidebar;