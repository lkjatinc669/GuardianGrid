import React from 'react'
import Home from '../pages/Home';
import NotFound from '../pages/NotFound';
import { Route, Routes } from 'react-router-dom';

//Strict Authentication Route
import Stage1 from '../auth-pages/Stage1';
import Stage10 from '../auth-pages/Stage10';
import Stage11 from '../auth-pages/Stage11';

const AppRoutes = () => {
    return (
        <Routes>
            <Route path="/auth/ip" element={<Stage1/>} />
            <Route path="/auth/signup" element={<Stage10/>} />
            <Route path="/auth/login" element={<Stage11/>} />

            {/* Static Routes */}
            <Route path="/" element={<Home />} />
            {/* <Route path="/about" element={<About />} />
            <Route path="/login" element={<Login />} /> */}

            {/* Dynamic Route: ':username' can be anything */}
            {/* <Route path="/profile/:username" element={<UserProfile />} /> */}

            {/* Catch-all: Redirects any unknown URL to our 404 page */}
            <Route path="*" element={<NotFound />} />
        </Routes>
    )
}

export default AppRoutes;