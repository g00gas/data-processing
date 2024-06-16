import React from 'react';
import { Link } from 'react-router-dom';

const Navbar: React.FC = () => {
    return (
        <nav className="navbar is-light">
            <div className="navbar-brand">
                <Link to="/" className="navbar-item">Apache Logs Dashboard</Link>
                <div className="navbar-burger burger" data-target="navbarMenu">
                    <span></span>
                    <span></span>
                    <span></span>
                </div>
            </div>
            <div id="navbarMenu" className="navbar-menu">
                <div className="navbar-start">
                    <Link to="/device-names" className="navbar-item">Nazwy urządzeń</Link>
                    <Link to="/http-methods" className="navbar-item">Metody HTTP</Link>
                    <Link to="/status-codes" className="navbar-item">Kody HTTP</Link>
                    <Link to="/user-agents" className="navbar-item">Agenty Użytkowników</Link>
                </div>
            </div>
        </nav>
    );
};

export default Navbar;
