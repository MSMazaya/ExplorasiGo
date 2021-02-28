import React from 'react'
import './style2.css'
const Phone = () => {
    return (
            <form class="form" method="POST">
                <h1>Gelang Anti Kekerasan</h1>
                <label htmlFor="pnumber">Input your phone number:</label>
                <input type="text" id="pnumber" name="phonenumber" placeholder="Masukkan Nomor Handphone"></input>
                <input type="submit" value="Submit"></input>
            </form>
    )
}

export default Phone

