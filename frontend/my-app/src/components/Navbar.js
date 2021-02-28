import React from 'react'
import './style1.css'
const Navbar = () => {
    return (
        <nav className="navbar">
            <ul className="littlenavbar">
                <li>Fitur</li>
                <li>Tentang Kami</li>
                <li className="masuk">Masuk</li>
                <li className="daftar">Daftar</li>
            </ul>
    </nav>
    )
}

export default Navbar
