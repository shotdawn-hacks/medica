import React from 'react';
import { Routes, Route, Navigate, Outlet, NavLink } from 'react-router-dom';
import ReportNew from './ReportNew';
import ReportArchive from './ReportArchive';
import useTitle from '../../hooks/useTitle';

export default function Reports() {
  useTitle('Reports');

  return (
    <Routes>
      <Route element={<ReportsLayout />}>
        <Route index element={<Navigate to="new" />} />
        <Route path="new" element={<ReportNew />} />
        <Route path="archive" element={<ReportArchive />} />
      </Route>
    </Routes>
  );
}

function ReportsLayout() {
  return (
    <main className="reports">
      <ModeSwitcher />
      <Outlet />
    </main>
  );
}

function ModeSwitcher() {
  return (
    <div className="report-change">
      <div className="report-change-container">
        <NavLink
          className={({ isActive }) => (isActive ? 'report-section left active' : 'report-section left')}
          to="new"
        >
          <p>Сгенерировать отчёт</p>
        </NavLink>
        <NavLink
          className={({ isActive }) => (isActive ? 'report-section right active' : 'report-section right')}
          to="archive"
        >
          <p>Архив отчётов</p>
        </NavLink>
      </div>
    </div>
  );
}
