import { Link } from "react-router-dom";

export default function NotFound() {
  return (
    <div className="fixed inset-0 flex items-center justify-center bg-linear-to-br from-slate-900 via-black to-slate-950 text-white z-50">
      
      {/* Glow background */}
      <div className="absolute w-100 h-100 bg-blue-500 opacity-20 blur-3xl animate-pulse rounded-full"></div>

      {/* Content */}
      <div className="z-10 text-center">
        
        {/* 404 with glitch effect */}
        <h1 className="text-[120px] font-extrabold tracking-widest animate-[glitch_1.5s_infinite]">
          404
        </h1>

        <h2 className="text-2xl mt-2">Lost in the Grid</h2>

        <p className="mt-2 text-gray-400">
          The page you're looking for doesn't exist or has been moved.
        </p>

        <Link
          to="/"
          className="inline-block mt-6 px-6 py-3 rounded-lg bg-blue-600 hover:bg-blue-700 transition transform hover:-translate-y-1 shadow-lg"
        >
          ⬅ Back to Dashboard
        </Link>
      </div>
    </div>
  );
}