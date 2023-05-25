import React, { useState } from 'react';
import { NavLink } from 'react-router-dom';
import { useSelector, useDispatch } from 'react-redux';
import Icon from '@mdi/react';
import {
  mdiHome,
  mdiFamilyTree,
  mdiTrendingUp,
  mdiCheckboxMultipleMarked,
  mdiPost,
  mdiCog,
  mdiChevronDown,
  mdiPlus,
} from '@mdi/js';

export default function Sidebar(props) {
  const hidden = false;
  return (
    <aside className={hidden ? 'sidebar sidebar-hidden' : 'sidebar'}>
      <h1 className="medica logo"/>
      <ul className="navigation">
        <NavigationSection
          path="../assets/home.png"
          title="Дашборд"
          status="raw"
          to="/dashboard"
          raw="true"
        />
        <NavigationSection
          path="../assets/report.png"
          title="Отчeты"
          status="soon"
          to="/skill-trees"
        />
        <NavigationSection
          path="../assets/staff.png"
          title="Сотрудники"
          status="raw"
          to="/habits"
          raw="true"
        />
        <NavigationSection
          path="../assets/analytics.png"
          title="Аналитика"
          to="/todo"
        />
        <NavigationSection
          path="../assets/lk.png"
          title="Личный кабинет"
          status="soon"
          to="/settings"
        />
      </ul>
    </aside>
  );
}

function NavigationSection(props) {
  const { path, to, title, status, raw } = props;
  return (
    <li>
      <NavLink
        className={({ isActive }) =>
          isActive ? 'navigation-section active' : 'navigation-section'
        }
        tabIndex="0"
        to={to}
      >
        <div className="icon-container">
          <img src={path} className="icon-auto" />
        </div>
        <p className="link">{title}</p>
      </NavLink>
    </li>
  );
}
