import React from 'react'
import { Link } from 'react-router-dom'
const Navbar = () => {
    return (
    <nav className="navbar">
        <ul className="littlenavbar">
            <li>Fitur</li>
            <li>Tentang Kami</li>
            <li><Link to="/phone">Phone</Link></li>
            <li className="masuk">Masuk</li>
            <li className="daftar">Daftar</li>
        </ul>
    </nav>
    )
}

export default Navbar
