import React, { useState } from 'react';
import {
  BrowserRouter,
  Routes,
  Route,
  Outlet,
  useLocation,
} from 'react-router-dom';
import Sidebar from './components/Sidebar';
import Dashboard from './pages/Dashboard';
import NotFound from './pages/404';

const App = () => {
  const [loggedIn, setLoggedIn] = useState(true);

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<PrivateRoutes loggedIn={loggedIn} />}>
          <Route path="/" element={<Dashboard />} />
          <Route path="*" element={<NotFound />} />
        </Route>
        <Route path="/login" element={<div/>} />
      </Routes>
    </BrowserRouter>
  );
};

const PrivateRoutes = (params) => {
  const { loggedIn, changeAuth } = params;
  const location = useLocation();

  const [sidebarHidden, setSidebarHidden] = useState(true);

  const toggleSidebar = () => {
    setSidebarHidden(!sidebarHidden);
  };

  return loggedIn ? (
    <div id="content">
      <Sidebar hidden={sidebarHidden} />
      <Outlet />
      {/* <SidebarMobile /> */}
    </div>
  ) : (
    <> </>
  );
};

export default App;
