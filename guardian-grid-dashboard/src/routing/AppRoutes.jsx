import React from 'react'
import { Route, Routes, Navigate } from 'react-router-dom';
import Home from '../pages/Home';
import AgentDetails from '../pages/AgentDetails';
import NotFound from '../pages/NotFound';

// Auth Pages
import Stage1 from '../auth-pages/Stage1';
import Stage10 from '../auth-pages/Stage10';
import Stage11 from '../auth-pages/Stage11';

const ProtectedRoute = ({ children }) => {
    const isAuthenticated = localStorage.getItem("auth") === "true";
    if (!isAuthenticated) {
        return <Navigate to="/auth/login" replace />;
    }
    return children;
};

const AppRoutes = () => {
    return (
        <Routes>
            {/* Public Auth Routes */}
            <Route path="/auth/ip" element={<Stage1/>} />
            <Route path="/auth/signup" element={<Stage10/>} />
            <Route path="/auth/login" element={<Stage11/>} />

            {/* Protected Application Routes */}
            <Route path="/" element={
                <ProtectedRoute>
                    <Home />
                </ProtectedRoute>
            } />
            
            <Route path="/agent/:id" element={
                <ProtectedRoute>
                    <AgentDetails />
                </ProtectedRoute>
            } />

            {/* Catch-all */}
            <Route path="*" element={<NotFound />} />
        </Routes>
    )
}

export default AppRoutes;
