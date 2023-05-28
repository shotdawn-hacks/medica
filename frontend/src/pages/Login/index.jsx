import React from 'react';
import { NavLink, useNavigate } from 'react-router-dom';
import useTitle from '../../hooks/useTitle';

export default function Login() {
  useTitle('Login');

  return (
    <div id="content-signup">
      <div className="login-container">
        <div className="login-header">
          <h1 className="login-header-text">Вход</h1>
        </div>
        <LoginForm />
      </div>
    </div>
  );
}

function LoginForm() {
  const navigate = useNavigate();

  const login = async (e) => {
    e.preventDefault();
    // const isSuccessful = await sendLoginRequest(data);
    // if (isSuccessful) navigate('/dashboard');
    navigate('..');
  };

  return (
    <form className="login-form" id="loginForm" onSubmit={login}>
      <div className="login-form-fields">
        <input className="login-form-field name" required placeholder="Логин" />
        <input
          className="login-form-field password"
          required
          placeholder="Пароль"
          type="password"
        />
        <h3 className="login-ref">
          <NavLink to="/">Forgot password?</NavLink>
        </h3>
        <button type="submit" className="login-form-field submit">
          Log in
        </button>
      </div>
    </form>
  );
}
