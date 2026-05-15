import React, { useState } from 'react';
import { BrowserRouter as Router } from 'react-router-dom';
import AppRoutes from './routing/AppRoutes';
import LoadingScreen from './components/LoadingScreen';
import "./App.css"

function App() {
  const [isAppReady, setIsAppReady] = useState(false);

  return (
    <Router>
      {!isAppReady && <LoadingScreen onComplete={() => setIsAppReady(true)} />}
      <div className={isAppReady ? "block" : "hidden"}>
        <AppRoutes />
      </div>
    </Router>
  );
}

export default App;
