import React from 'react';
import Icon from '@mdi/react';
import {
  mdiMagnify,
  mdiBell,
  mdiStarOutline,
  mdiEyeOutline,
  mdiShareVariantOutline,
} from '@mdi/js';
import useTitle from '../../hooks/useTitle';

export default function Dashboard() {
  useTitle('Dashboard');

  return (
    <div className="dashboard">
    </div>
  );
}

function ProjectCard() {
  return (
    <a className="card" tabIndex="0">
      <div className="card-text">
        <h4>Signup-form</h4>
        <p>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
          eiusmod tempor incididunt ut labore et dolore magna aliqua.
        </p>
      </div>
      <div className="card-footer">
        <button>
          <Icon path={mdiStarOutline} alt="fav" className="icon" />
        </button>
        <button>
          <Icon path={mdiEyeOutline} alt="github" className="icon" />
        </button>
        <button>
          <Icon
            path={mdiShareVariantOutline}
            alt="share-variant-outline"
            className="icon"
          />
        </button>
      </div>
    </a>
  );
}

function Post(props) {
  const { title } = props;
  return (
    <a className="post" tabIndex="0">
      <h4>{title}</h4>
      <p>
        Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed
        do eiusmod tempor incididunt ut labore et dolore magna aliqua.
      </p>
    </a>
  );
}

function Skill(props) {
  const { color, name } = props;
  return (
    <button className="skill-icon" style={{ backgroundColor: color }}>
      {name}
    </button>
  );
}
